package gandalff

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type XptVersionType uint8

const (
	XPT_VERSION_5 XptVersionType = iota + 5
	XPT_VERSION_6
	XPT_VERSION_8
	XPT_VERSION_9
)

type XptReader struct {
	version XptVersionType
	path    string
	reader  io.Reader
	ctx     *Context
}

func NewXptReader(ctx *Context) *XptReader {
	return &XptReader{
		version: XPT_VERSION_5,
		path:    "",
		reader:  nil,
		ctx:     ctx,
	}
}

func (r *XptReader) SetVersion(version XptVersionType) *XptReader {
	r.version = version
	return r
}

func (r *XptReader) SetPath(path string) *XptReader {
	r.path = path
	return r
}

func (r *XptReader) SetReader(reader io.Reader) *XptReader {
	r.reader = reader
	return r
}

func (r *XptReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("XptReader: no reader specified")}
	}

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XptReader: no context specified")}
	}

	var err error
	var names []string
	var series []Series

	switch r.version {
	case XPT_VERSION_5, XPT_VERSION_6:
		names, series, err = readXPTv56(r.reader, r.ctx)
	case XPT_VERSION_8, XPT_VERSION_9:
		names, series, err = readXPTv89(r.reader, r.ctx)
	default:
		return BaseDataFrame{err: fmt.Errorf("XptReader: unknown version")}
	}

	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

type XptWriter struct {
	version   XptVersionType
	path      string
	writer    io.Writer
	dataframe DataFrame
}

func NewXptWriter() *XptWriter {
	return &XptWriter{
		version:   XPT_VERSION_5,
		path:      "",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *XptWriter) SetVersion(version XptVersionType) *XptWriter {
	w.version = version
	return w
}

func (w *XptWriter) SetPath(path string) *XptWriter {
	w.path = path
	return w
}

func (w *XptWriter) SetWriter(writer io.Writer) *XptWriter {
	w.writer = writer
	return w
}

func (w *XptWriter) SetDataFrame(dataframe DataFrame) *XptWriter {
	w.dataframe = dataframe
	return w
}

///////////////////////////////////////     SAS XPT v5/6     ///////////////////////////////////////

const (
	valueHeader = "HEADER RECORD*******"
	valueSAS    = "SAS     "
	valueLIB    = "SASLIB  "
)

// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-5-or-6-data-set-in-sas-transport-xport-format.pdf
const FIRST_HEADER_V56 = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000"

// This functions reads a SAS XPT file (versions 5/6).
func readXPTv56(reader io.Reader, ctx *Context) ([]string, []Series, error) {
	if ctx == nil {
		return nil, nil, fmt.Errorf("readCSV: no context specified")
	}

	var err error

	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for n, err := reader.Read(buffer); err == nil; n, err = reader.Read(buffer) {
		content = append(content, buffer[:n]...)
	}

	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	offset := 0

	// check header
	if string(content[0:20]) != valueHeader {
		return nil, nil, fmt.Errorf("readXPV56: invalid header")
	}
	offset += 80

	// check the first real header record
	if string(content[offset:offset+8]) != valueSAS ||
		string(content[offset+8:offset+16]) != valueSAS ||
		string(content[offset+16:offset+24]) != valueLIB {
		return nil, nil, fmt.Errorf("readXPV56: invalid first real header")
	}

	version := strings.Trim(string(content[offset+24:offset+32]), " ")
	if version != "" { // TODO: check version
		return nil, nil, fmt.Errorf("readXPTv56: invalid version '%s'", version)
	}
	offset += 80

	// check the second real header record
	offset += 80

	return nil, nil, nil
}

// This functions writes a SAS XPT file (versions 5/6).
func writeXPTv56(path string) error {
	buff := make([]byte, 0)

	buff = append(buff, []byte(fmt.Sprintf(
		"%s%8s%8s%8s%8s",
		FIRST_HEADER_V56,       // 1-80 HEADER RECORD
		"SAS", "SAS", "SASLIB", // 81-104
		time.Now().Format("ddMMMyy:hh:mm:ss")), // 105-128
	)...)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
}

///////////////////////////////////////     SAS XPT v8/9

// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-8-or-9-data-set-in-sas-transport-format.pdf

// This functions reads a SAS XPT file (versions 8/9).
func readXPTv89(reader io.Reader, ctx *Context) ([]string, []Series, error) {
	if ctx == nil {
		return nil, nil, fmt.Errorf("readXPTv89: no context specified")
	}

	var err error

	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for n, err := reader.Read(buffer); err == nil; n, err = reader.Read(buffer) {
		content = append(content, buffer[:n]...)
	}

	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	offset := 0

	///////////////////////////////////////
	// 1	The first header record consists ofthe following characterstring, in ASCII:
	// 		HEADER RECORD*******LIBV8 HEADER RECORD!!!!!!!000000000000000000000000000000
	if string(content[0:20]) != valueHeader {
		return nil, nil, fmt.Errorf("readXPTv89: invalid header")
	}
	offset += 80

	///////////////////////////////////////
	// 2	The first real header record
	if string(content[offset:offset+8]) != valueSAS ||
		string(content[offset+8:offset+16]) != valueSAS ||
		string(content[offset+16:offset+24]) != valueLIB {
		return nil, nil, fmt.Errorf("readXPTv89: invalid first real header")
	}

	version := strings.Trim(string(content[offset+24:offset+32]), " ")
	if version != "9.4" {
		return nil, nil, fmt.Errorf("readXPTv89: invalid version '%s'", version)
	}
	offset += 80

	///////////////////////////////////////
	// 3	Second real header record: ddMMMyy:hh:mm:ss
	offset += 80

	///////////////////////////////////////
	// 4	Member header records
	if string(content[offset:offset+20]) != valueHeader {
		return nil, nil, fmt.Errorf("readXPTv89: invalid member header")
	}

	namestrSize, err := strconv.Atoi(string(content[offset+74 : offset+78]))
	if err != nil {
		return nil, nil, fmt.Errorf("readXPTv89: invalid NAMESTR size '%s'", string(content[offset+74:offset+78]))
	}
	offset += 80

	switch namestrSize {
	case 140:
		// TODO: read namestr
	default:
		return nil, nil, fmt.Errorf("readXPTv89: invalid NAMESTR size '%d'", namestrSize)
	}

	// skip the member header
	offset += 80

	///////////////////////////////////////
	// 5	Member header data

	// skip the member header data
	offset += 80
	// skip the header record
	offset += 80

	///////////////////////////////////////
	// 6	Namestr headerrecord

	///////////////////////////////////////
	// 7	Namestr records

	return nil, nil, nil
}

// This functions writes a SAS XPT file (versions 8/9).
func writeXPTv89(path string) error {
	const firstHeaderRecord = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000  "

	var osVersion string
	switch runtime.GOOS {
	case "darwin":
		osVersion = "MacOS"
	case "linux":
		osVersion = "Linux"
	case "windows":
		osVersion = "X64_10HO"
	default:
		osVersion = runtime.GOOS
	}

	buff := make([]byte, 0)

	buff = append(buff, []byte(fmt.Sprintf(
		"%s%8s%8s%8s%8s%8s%24s%16s%80s",
		firstHeaderRecord,                     // 1-80 		First header record
		"SAS     ",                            // 81-88 	SAS
		"SAS     ",                            // 89-96 	SAS
		"SASLIB  ",                            // 97-104 	SASLIB
		"9.4     ",                            // 105-112 	SAS Version
		osVersion,                             // 113-120 	Operating System
		"",                                    //
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 153-176   Date/time created
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 177-200 	Second header record, date/time modified
	))...)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
}
