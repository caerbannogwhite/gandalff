package gandalff

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
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

const (
	valueHeaderStart = "HEADER RECORD*******"
	valueSAS         = "SAS     "
	valueLIB         = "SASLIB  "
)

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
func readXPTv56(reader io.Reader, ctx *Context) ([]string, []Series, error) {
	if ctx == nil {
		return nil, nil, fmt.Errorf("readXPTv56: no context specified")
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
	// 		HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000
	if string(content[0:20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv56: invalid header")
	}
	offset += 80

	///////////////////////////////////////
	// 2	The first real header record
	if string(content[offset:offset+8]) != valueSAS ||
		string(content[offset+8:offset+16]) != valueSAS ||
		string(content[offset+16:offset+24]) != valueLIB {
		return nil, nil, fmt.Errorf("readXPTv56: invalid first real header")
	}

	version := strings.Trim(string(content[offset+24:offset+32]), " ")
	if version != "" { // TODO: check version
		return nil, nil, fmt.Errorf("readXPTv56: invalid version '%s'", version)
	}
	offset += 80

	///////////////////////////////////////
	// 3	Second real header record: ddMMMyy:hh:mm:ss
	offset += 80

	///////////////////////////////////////
	// 4	Member header records
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv56: invalid member header")
	}

	namestrSize := 140
	offset += 80

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
	var varsNum int
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv56: invalid namestr header")
	}

	// get number of variables
	n, err := strconv.ParseInt(string(content[offset+48:offset+58]), 10, 32)
	if err != nil {
		return nil, nil, fmt.Errorf("readXPTv56: invalid number of variables '%s'", string(content[offset+24:offset+32]))
	}
	varsNum = int(n)
	offset += 80

	///////////////////////////////////////
	// 7	Namestr records

	names := make([]string, varsNum)
	namestrs := make([]__NAMESTRv89, varsNum)

	// read namestr
	for i := 0; i < varsNum; i++ {

		// ntype  int16    // VARIABLE TYPE: 1=NUMERIC, 2=CHAR 		(bytes: 000 to 002)
		// nhfun  int16    // HASH OF NNAME (always 0)				(bytes: 002 to 004)
		// nlng   int16    // LENGTH OF VARIABLE IN OBSERVATION		(bytes: 004 to 006)
		// nvar0  int16    // VARNUM								(bytes: 006 to 008)
		// nname  [8]byte  // NAME OF VARIABLE						(bytes: 008 to 016)
		// nlabel [40]byte // LABEL OF VARIABLE						(bytes: 016 to 056)
		// nform  [8]byte  // NAME OF FORMAT						(bytes: 056 to 064)
		// nfl    int16    // FORMAT FIELD LENGTH OR 0				(bytes: 064 to 066)
		// nfd    int16    // FORMAT NUMBER OF DECIMALS				(bytes: 066 to 068)
		// nfj    int16    // 0=LEFT JUSTIFICATION, 1=RIGHT JUST	(bytes: 068 to 070)
		// nfill  [2]byte  // (UNUSED, FOR ALIGNMENT AND FUTURE)	(bytes: 070 to 072)
		// niform [8]byte  // NAME OF INPUT FORMAT					(bytes: 072 to 080)
		// nifl   int16    // INFORMAT LENGTH ATTRIBUTE				(bytes: 080 to 082)
		// nifd   int16    // INFORMAT NUMBER OF DECIMALS			(bytes: 082 to 084)
		// npos   int32    // POSITION OF VALUE IN OBSERVATION		(bytes: 084 to 088)
		// rest   [52]byte // remaining fields are irrelevant		(bytes: 088 to 140)

		namestrs[i].ntype = int16(binary.BigEndian.Uint16(content[offset : offset+2]))
		// namestrs[i].nhfun = int16(binary.BigEndian.Uint16(content[offset+2 : offset+4]))
		namestrs[i].nlng = int16(binary.BigEndian.Uint16(content[offset+4 : offset+6]))
		namestrs[i].nvar0 = int16(binary.BigEndian.Uint16(content[offset+6 : offset+8]))
		copy(namestrs[i].nname[:], content[offset+8:offset+16])
		copy(namestrs[i].nlabel[:], content[offset+16:offset+56])
		// copy(namestrs[i].nform[:], content[offset+56:offset+64])
		// namestrs[i].nfl = int16(binary.BigEndian.Uint16(content[offset+64 : offset+66]))
		// namestrs[i].nfd = int16(binary.BigEndian.Uint16(content[offset+66 : offset+68]))
		// namestrs[i].nfj = int16(binary.BigEndian.Uint16(content[offset+68 : offset+70]))
		// copy(namestrs[i].nfill[:], content[offset+70:offset+72])
		// copy(namestrs[i].niform[:], content[offset+72:offset+80])
		// namestrs[i].nifl = int16(binary.BigEndian.Uint16(content[offset+80 : offset+82]))
		// namestrs[i].nifd = int16(binary.BigEndian.Uint16(content[offset+82 : offset+84]))
		namestrs[i].npos = int32(binary.BigEndian.Uint32(content[offset+84 : offset+88]))
		// copy(namestrs[i].rest[:], content[offset+88:offset+140])

		names[i] = strings.Trim(string(namestrs[i].nname[:]), " ")
		offset += namestrSize
	}

	// skip the padding
	if p := ((namestrSize * varsNum) % 80); p != 0 {
		offset += 80 - p
	}

	///////////////////////////////////////
	// 8	Observation header

	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv56: invalid observation header")
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
			return nil, nil, fmt.Errorf("readXPTv56: invalid variable type '%d'", namestrs[i].ntype)
		}
	}

	// read observations by rows
	for offset < len(content) {

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
			buffer := content[offset+int(namestrs[i].npos) : offset+int(namestrs[i].npos)+int(namestrs[i].nlng)]

			switch namestrs[i].ntype {
			case 1:
				f, err := NewSasFloat(buffer).ToIeee()
				if err != nil {
					return nil, nil, err
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}
				values[i] = append(values[i].([]float64), f)

			case 2:
				s := string(buffer)

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
	}

	series := make([]Series, varsNum)
	for i := 0; i < varsNum; i++ {
		switch t := values[i].(type) {
		case []float64:
			series[i] = NewSeriesFloat64(t, nulls[i], false, ctx)
		case []string:
			series[i] = NewSeriesString(t, nulls[i], false, ctx)
		}
	}

	return names, series, nil
}

// This functions writes a SAS XPT file (versions 5/6).
func writeXPTv56(path string) error {
	buff := make([]byte, 0)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
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

func (nms *__NAMESTRv89) ToString() string {
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
	if string(content[0:20]) != valueHeaderStart {
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
	if string(content[offset:offset+20]) != valueHeaderStart {
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
	var varsNum int
	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv89: invalid namestr header")
	}

	// get number of variables
	n, err := strconv.ParseInt(string(content[offset+48:offset+58]), 10, 32)
	if err != nil {
		return nil, nil, fmt.Errorf("readXPTv89: invalid number of variables '%s'", string(content[offset+24:offset+32]))
	}
	varsNum = int(n)
	offset += 80

	///////////////////////////////////////
	// 7	Namestr records

	names := make([]string, varsNum)
	namestrs := make([]__NAMESTRv89, varsNum)

	// read namestr
	for i := 0; i < varsNum; i++ {

		// ntype    int16    // VARIABLE TYPE: 1=NUMERIC, 2=CHAR	(bytes: 000 to 002)
		// nhfun    int16    // HASH OF NNAME (always 0)			(bytes: 002 to 004)
		// nlng     int16    // LENGTH OF VARIABLE IN OBSERVATION	(bytes: 004 to 006)
		// nvar0    int16    // VARNUM								(bytes: 006 to 008)
		// nname    [8]byte  // NAME OF VARIABLE					(bytes: 008 to 016)
		// nlabel   [40]byte // LABEL OF VARIABLE					(bytes: 016 to 056)
		// nform    [8]byte  // NAME OF FORMAT						(bytes: 056 to 064)
		// nfl      int16    // FORMAT FIELD LENGTH OR 0			(bytes: 064 to 066)
		// nfd      int16    // FORMAT NUMBER OF DECIMALS			(bytes: 066 to 068)
		// nfj      int16    // 0=LEFT JUSTIFICATION, 1=RIGHT JUST	(bytes: 068 to 070)
		// nfill    [2]byte  // (UNUSED, FOR ALIGNMENT AND FUTURE)	(bytes: 070 to 072)
		// niform   [8]byte  // NAME OF INPUT FORMAT				(bytes: 072 to 080)
		// nifl     int16    // INFORMAT LENGTH ATTRIBUTE			(bytes: 080 to 082)
		// nifd     int16    // INFORMAT NUMBER OF DECIMALS			(bytes: 082 to 084)
		// npos     int32    // POSITION OF VALUE IN OBSERVATION	(bytes: 084 to 088)
		// longname [32]byte // long name for Version 8-style		(bytes: 088 to 120)
		// lablen   int16    // length of label						(bytes: 120 to 122)
		// rest     [18]byte // remaining fields are irrelevant		(bytes: 122 to 140)

		namestrs[i].ntype = int16(binary.BigEndian.Uint16(content[offset : offset+2]))
		// namestrs[i].nhfun = int16(binary.BigEndian.Uint16(content[offset+2 : offset+4]))
		namestrs[i].nlng = int16(binary.BigEndian.Uint16(content[offset+4 : offset+6]))
		namestrs[i].nvar0 = int16(binary.BigEndian.Uint16(content[offset+6 : offset+8]))
		copy(namestrs[i].nname[:], content[offset+8:offset+16])
		copy(namestrs[i].nlabel[:], content[offset+16:offset+56])
		// copy(namestrs[i].nform[:], content[offset+56:offset+64])
		// namestrs[i].nfl = int16(binary.BigEndian.Uint16(content[offset+64 : offset+66]))
		// namestrs[i].nfd = int16(binary.BigEndian.Uint16(content[offset+66 : offset+68]))
		// namestrs[i].nfj = int16(binary.BigEndian.Uint16(content[offset+68 : offset+70]))
		// copy(namestrs[i].niform[:], content[offset+72:offset+80])
		// namestrs[i].nifl = int16(binary.BigEndian.Uint16(content[offset+80 : offset+82]))
		// namestrs[i].nifd = int16(binary.BigEndian.Uint16(content[offset+82 : offset+84]))
		namestrs[i].npos = int32(binary.BigEndian.Uint32(content[offset+84 : offset+88]))
		copy(namestrs[i].longname[:], content[offset+88:offset+120])
		namestrs[i].lablen = int16(binary.BigEndian.Uint16(content[offset+122 : offset+124]))

		names[i] = strings.Trim(string(namestrs[i].nname[:]), " ")
		offset += namestrSize
	}

	// skip the padding
	if p := ((namestrSize * varsNum) % 80); p != 0 {
		offset += 80 - p
	}

	///////////////////////////////////////
	// 8	Observation header

	if string(content[offset:offset+20]) != valueHeaderStart {
		return nil, nil, fmt.Errorf("readXPTv89: invalid observation header")
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
			return nil, nil, fmt.Errorf("readXPTv89: invalid variable type '%d'", namestrs[i].ntype)
		}
	}

	// read observations by rows
	for offset < len(content) {

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
			buffer := content[offset+int(namestrs[i].npos) : offset+int(namestrs[i].npos)+int(namestrs[i].nlng)]

			switch namestrs[i].ntype {
			case 1:
				f, err := NewSasFloat(buffer).ToIeee()
				if err != nil {
					return nil, nil, err
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}
				values[i] = append(values[i].([]float64), f)

			case 2:
				s := string(buffer)

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
	}

	series := make([]Series, varsNum)
	for i := 0; i < varsNum; i++ {
		switch t := values[i].(type) {
		case []float64:
			series[i] = NewSeriesFloat64(t, nulls[i], false, ctx)
		case []string:
			series[i] = NewSeriesString(t, nulls[i], false, ctx)
		}
	}

	return names, series, nil
}

// This functions writes a SAS XPT file (versions 8/9).
func writeXPTv89(path string) error {
	const firstHeaderRecord = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000  "

	var osVersion string
	switch runtime.GOOS {
	// case "darwin":
	// 	osVersion = "MacOS"
	// case "linux":
	// 	osVersion = "Linux"
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
