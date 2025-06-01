package io

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/caerbannogwhite/gandalff"
	"github.com/caerbannogwhite/gandalff/meta"
	"github.com/caerbannogwhite/gandalff/series"
)

// https://www.loc.gov/preservation/digital/formats/fdd/fdd000464.shtml

type XptVersionType uint8

const SAS_YEAR_THRESHOLD = 50

const (
	XPT_VERSION_5 XptVersionType = iota + 5
	XPT_VERSION_6
	XPT_VERSION_8 XptVersionType = iota + 6
	XPT_VERSION_9
)

type XptReader struct {
	guessVersion    bool
	maxObservations int
	version         XptVersionType
	byteOrder       binary.ByteOrder
	path            string
	reader          io.Reader
	ctx             *gandalff.Context
}

func NewXptReader(ctx *gandalff.Context) *XptReader {
	return &XptReader{
		guessVersion:    false,
		maxObservations: -1,
		version:         XPT_VERSION_8,
		byteOrder:       binary.BigEndian,
		path:            "",
		reader:          nil,
		ctx:             ctx,
	}
}

func (r *XptReader) SetMaxObservations(maxObservations int) *XptReader {
	r.maxObservations = maxObservations
	return r
}

func (r *XptReader) SetVersion(version XptVersionType) *XptReader {
	r.version = version
	return r
}

func (r *XptReader) GuessVersion() *XptReader {
	r.guessVersion = true
	return r
}

func (r *XptReader) SetByteOrder(byteOrder binary.ByteOrder) *XptReader {
	r.byteOrder = byteOrder
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

func (r *XptReader) Read() *IoData {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return &IoData{Error: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return &IoData{Error: fmt.Errorf("XptReader: no reader specified")}
	}

	if r.ctx == nil {
		return &IoData{Error: fmt.Errorf("XptReader: no context specified")}
	}

	var version XptVersionType
	var content []byte
	var err error

	if r.guessVersion {
		version, content, err = guessXptVersion(r.reader, r.ctx)
		if err != nil {
			return &IoData{Error: err}
		}

		r.version = version
		r.reader = bytes.NewReader(content)
	}

	var ioData *IoData

	switch r.version {
	case XPT_VERSION_5, XPT_VERSION_6:
		ioData, err = readXptV56(r.reader, r.maxObservations, r.byteOrder, r.ctx)
	case XPT_VERSION_8, XPT_VERSION_9:
		ioData, err = readXptV89(r.reader, r.maxObservations, r.byteOrder, r.ctx)
	default:
		return &IoData{Error: fmt.Errorf("XptReader: unknown version")}
	}

	if err != nil {
		return &IoData{Error: err}
	}

	return ioData
}

type XptWriter struct {
	version   XptVersionType
	byteOrder binary.ByteOrder
	path      string
	writer    io.Writer
	ioData    *IoData
}

func NewXptWriter() *XptWriter {
	return &XptWriter{
		version:   XPT_VERSION_8,
		byteOrder: binary.BigEndian,
		path:      "",
		writer:    nil,
		ioData:    nil,
	}
}

func (w *XptWriter) SetVersion(version XptVersionType) *XptWriter {
	w.version = version
	return w
}

func (w *XptWriter) SetByteOrder(byteOrder binary.ByteOrder) *XptWriter {
	w.byteOrder = byteOrder
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

func (w *XptWriter) SetIoData(ioData *IoData) *XptWriter {
	w.ioData = ioData
	return w
}

func (w *XptWriter) Write() error {
	if w.ioData == nil {
		return fmt.Errorf("XptWriter: no ioData specified")
	}

	if w.path != "" {
		file, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return fmt.Errorf("XptWriter: %w", err)
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		return fmt.Errorf("XptWriter: no writer specified")
	}

	var err error
	switch w.version {
	case XPT_VERSION_5, XPT_VERSION_6:
		err = writeXPTv56(w.ioData, w.writer, w.byteOrder)
	case XPT_VERSION_8, XPT_VERSION_9:
		err = writeXPTv89(w.ioData, w.writer, w.byteOrder)
	default:
		return fmt.Errorf("XptWriter: unknown SAS version '%d'", w.version)
	}

	if err != nil {
		return fmt.Errorf("XptWriter: %w", err)
	}

	return nil
}

const (
	valueHeaderStart = "HEADER RECORD*******"
	valueSas         = "SAS     "
	valueSasLib      = "SASLIB  "
	valueSasData     = "SASDATA "
)

// This functions guesses the version of a SAS XPT file.
func guessXptVersion(reader io.Reader, ctx *gandalff.Context) (XptVersionType, []byte, error) {
	if ctx == nil {
		return 0, nil, fmt.Errorf("guessXptVersion: no context specified")
	}

	var err error

	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for n, err := reader.Read(buffer); err == nil; n, err = reader.Read(buffer) {
		content = append(content, buffer[:n]...)
	}

	if err != nil && err != io.EOF {
		return 0, nil, fmt.Errorf("guessXptVersion: %w", err)
	}

	offset := 0

	///////////////////////////////////////
	// 1	The first header record consists ofthe following characterstring, in ASCII:
	// 		HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000
	if string(content[0:20]) != valueHeaderStart {
		return XPT_VERSION_8, content, nil
	}
	offset += 80

	///////////////////////////////////////
	// 2	The first real header record
	if string(content[offset:offset+8]) != valueSas ||
		string(content[offset+8:offset+16]) != valueSas ||
		string(content[offset+16:offset+24]) != valueSasLib {
		return XPT_VERSION_8, content, nil
	}

	sasLibVersion := strings.Trim(string(content[offset+24:offset+32]), " ")
	switch strings.Split(sasLibVersion, ".")[0] {
	case "5":
		return XPT_VERSION_5, content, nil
	case "6":
		return XPT_VERSION_6, content, nil
	case "8":
		return XPT_VERSION_8, content, nil
	case "9":
		return XPT_VERSION_9, content, nil
	default:
		return 0, nil, fmt.Errorf("guessXptVersion: invalid version '%s'", sasLibVersion)
	}
}

///////////////////////////////////////     SAS XPT v5/6
//
// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-5-or-6-data-set-in-sas-transport-xport-format.pdf

type __NAMESTRv56 struct {
	ntype  int16    // VARIABLE TYPE: 1=NUMERIC, 2=CHAR 	(bytes: 000 to 002)
	nhfun  int16    // HASH OF NNAME (always 0)				(bytes: 002 to 004)
	nlng   int16    // LENGTH OF VARIABLE IN OBSERVATION	(bytes: 004 to 006)
	nvar0  int16    // VARNUM								(bytes: 006 to 008)
	nname  [8]byte  // NAME OF VARIABLE						(bytes: 008 to 016)
	nlabel [40]byte // LABEL OF VARIABLE					(bytes: 016 to 056)
	nform  [8]byte  // NAME OF FORMAT						(bytes: 056 to 064)
	nfl    int16    // FORMAT FIELD LENGTH OR 0				(bytes: 064 to 066)
	nfd    int16    // FORMAT NUMBER OF DECIMALS			(bytes: 066 to 068)
	nfj    int16    // 0=LEFT JUSTIFICATION, 1=RIGHT JUST	(bytes: 068 to 070)
	nfill  [2]byte  // (UNUSED, FOR ALIGNMENT AND FUTURE)	(bytes: 070 to 072)
	niform [8]byte  // NAME OF INPUT FORMAT					(bytes: 072 to 080)
	nifl   int16    // INFORMAT LENGTH ATTRIBUTE			(bytes: 080 to 082)
	nifd   int16    // INFORMAT NUMBER OF DECIMALS			(bytes: 082 to 084)
	npos   int32    // POSITION OF VALUE IN OBSERVATION		(bytes: 084 to 088)
	rest   [52]byte // remaining fields are irrelevant		(bytes: 088 to 140)
}

func NewNamestrV56() *__NAMESTRv56 {
	return &__NAMESTRv56{
		ntype: 0,
		nhfun: 0,
		nlng:  0,
		nvar0: 0,
		nname: [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nlabel: [40]byte{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		},
		nform:  [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nfl:    0,
		nfd:    0,
		nfj:    0,
		nfill:  [2]byte{},
		niform: [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nifl:   0,
		nifd:   0,
		npos:   0,
		rest:   [52]byte{},
	}
}

func (nms *__NAMESTRv56) FromBinary(buffer []byte, byteOrder binary.ByteOrder) {
	nms.ntype = int16(byteOrder.Uint16(buffer[0:2]))
	nms.nhfun = int16(byteOrder.Uint16(buffer[2:4]))
	nms.nlng = int16(byteOrder.Uint16(buffer[4:6]))
	nms.nvar0 = int16(byteOrder.Uint16(buffer[6:8]))
	copy(nms.nname[:], buffer[8:16])
	copy(nms.nlabel[:], buffer[16:56])
	copy(nms.nform[:], buffer[56:64])
	nms.nfl = int16(byteOrder.Uint16(buffer[64:66]))
	nms.nfd = int16(byteOrder.Uint16(buffer[66:68]))
	nms.nfj = int16(byteOrder.Uint16(buffer[68:70]))
	copy(nms.nfill[:], buffer[70:72])
	copy(nms.niform[:], buffer[72:80])
	nms.nifl = int16(byteOrder.Uint16(buffer[80:82]))
	nms.nifd = int16(byteOrder.Uint16(buffer[82:84]))
	nms.npos = int32(byteOrder.Uint32(buffer[84:88]))
	copy(nms.rest[:], buffer[88:140])
}

func (nms *__NAMESTRv56) ToBinary(byteOrder binary.ByteOrder) []byte {
	buffer := make([]byte, 140)

	byteOrder.PutUint16(buffer[0:2], uint16(nms.ntype))
	byteOrder.PutUint16(buffer[2:4], uint16(nms.nhfun))
	byteOrder.PutUint16(buffer[4:6], uint16(nms.nlng))
	byteOrder.PutUint16(buffer[6:8], uint16(nms.nvar0))
	copy(buffer[8:16], nms.nname[:])
	copy(buffer[16:56], nms.nlabel[:])
	copy(buffer[56:64], nms.nform[:])
	byteOrder.PutUint16(buffer[64:66], uint16(nms.nfl))
	byteOrder.PutUint16(buffer[66:68], uint16(nms.nfd))
	byteOrder.PutUint16(buffer[68:70], uint16(nms.nfj))
	copy(buffer[70:72], nms.nfill[:])
	copy(buffer[72:80], nms.niform[:])
	byteOrder.PutUint16(buffer[80:82], uint16(nms.nifl))
	byteOrder.PutUint16(buffer[82:84], uint16(nms.nifd))
	byteOrder.PutUint32(buffer[84:88], uint32(nms.npos))
	copy(buffer[88:140], nms.rest[:])

	return buffer
}

func (nms *__NAMESTRv56) ToString() string {
	return fmt.Sprintf(
		"NAMESTRv56[\n"+
			"\tntype:  %d\n"+
			"\tnhfun:  %d\n"+
			"\tnlng:   %d\n"+
			"\tnvar0:  %d\n"+
			"\tnname:  %s\n"+
			"\tnlabel: %s\n"+
			"\tnform:  %s\n"+
			"\tnfl:    %d\n"+
			"\tnfd:    %d\n"+
			"\tnfj:    %d\n"+
			"\tnfill:  %s\n"+
			"\tniform: %s\n"+
			"\tnifl:   %d\n"+
			"\tnifd:   %d\n"+
			"\tnpos:   %d\n"+
			"\trest:   %s\n"+
			"]\n",
		nms.ntype,
		nms.nhfun,
		nms.nlng,
		nms.nvar0,
		string(nms.nname[:]),
		string(nms.nlabel[:]),
		string(nms.nform[:]),
		nms.nfl,
		nms.nfd,
		nms.nfj,
		string(nms.nfill[:]),
		string(nms.niform[:]),
		nms.nifl,
		nms.nifd,
		nms.npos,
		string(nms.rest[:]),
	)
}

// This functions reads a SAS XPT file (versions 5/6).
func readXptV56(reader io.Reader, maxObservations int, byteOrder binary.ByteOrder, ctx *gandalff.Context) (*IoData, error) {
	if ctx == nil {
		return nil, fmt.Errorf("readXptV56: no context specified")
	}
	var err error
	var fileMeta FileMeta

	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for n, err := reader.Read(buffer); err == nil; n, err = reader.Read(buffer) {
		content = append(content, buffer[:n]...)
	}

	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("readXptV56: %w", err)
	}

	offset := 0

	///////////////////////////////////////
	// 1	The first header record consists ofthe following characterstring, in ASCII:
	// 		HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000
	if string(content[0:20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV56: invalid header")
	}
	offset += 80

	///////////////////////////////////////
	// 2	The first real header record
	if string(content[offset:offset+8]) != valueSas ||
		string(content[offset+8:offset+16]) != valueSas ||
		string(content[offset+16:offset+24]) != valueSasLib {
		return nil, fmt.Errorf("readXptV56: invalid first real header")
	}

	sasLibVersion := strings.Trim(string(content[offset+24:offset+32]), " ")
	switch v := strings.Split(sasLibVersion, ".")[0]; v {
	case "5", "6":
		fileMeta.SasLibVersion = sasLibVersion

	default:
		return nil, fmt.Errorf("readXptV56: invalid version '%s'", sasLibVersion)
	}

	// Read SAS OS
	fileMeta.SasOs = string(content[offset+32 : offset+40])

	// Read Creation Date
	// ie: 04APR12:22:16:21
	creationDate := strings.Trim(string(content[offset+64:offset+80]), " ")
	fileMeta.Created, err = parseSasDate(creationDate)
	if err != nil {
		return nil, fmt.Errorf("readXptV56: invalid creation date '%s'", creationDate)
	}
	offset += 80

	///////////////////////////////////////
	// 3	Second real header record: ddMMMyy:hh:mm:ss
	lastModifiedDate := strings.Trim(string(content[offset:offset+80]), " ")
	fileMeta.LastModified, err = parseSasDate(lastModifiedDate)
	if err != nil {
		return nil, fmt.Errorf("readXptV56: invalid last modified date '%s'", lastModifiedDate)
	}
	offset += 80

	///////////////////////////////////////
	// 4	Member header records
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV56: invalid member header")
	}

	namestrSize := 140
	offset += 80

	// skip the member header
	offset += 80

	///////////////////////////////////////
	// 5	Member header data
	dsName := string(content[offset+8 : offset+16])
	fileMeta.SasDsName = strings.Trim(dsName, " ")

	sasDataVersion := string(content[offset+24 : offset+32])
	fileMeta.SasDataVersion = strings.Trim(sasDataVersion, " ")

	// skip the member header data
	offset += 80
	// skip the header record
	offset += 80

	///////////////////////////////////////
	// 6	Namestr headerrecord
	var varsNum int
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV56: invalid namestr header")
	}

	// get number of variables
	n, err := strconv.ParseInt(string(content[offset+48:offset+58]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("readXptV56: invalid number of variables '%s'", string(content[offset+24:offset+32]))
	}
	varsNum = int(n)
	offset += 80

	///////////////////////////////////////
	// 7	Namestr records

	seriesMeta := make([]SeriesMeta, varsNum)
	namestrs := make([]__NAMESTRv89, varsNum)

	// read namestr
	for i := 0; i < varsNum; i++ {
		namestrs[i].FromBinary(content[offset:offset+140], byteOrder)
		type_ := meta.Float64Type
		if namestrs[i].ntype == 2 {
			type_ = meta.StringType
		}

		seriesMeta[i] = SeriesMeta{
			Name:   strings.Trim(string(namestrs[i].nname[:]), " "),
			Label:  strings.Trim(string(namestrs[i].nlabel[:]), " "),
			Length: int(namestrs[i].nlng),
			Type:   type_,
		}

		offset += namestrSize
	}

	// skip the padding
	if p := ((namestrSize * varsNum) % 80); p != 0 {
		offset += 80 - p
	}

	///////////////////////////////////////
	// 8	Observation header

	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV56: invalid observation header")
	}

	// skip the observation header
	offset += 80

	///////////////////////////////////////
	// 9	Data records

	nulls := make([][]bool, varsNum)
	values := make([]interface{}, varsNum)

	for i := 0; i < varsNum; i++ {
		nulls[i] = make([]bool, 0)

		switch namestrs[i].ntype {
		case 1:
			values[i] = make([]float64, 0)
		case 2:
			values[i] = make([]string, 0)
		default:
			return nil, fmt.Errorf("readXptV56: invalid variable type '%d'", namestrs[i].ntype)
		}
	}

	// read observations by rows
	if maxObservations < 0 {
		maxObservations = math.MaxInt32
	}

	var tmp []byte
	rowCounter := 0
	for offset < len(content) && rowCounter < maxObservations {

		allNulls := true
		for i := offset; i < len(content); i++ {
			if content[i] != '\x20' {
				allNulls = false
				break
			}
		}

		if allNulls {
			break
		}

		rowLen := 0
		for i := 0; i < varsNum; i++ {
			tmp = make([]byte, namestrs[i].nlng)
			copy(tmp, content[offset+int(namestrs[i].npos):offset+int(namestrs[i].npos)+int(namestrs[i].nlng)])

			switch namestrs[i].ntype {

			// NUMERIC
			case 1:
				f, err := NewSasFloat(tmp).ToIeee(byteOrder)
				if err != nil {
					return nil, fmt.Errorf("readXptV56: %w", err)
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}

				values[i] = append(values[i].([]float64), f)

			// CHAR
			case 2:
				s := strings.Trim(string(tmp), " ")

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
		rowCounter++
	}

	_series := make([]series.Series, varsNum)
	for i := 0; i < varsNum; i++ {
		switch t := values[i].(type) {
		case []float64:
			_series[i] = series.NewSeriesFloat64(t, nulls[i], false, ctx)
		case []string:
			_series[i] = series.NewSeriesString(t, nulls[i], false, ctx)
		}
	}

	return &IoData{
		FileMeta:   fileMeta,
		SeriesMeta: seriesMeta,
		Series:     _series,
	}, nil
}

// This functions writes a SAS XPT file (versions 5/6).
func writeXPTv56(ioData *IoData, writer io.Writer, byteOrder binary.ByteOrder) error {
	// TODO: implement
	return fmt.Errorf("writeXPTv56: not implemented")
}

///////////////////////////////////////     SAS XPT v8/9
//
// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-8-or-9-data-set-in-sas-transport-format.pdf

type __NAMESTRv89 struct {
	ntype    int16    // VARIABLE TYPE: 1=NUMERIC, 2=CHAR	(bytes: 000 to 002)
	nhfun    int16    // HASH OF NNAME (always 0)			(bytes: 002 to 004)
	nlng     int16    // LENGTH OF VARIABLE IN OBSERVATION	(bytes: 004 to 006)
	nvar0    int16    // VARNUM								(bytes: 006 to 008)
	nname    [8]byte  // NAME OF VARIABLE					(bytes: 008 to 016)
	nlabel   [40]byte // LABEL OF VARIABLE					(bytes: 016 to 056)
	nform    [8]byte  // NAME OF FORMAT						(bytes: 056 to 064)
	nfl      int16    // FORMAT FIELD LENGTH OR 0			(bytes: 064 to 066)
	nfd      int16    // FORMAT NUMBER OF DECIMALS			(bytes: 066 to 068)
	nfj      int16    // 0=LEFT JUSTIFICATION, 1=RIGHT JUST	(bytes: 068 to 070)
	nfill    [2]byte  // (UNUSED, FOR ALIGNMENT AND FUTURE)	(bytes: 070 to 072)
	niform   [8]byte  // NAME OF INPUT FORMAT				(bytes: 072 to 080)
	nifl     int16    // INFORMAT LENGTH ATTRIBUTE			(bytes: 080 to 082)
	nifd     int16    // INFORMAT NUMBER OF DECIMALS		(bytes: 082 to 084)
	npos     int32    // POSITION OF VALUE IN OBSERVATION	(bytes: 084 to 088)
	longname [32]byte // long name for Version 8-style		(bytes: 088 to 120)
	lablen   int16    // length of label					(bytes: 120 to 122)
	rest     [18]byte // remaining fields are irrelevant	(bytes: 122 to 140)
}

func NewNamestrV89() *__NAMESTRv89 {
	return &__NAMESTRv89{
		ntype: 0,
		nhfun: 0,
		nlng:  0,
		nvar0: 0,
		nname: [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nlabel: [40]byte{
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
			' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
		},
		nform:  [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nfl:    0,
		nfd:    0,
		nfj:    0,
		nfill:  [2]byte{},
		niform: [8]byte{' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '},
		nifl:   0,
		nifd:   0,
		npos:   0,
		rest:   [18]byte{},
	}
}

func (nms *__NAMESTRv89) FromBinary(buffer []byte, byteOrder binary.ByteOrder) {
	nms.ntype = int16(byteOrder.Uint16(buffer[0:2]))
	nms.nhfun = int16(byteOrder.Uint16(buffer[2:4]))
	nms.nlng = int16(byteOrder.Uint16(buffer[4:6]))
	nms.nvar0 = int16(byteOrder.Uint16(buffer[6:8]))
	copy(nms.nname[:], buffer[8:16])
	copy(nms.nlabel[:], buffer[16:56])
	copy(nms.nform[:], buffer[56:64])
	nms.nfl = int16(byteOrder.Uint16(buffer[64:66]))
	nms.nfd = int16(byteOrder.Uint16(buffer[66:68]))
	nms.nfj = int16(byteOrder.Uint16(buffer[68:70]))
	copy(nms.nfill[:], buffer[70:72])
	copy(nms.niform[:], buffer[72:80])
	nms.nifl = int16(byteOrder.Uint16(buffer[80:82]))
	nms.nifd = int16(byteOrder.Uint16(buffer[82:84]))
	nms.npos = int32(byteOrder.Uint32(buffer[84:88]))
	copy(nms.longname[:], buffer[88:120])
	nms.lablen = int16(byteOrder.Uint16(buffer[120:122]))
	// copy(nms.rest[:], buffer[122:140])
}

func (nms *__NAMESTRv89) ToBinary(byteOrder binary.ByteOrder) []byte {
	buffer := make([]byte, 140)

	byteOrder.PutUint16(buffer[0:2], uint16(nms.ntype))
	byteOrder.PutUint16(buffer[2:4], uint16(nms.nhfun))
	byteOrder.PutUint16(buffer[4:6], uint16(nms.nlng))
	byteOrder.PutUint16(buffer[6:8], uint16(nms.nvar0))
	copy(buffer[8:16], nms.nname[:])
	copy(buffer[16:56], nms.nlabel[:])
	copy(buffer[56:64], nms.nform[:])
	byteOrder.PutUint16(buffer[64:66], uint16(nms.nfl))
	byteOrder.PutUint16(buffer[66:68], uint16(nms.nfd))
	byteOrder.PutUint16(buffer[68:70], uint16(nms.nfj))
	copy(buffer[70:72], nms.nfill[:])
	copy(buffer[72:80], nms.niform[:])
	byteOrder.PutUint16(buffer[80:82], uint16(nms.nifl))
	byteOrder.PutUint16(buffer[82:84], uint16(nms.nifd))
	byteOrder.PutUint32(buffer[84:88], uint32(nms.npos))
	copy(buffer[88:120], nms.longname[:])
	byteOrder.PutUint16(buffer[120:122], uint16(nms.lablen))
	// copy(buffer[122:140], nms.rest[:])

	return buffer
}

func (nms *__NAMESTRv89) String() string {
	return fmt.Sprintf(
		"NAMESTRv89[\n"+
			"\tntype:    %d\n"+
			"\tnhfun:    %d\n"+
			"\tnlng:     %d\n"+
			"\tnvar0:    %d\n"+
			"\tnname:    %s\n"+
			"\tnlabel:   %s\n"+
			"\tnform:    %s\n"+
			"\tnfl:      %d\n"+
			"\tnfd:      %d\n"+
			"\tnfj:      %d\n"+
			"\tnfill:    %s\n"+
			"\tniform:   %s\n"+
			"\tnifl:     %d\n"+
			"\tnifd:     %d\n"+
			"\tnpos:     %d\n"+
			"\tlongname: %s\n"+
			"\tlablen:   %d\n"+
			"\trest:     %s\n"+
			"]\n",
		nms.ntype,
		nms.nhfun,
		nms.nlng,
		nms.nvar0,
		string(nms.nname[:]),
		string(nms.nlabel[:]),
		string(nms.nform[:]),
		nms.nfl,
		nms.nfd,
		nms.nfj,
		string(nms.nfill[:]),
		string(nms.niform[:]),
		nms.nifl,
		nms.nifd,
		nms.npos,
		string(nms.longname[:]),
		nms.lablen,
		string(nms.rest[:]),
	)
}

// This functions reads a SAS XPT file (versions 8/9).
func readXptV89(reader io.Reader, maxObservations int, byteOrder binary.ByteOrder, ctx *gandalff.Context) (*IoData, error) {
	if ctx == nil {
		return nil, fmt.Errorf("readXptV89: no context specified")
	}

	var err error
	var fileMeta FileMeta

	content := make([]byte, 0)
	buffer := make([]byte, 1024)
	for n, err := reader.Read(buffer); err == nil; n, err = reader.Read(buffer) {
		content = append(content, buffer[:n]...)
	}

	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("readXptV89: %w", err)
	}

	offset := 0

	///////////////////////////////////////
	// 1	The first header record consists ofthe following characterstring, in ASCII:
	// 		HEADER RECORD*******LIBV8 HEADER RECORD!!!!!!!000000000000000000000000000000
	if string(content[0:20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV89: invalid header")
	}
	offset += 80

	///////////////////////////////////////
	// 2	The first real header record
	if string(content[offset:offset+8]) != valueSas ||
		string(content[offset+8:offset+16]) != valueSas ||
		string(content[offset+16:offset+24]) != valueSasLib {
		return nil, fmt.Errorf("readXptV89: invalid first real header")
	}

	sasLibVersion := strings.Trim(string(content[offset+24:offset+32]), " ")
	switch v := strings.Split(sasLibVersion, ".")[0]; v {
	case "8", "9":
		fileMeta.SasLibVersion = sasLibVersion

	default:
		return nil, fmt.Errorf("readXptV89: invalid version '%s'", sasLibVersion)
	}

	// Read SAS OS
	fileMeta.SasOs = string(content[offset+32 : offset+40])

	// Read Creation Date
	// ie: 04APR12:22:16:21
	creationDate := strings.Trim(string(content[offset+64:offset+80]), " ")
	fileMeta.Created, err = parseSasDate(creationDate)
	if err != nil {
		return nil, fmt.Errorf("readXptV89: invalid creation date '%s'", creationDate)
	}
	offset += 80

	///////////////////////////////////////
	// 3	Second real header record: ddMMMyy:hh:mm:ss
	lastModifiedDate := strings.Trim(string(content[offset:offset+80]), " ")
	fileMeta.LastModified, err = parseSasDate(lastModifiedDate)
	if err != nil {
		return nil, fmt.Errorf("readXptV89: invalid last modified date '%s'", lastModifiedDate)
	}
	offset += 80

	///////////////////////////////////////
	// 4	Member header records
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV89: invalid member header")
	}

	namestrSize, err := strconv.Atoi(string(content[offset+74 : offset+78]))
	if err != nil {
		return nil, fmt.Errorf("readXptV89: invalid NAMESTR size '%s'", string(content[offset+74:offset+78]))
	}
	offset += 80

	switch namestrSize {
	case 140:
		// TODO: read namestr
	default:
		return nil, fmt.Errorf("readXptV89: invalid NAMESTR size '%d'", namestrSize)
	}

	// skip the member header
	offset += 80

	///////////////////////////////////////
	// 5	Member header data
	dsName := string(content[offset+8 : offset+16])
	fileMeta.SasDsName = strings.Trim(dsName, " ")

	sasDataVersion := string(content[offset+24 : offset+32])
	fileMeta.SasDataVersion = strings.Trim(sasDataVersion, " ")

	// skip the member header data
	offset += 80
	// skip the header record
	offset += 80

	///////////////////////////////////////
	// 6	Namestr headerrecord
	var varsNum int
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV89: invalid namestr header")
	}

	// get number of variables
	n, err := strconv.ParseInt(string(content[offset+48:offset+58]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("readXptV89: invalid number of variables '%s'", string(content[offset+24:offset+32]))
	}
	varsNum = int(n)
	offset += 80

	///////////////////////////////////////
	// 7	Namestr records

	seriesMeta := make([]SeriesMeta, varsNum)
	namestrs := make([]__NAMESTRv89, varsNum)

	// read namestr
	for i := 0; i < varsNum; i++ {
		namestrs[i].FromBinary(content[offset:offset+140], byteOrder)
		type_ := meta.Float64Type
		if namestrs[i].ntype == 2 {
			type_ = meta.StringType
		}

		seriesMeta[i] = SeriesMeta{
			Name:   strings.Trim(string(namestrs[i].nname[:]), " "),
			Label:  strings.Trim(string(namestrs[i].nlabel[:]), " "),
			Length: int(namestrs[i].nlng),
			Type:   type_,
		}

		offset += namestrSize
	}

	// skip the padding
	if p := ((namestrSize * varsNum) % 80); p != 0 {
		offset += 80 - p
	}

	///////////////////////////////////////
	// 8	Observation header

	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, fmt.Errorf("readXptV89: invalid observation header")
	}

	// skip the observation header
	offset += 80

	///////////////////////////////////////
	// 9	Data records

	nulls := make([][]bool, varsNum)
	values := make([]interface{}, varsNum)

	for i := 0; i < varsNum; i++ {
		nulls[i] = make([]bool, 0)

		switch namestrs[i].ntype {
		case 1:
			values[i] = make([]float64, 0)
		case 2:
			values[i] = make([]string, 0)
		default:
			return nil, fmt.Errorf("readXptV89: invalid variable type '%d'", namestrs[i].ntype)
		}
	}

	// read observations by rows
	if maxObservations < 0 {
		maxObservations = math.MaxInt32
	}

	var tmp []byte
	rowCounter := 0
	for offset < len(content) && rowCounter < maxObservations {

		allNulls := true
		for i := offset; i < len(content); i++ {
			if content[i] != '\x20' {
				allNulls = false
				break
			}
		}

		if allNulls {
			break
		}

		rowLen := 0
		for i := 0; i < varsNum; i++ {
			tmp = make([]byte, namestrs[i].nlng)
			copy(tmp, content[offset+int(namestrs[i].npos):offset+int(namestrs[i].npos)+int(namestrs[i].nlng)])

			switch namestrs[i].ntype {

			// NUMERIC
			case 1:
				f, err := NewSasFloat(tmp).ToIeee(byteOrder)
				if err != nil {
					return nil, fmt.Errorf("readXptV89: %w", err)
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}

				values[i] = append(values[i].([]float64), f)

			// CHAR
			case 2:
				s := strings.Trim(string(tmp), " ")

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
		rowCounter++
	}

	_series := make([]series.Series, varsNum)
	for i := 0; i < varsNum; i++ {
		switch t := values[i].(type) {
		case []float64:
			_series[i] = series.NewSeriesFloat64(t, nulls[i], false, ctx)
		case []string:
			_series[i] = series.NewSeriesString(t, nulls[i], false, ctx)
		}
	}

	return &IoData{
		FileMeta:   fileMeta,
		SeriesMeta: seriesMeta,
		Series:     _series,
	}, nil
}

// This functions writes a SAS XPT file (versions 8/9).
func writeXPTv89(ioData *IoData, writer io.Writer, byteOrder binary.ByteOrder) error {

	const xptV89Template = "" +
		"HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000  " +
		"SAS     SAS     SASLIB  {{.SasLibVersion}}{{.SasOs}}                        {{.SasCreateDt}}" +
		"{{.SasCreateDt}}                                                                " +
		"HEADER RECORD*******MEMBER  HEADER RECORD!!!!!!!000000000000000001600000000140  " +
		"HEADER RECORD*******DSCRPTR HEADER RECORD!!!!!!!000000000000000000000000000000  " +
		"SAS     VALUES  SASDATA {{.SasDataVersion}}{{.SasOs}}                        {{.SasCreateDt}}" +
		"{{.SasCreateDt}}                                                                " +
		"HEADER RECORD*******NAMESTR HEADER RECORD!!!!!!!{{.VarsN}}00000000000000000000  "

	const xptV89ObsHeader = "" +
		"HEADER RECORD*******OBS     HEADER RECORD!!!!!!!000000000000000000000000000000  "

	type xptV89TemplateData struct {
		SasLibVersion     string
		SasDataVersion    string
		SasOs             string
		SasCreateDt       string
		SasLastModifiedDt string
		VarsN             string
	}

	tmpl, err := template.New("xptV89").Parse(xptV89Template)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, xptV89TemplateData{
		SasLibVersion:     "9.4     ",
		SasDataVersion:    "9.4     ",
		SasOs:             "X64_10HO",
		SasCreateDt:       formatDateTimeSAS(time.Now()),
		SasLastModifiedDt: formatDateTimeSAS(time.Now()),
		VarsN:             fmt.Sprintf("%010d", ioData.NCols()),
	})
	if err != nil {
		return err
	}

	offset := 0
	stringVarLengths := make([]int, ioData.NCols())

	var _series series.Series
	for i := 0; i < ioData.NCols(); i++ {
		_series = ioData.At(i)

		namestr := NewNamestrV89()
		namestr.npos = int32(offset)

		switch s := _series.(type) {
		case series.Bools:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case series.Ints:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case series.Int64s:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case series.Float64s:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case series.Strings:
			for _, v := range s.Data().([]string) {
				if len(v) > stringVarLengths[i] {
					stringVarLengths[i] = len(v)
				}
			}

			namestr.ntype = 2
			namestr.nlng = int16(stringVarLengths[i])
			offset += stringVarLengths[i]

		// TODO: implement
		// case preludiometa.TimeType:
		// 	namestr.ntype = 2
		// 	namestr.nlng = 0
		// 	offset += 0

		case series.Durations:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		default:
			return fmt.Errorf("writeXPTv89: invalid variable type '%v'", _series.Type())
		}

		namestr.nvar0 = int16(i + 1)
		copy(namestr.nname[:], []byte(fmt.Sprintf("%-8s", ioData.SeriesMeta[i].Name)[0:8])) // TODO: check if are repeated names
		// copy(namestr.nlabel[:], []byte(df.NameAt(i))[0:40]) // TODO: add labels to writer

		_, err = writer.Write(namestr.ToBinary(byteOrder))
		if err != nil {
			return err
		}
	}

	// add padding
	if p := ((140 * ioData.NCols()) % 80); p != 0 {
		_, err = writer.Write(bytes.Repeat([]byte{0x20}, 80-p))
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte(xptV89ObsHeader))
	if err != nil {
		return err
	}

	offset = 0
	for i := 0; i < ioData.NRows(); i++ {
		for j := 0; j < ioData.NCols(); j++ {
			_series = ioData.At(j)

			switch _series.(type) {

			// Numeric types
			case series.Bools, series.Ints, series.Int64s, series.Float64s:
				var val float64
				if _series.IsNull(i) {
					val = math.NaN()
				} else {
					switch s := _series.(type) {
					case series.Bools:
						val = 0
						if s.Get(i).(bool) {
							val = 1
						}
					case series.Ints:
						val = float64(s.Get(i).(int))
					case series.Int64s:
						val = float64(s.Get(i).(int64))
					case series.Float64s:
						val = s.Get(i).(float64)
					case series.Durations:
						val = float64(s.Get(i).(time.Duration))
					}
				}

				sf := NewSasFloat([]byte{})
				err = sf.FromIeee(val, byteOrder)
				if err != nil {
					return err
				}

				_, err = writer.Write([]byte(*sf))
				if err != nil {
					return err
				}

				offset += 8

			// String types
			case series.Strings:
				val := ""
				if !_series.IsNull(i) {
					val = _series.Get(i).(string)
				}

				_, err = writer.Write([]byte(fmt.Sprintf("%-*s", stringVarLengths[j], val)))
				if err != nil {
					return err
				}

				offset += stringVarLengths[j]

				// TODO: implement
				// case Times:
			}
		}
	}

	// add padding
	if p := (offset % 80); p != 0 {
		_, err = writer.Write(bytes.Repeat([]byte{0x20}, 80-p))
		if err != nil {
			return err
		}
	}

	return nil
}
