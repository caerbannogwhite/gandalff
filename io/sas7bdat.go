package io

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

// PageType defines the different types of pages
type PageType int

const (
	PageTypeMeta PageType = iota
	PageTypeData
	PageTypeMix
	PageTypeAMD
)

// SubheaderType defines the different types of subheaders
type SubheaderType int

const (
	SubheaderRowSize SubheaderType = iota
	SubheaderColSize
	SubheaderSubhdrCounts
	SubheaderColText
	SubheaderColName
	SubheaderColAttrs
	SubheaderColFormatAndLabel
	SubheaderColList
	SubheaderUnknown
)

// ColumnType defines the supported data types
type ColumnType int

const (
	ColumnTypeNumeric ColumnType = iota
	ColumnTypeCharacter
	ColumnTypeDate
	ColumnTypeTime
	ColumnTypeDateTime
	ColumnTypeCurrency
)

// Page represents a single page in the SAS7BDAT file
type Page struct {
	Type        int16
	BlockCount  int16
	SubhdrCount int16
	PageType    PageType
	Subheaders  []*Subheader
	Data        []byte
	Deleted     bool
}

// Subheader represents metadata within a page
type Subheader struct {
	Signature   SubheaderSignature
	Compression CompressionType
	Type        SubheaderType
	Length      int64
	Data        []byte
}

// Column represents a dataset column
type Column struct {
	Index  int
	Name   string
	Label  string
	Format string
	Type   ColumnType
	Length int
	Offset int64
}

// sas7bdatFile represents a parsed SAS7BDAT file
type sas7bdatFile struct {
	reader     io.ReadSeeker
	header     *FileHeader
	pages      []*Page
	columns    []*Column
	byteOrder  binary.ByteOrder
	pageSize   int64
	pageCount  int64
	headerSize int64
}

// FileHeader contains the file header information
type FileHeader struct {
	Magic           [32]byte
	A1              [8]byte
	A2              [8]byte
	DateCreated     time.Time
	DateModified    time.Time
	HeaderSize      int64
	PageSize        int64
	PageCount       int64
	SASRelease      string
	SASServer       string
	OSVersion       string
	OSMaker         string
	OSName          string
	DatasetName     string
	FileType        string
	U64             bool
	Compression     string
	RowLength       int64
	RowCount        int64
	ColCountP1      int64
	ColCountP2      int64
	MixPageRowCount int64
	LCS             int64
	LCP             int64
}

// ParseHeader reads and parses the SAS7BDAT file header
func (f *sas7bdatFile) parseHeader() error {
	// Read the first 288 bytes for initial header analysis
	headerBytes := make([]byte, 288)
	if _, err := f.reader.Read(headerBytes); err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	f.header = &FileHeader{}

	// Check magic number (offset 0, length 32)
	copy(f.header.Magic[:], headerBytes[0:32])
	if !f.validateMagicNumber() {
		return fmt.Errorf("invalid magic number")
	}

	// Detect endianness based on known patterns at offset 32-33
	if binary.LittleEndian.Uint16(headerBytes[32:34]) == 0x8000 {
		f.byteOrder = binary.LittleEndian
	} else if binary.BigEndian.Uint16(headerBytes[32:34]) == 0x8000 {
		f.byteOrder = binary.BigEndian
	} else {
		return fmt.Errorf("unable to determine byte order")
	}

	// Parse page size (offset 200, length 4 or 8 depending on format)
	f.pageSize = int64(f.byteOrder.Uint32(headerBytes[200:204]))
	if f.pageSize == 0 {
		// Try 8-byte format for 64-bit files
		f.pageSize = int64(f.byteOrder.Uint64(headerBytes[200:208]))
	}

	// Parse page count
	f.pageCount = int64(f.byteOrder.Uint32(headerBytes[204:208]))
	if f.pageCount == 0 {
		f.pageCount = int64(f.byteOrder.Uint64(headerBytes[208:216]))
	}

	return f.parseExtendedHeader()
}

func (f *sas7bdatFile) validateMagicNumber() bool {
	expectedMagic := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0xc2, 0xea, 0x81, 0x60,
		0xb3, 0x14, 0x11, 0xcf, 0xbd, 0x92, 0x08, 0x00,
		0x09, 0xc7, 0x31, 0x8c, 0x18, 0x1f, 0x10, 0x11}

	for i, b := range expectedMagic {
		if f.header.Magic[i] != b {
			return false
		}
	}
	return true
}

func (f *sas7bdatFile) parseExtendedHeader() error {
	// Seek to beginning and read full header
	if _, err := f.reader.Seek(0, io.SeekStart); err != nil {
		return err
	}

	headerSize := f.determineHeaderSize()
	f.headerSize = headerSize

	headerBytes := make([]byte, headerSize)
	if _, err := f.reader.Read(headerBytes); err != nil {
		return fmt.Errorf("failed to read extended header: %w", err)
	}

	// Parse dataset name (typically at offset 92, length 64)
	f.header.DatasetName = f.extractString(headerBytes[92:156])

	// Parse file type information
	f.header.FileType = f.extractString(headerBytes[156:164])

	// Parse timestamps (SAS dates are seconds since 1960-01-01)
	if len(headerBytes) > 164 {
		sasTime := f.byteOrder.Uint64(headerBytes[164:172])
		f.header.DateCreated = time.Unix(int64(sasTime)-int64(315619200), 0) // Adjust for epoch difference
	}

	// Detect u64 format extensions
	f.header.U64 = f.detectU64Format(headerBytes)

	return nil
}

// ParsePages reads and processes all pages in the file
func (f *sas7bdatFile) parsePages() error {
	f.pages = make([]*Page, f.pageCount)

	for i := int64(0); i < f.pageCount; i++ {
		offset := f.headerSize + (i * f.pageSize)
		if _, err := f.reader.Seek(offset, io.SeekStart); err != nil {
			return fmt.Errorf("failed to seek to page %d: %w", i, err)
		}

		page, err := f.parsePage()
		if err != nil {
			return fmt.Errorf("failed to parse page %d: %w", i, err)
		}
		f.pages[i] = page
	}

	return f.processSubheaders()
}

func (f *sas7bdatFile) parsePage() (*Page, error) {
	pageData := make([]byte, f.pageSize)
	if _, err := f.reader.Read(pageData); err != nil {
		return nil, err
	}

	page := &Page{
		Data: pageData,
	}

	// Parse page header (first 24 bytes typically)
	page.Type = int16(f.byteOrder.Uint16(pageData[16:18]))
	page.BlockCount = int16(f.byteOrder.Uint16(pageData[18:20]))
	page.SubhdrCount = int16(f.byteOrder.Uint16(pageData[20:22]))

	// Determine page type based on content
	page.PageType = f.determinePageType(page)

	// Parse subheaders if present
	if page.SubhdrCount > 0 {
		if err := f.parseSubheaders(page); err != nil {
			return nil, err
		}
	}

	return page, nil
}

func (f *sas7bdatFile) parseSubheaders(page *Page) error {
	offset := 24 // Standard page header size

	for i := 0; i < int(page.SubhdrCount); i++ {
		if offset+12 > len(page.Data) {
			break
		}

		subheader := &Subheader{}

		// Parse subheader pointer (offset, length, compression, type)
		subheaderOffset := f.byteOrder.Uint32(page.Data[offset : offset+4])
		subheaderLength := f.byteOrder.Uint32(page.Data[offset+4 : offset+8])
		compressionInfo := page.Data[offset+8]
		subheaderType := page.Data[offset+9]

		subheader.Length = int64(subheaderLength)
		subheader.Compression = CompressionType(compressionInfo)
		subheader.Type = f.identifySubheaderType(subheaderType)

		// Extract subheader data
		if subheaderOffset < uint32(len(page.Data)) &&
			subheaderOffset+subheaderLength <= uint32(len(page.Data)) {
			subheader.Data = page.Data[subheaderOffset : subheaderOffset+subheaderLength]
		}

		page.Subheaders = append(page.Subheaders, subheader)
		offset += 12 // Move to next subheader pointer
	}

	return nil
}

// NumericFormat represents different numeric storage formats
type NumericFormat struct {
	Size     int
	SignBits int
	ExpBits  int
	MantBits int
}

var numericFormats = map[int]NumericFormat{
	3: {24, 1, 11, 12}, // 24-bit format
	4: {32, 1, 11, 20}, // 32-bit format
	5: {40, 1, 11, 28}, // 40-bit format
	6: {48, 1, 11, 36}, // 48-bit format
	7: {56, 1, 11, 44}, // 56-bit format
	8: {64, 1, 11, 52}, // 64-bit format
}

func (f *sas7bdatFile) parseNumericValue(data []byte, format NumericFormat) (float64, error) {
	if len(data) < format.Size {
		return 0, fmt.Errorf("insufficient data for numeric format")
	}

	// Handle different numeric formats based on size
	switch format.Size {
	case 8:
		// Standard IEEE 754 double precision
		bits := f.byteOrder.Uint64(data[:8])
		return math.Float64frombits(bits), nil
	case 4:
		// IEEE 754 single precision
		bits := f.byteOrder.Uint32(data[:4])
		return float64(math.Float32frombits(bits)), nil
	default:
		// Custom SAS numeric formats
		return f.parseCustomNumeric(data[:format.Size], format)
	}
}

func (f *sas7bdatFile) parseCustomNumeric(data []byte, format NumericFormat) (float64, error) {
	// Extract sign, exponent, and mantissa based on format specification
	var value uint64
	for i, b := range data {
		value |= uint64(b) << (8 * (len(data) - 1 - i))
	}

	// Apply SAS-specific numeric conversion logic
	sign := (value >> (format.ExpBits + format.MantBits)) & 1
	exponent := (value >> format.MantBits) & ((1 << format.ExpBits) - 1)
	mantissa := value & ((1 << format.MantBits) - 1)

	// Convert to IEEE 754 equivalent
	if exponent == 0 && mantissa == 0 {
		return 0.0, nil
	}

	// Adjust for SAS-specific bias and scaling
	adjustedExp := int64(exponent) - (1 << (format.ExpBits - 1)) + 1023
	adjustedMant := mantissa << (52 - format.MantBits)

	result := (sign << 63) | (uint64(adjustedExp) << 52) | adjustedMant
	return math.Float64frombits(result), nil
}

func (f *sas7bdatFile) parseTimeValue(data []byte, format string) (time.Time, error) {
	numericValue, err := f.parseNumericValue(data, numericFormats[len(data)])
	if err != nil {
		return time.Time{}, err
	}

	sasEpoch := time.Date(1960, 1, 1, 0, 0, 0, 0, time.UTC)

	switch format {
	case "DATE":
		// Days since January 1, 1960
		return sasEpoch.AddDate(0, 0, int(numericValue)), nil
	case "TIME":
		// Seconds since midnight
		hours := int(numericValue) / 3600
		minutes := (int(numericValue) % 3600) / 60
		seconds := int(numericValue) % 60
		return time.Date(0, 1, 1, hours, minutes, seconds, 0, time.UTC), nil
	case "DATETIME":
		// Seconds since January 1, 1960
		return sasEpoch.Add(time.Duration(numericValue) * time.Second), nil
	default:
		return time.Time{}, fmt.Errorf("unsupported time format: %s", format)
	}
}

// PlatformInfo contains platform-specific parsing parameters
type PlatformInfo struct {
	Name        string
	Endianness  binary.ByteOrder
	Is64Bit     bool
	IsU64       bool
	PointerSize int
	IntSize     int
}

func (f *sas7bdatFile) detectPlatform() (*PlatformInfo, error) {
	platform := &PlatformInfo{
		Endianness: f.byteOrder,
	}

	// Analyze header fields to determine platform characteristics
	if len(f.header.OSName) > 0 {
		platform.Name = f.header.OSName

		// Platform-specific configurations
		switch {
		case strings.Contains(platform.Name, "WIN"):
			platform.Is64Bit = f.header.U64
			platform.PointerSize = f.getPointerSize(platform.Is64Bit)
		case strings.Contains(platform.Name, "Linux"):
			platform.Is64Bit = f.header.U64
			platform.PointerSize = f.getPointerSize(platform.Is64Bit)
		case strings.Contains(platform.Name, "SunOS"):
			platform.IsU64 = true
			platform.Is64Bit = true
			platform.PointerSize = 8
			platform.Endianness = binary.BigEndian
		default:
			// Default assumptions for unknown platforms
			platform.PointerSize = 4
		}
	}

	return platform, nil
}

func (f *sas7bdatFile) getPointerSize(is64Bit bool) int {
	if is64Bit {
		return 8
	}
	return 4
}

func (f *sas7bdatFile) readInteger(data []byte, size int) (int64, error) {
	switch size {
	case 1:
		return int64(data[0]), nil
	case 2:
		return int64(f.byteOrder.Uint16(data)), nil
	case 4:
		return int64(f.byteOrder.Uint32(data)), nil
	case 8:
		return int64(f.byteOrder.Uint64(data)), nil
	default:
		return 0, fmt.Errorf("unsupported integer size: %d", size)
	}
}

func (f *sas7bdatFile) adaptToU64Format() {
	if f.header.U64 {
		// Adjust parsing parameters for u64 format
		f.adjustFieldOffsets()
		f.adjustSubheaderParsing()
	}
}

// GetDataset extracts the complete dataset as a slice of maps
func (f *sas7bdatFile) getDataset() ([]map[string]interface{}, error) {
	if len(f.columns) == 0 {
		return nil, fmt.Errorf("no columns defined")
	}

	var dataset []map[string]interface{}

	for _, page := range f.pages {
		if page.PageType == PageTypeData || page.PageType == PageTypeMix {
			rows, err := f.extractRowsFromPage(page)
			if err != nil {
				return nil, fmt.Errorf("failed to extract rows: %w", err)
			}
			dataset = append(dataset, rows...)
		}
	}

	return dataset, nil
}

func (f *sas7bdatFile) extractRowsFromPage(page *Page) ([]map[string]interface{}, error) {
	var rows []map[string]interface{}

	// Determine data start offset (after subheaders)
	dataOffset := f.calculateDataOffset(page)
	rowSize := f.header.RowLength

	// Extract individual rows
	for offset := dataOffset; offset+int(rowSize) <= len(page.Data); offset += int(rowSize) {
		row := make(map[string]interface{})
		rowData := page.Data[offset : offset+int(rowSize)]

		// Parse each column value
		for _, col := range f.columns {
			if col.Offset+int64(col.Length) <= int64(len(rowData)) {
				colData := rowData[col.Offset : col.Offset+int64(col.Length)]
				value, err := f.parseColumnValue(colData, col)
				if err != nil {
					return nil, fmt.Errorf("failed to parse column %s: %w", col.Name, err)
				}
				row[col.Name] = value
			}
		}

		rows = append(rows, row)
	}

	return rows, nil
}

func (f *sas7bdatFile) parseColumnValue(data []byte, col *Column) (interface{}, error) {
	switch col.Type {
	case ColumnTypeNumeric:
		format := numericFormats[col.Length]
		return f.parseNumericValue(data, format)
	case ColumnTypeCharacter:
		return f.parseStringValue(data), nil
	case ColumnTypeDate, ColumnTypeTime, ColumnTypeDateTime:
		return f.parseTimeValue(data, col.Format)
	default:
		return nil, fmt.Errorf("unsupported column type: %v", col.Type)
	}
}

// GetColumnInfo returns metadata about dataset columns
func (f *sas7bdatFile) GetColumnInfo() []*Column {
	return f.columns
}

// GetFileInfo returns general file metadata
func (f *sas7bdatFile) GetFileInfo() *FileHeader {
	return f.header
}

func readSas7bdata(reader io.ReadSeeker) (*IoData, error) {

	parser := &sas7bdatFile{
		reader: reader,
	}

	// Parse file header
	if err := parser.parseHeader(); err != nil {
		return nil, fmt.Errorf("header parsing failed: %w", err)
	}

	// Detect and configure platform-specific settings
	platform, err := parser.detectPlatform()
	if err != nil {
		return nil, fmt.Errorf("platform detection failed: %w", err)
	}

	parser.configurePlatform(platform)

	// Parse all pages and extract column metadata
	if err := parser.parsePages(); err != nil {
		return nil, fmt.Errorf("page parsing failed: %w", err)
	}

	return nil, nil
}
