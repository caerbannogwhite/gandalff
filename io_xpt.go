package gandalff

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
)

// https://www.loc.gov/preservation/digital/formats/fdd/fdd000464.shtml

type XptVersionType uint8

const (
	XPT_VERSION_5 XptVersionType = iota + 5
	XPT_VERSION_6
	XPT_VERSION_8 XptVersionType = iota + 6
	XPT_VERSION_9
)

type XptReader struct {
	maxObservations int
	version         XptVersionType
	byteOrder       binary.ByteOrder
	path            string
	reader          io.Reader
	ctx             *Context
}

func NewXptReader(ctx *Context) *XptReader {
	return &XptReader{
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

func (r *XptReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: r.ctx}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("XptReader: no reader specified"), ctx: r.ctx}
	}

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XptReader: no context specified"), ctx: r.ctx}
	}

	var err error
	var names []string
	var series []Series

	switch r.version {
	case XPT_VERSION_5, XPT_VERSION_6:
		names, series, err = readXPTv56(r.reader, r.maxObservations, r.byteOrder, r.ctx)
	case XPT_VERSION_8, XPT_VERSION_9:
		names, series, err = readXPTv89(r.reader, r.maxObservations, r.byteOrder, r.ctx)
	default:
		return BaseDataFrame{err: fmt.Errorf("XptReader: unknown version"), ctx: r.ctx}
	}

	if err != nil {
		return BaseDataFrame{err: err, ctx: r.ctx}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

type XptWriter struct {
	version   XptVersionType
	byteOrder binary.ByteOrder
	path      string
	writer    io.Writer
	dataframe DataFrame
}

func NewXptWriter() *XptWriter {
	return &XptWriter{
		version:   XPT_VERSION_8,
		byteOrder: binary.BigEndian,
		path:      "",
		writer:    nil,
		dataframe: nil,
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

func (w *XptWriter) SetDataFrame(dataframe DataFrame) *XptWriter {
	w.dataframe = dataframe
	return w
}

func (w *XptWriter) Write() DataFrame {
	if w.dataframe == nil {
		return BaseDataFrame{err: fmt.Errorf("XptWriter: no dataframe specified"), ctx: w.dataframe.GetContext()}
	}

	if w.dataframe.IsErrored() {
		return w.dataframe
	}

	if w.path != "" {
		file, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		return BaseDataFrame{err: fmt.Errorf("XptWriter: no writer specified"), ctx: w.dataframe.GetContext()}
	}

	var err error
	switch w.version {
	case XPT_VERSION_5, XPT_VERSION_6:
		err = writeXPTv56(w.dataframe, w.writer, w.byteOrder)
	case XPT_VERSION_8, XPT_VERSION_9:
		err = writeXPTv89(w.dataframe, w.writer, w.byteOrder)
	default:
		return BaseDataFrame{err: fmt.Errorf("XptWriter: unknown SAS version '%d'", w.version), ctx: w.dataframe.GetContext()}
	}

	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
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
func readXPTv56(reader io.Reader, maxObservations int, byteOrder binary.ByteOrder, ctx *Context) ([]string, []Series, error) {
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
	switch strings.Split(version, ".")[0] {
	case "5", "6":

	default:
		return nil, nil, fmt.Errorf("readXPTv56: invalid version '%s'", version)
	}

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
	namestrs := make([]__NAMESTRv56, varsNum)

	// read namestr
	for i := 0; i < varsNum; i++ {
		namestrs[i].FromBinary(content[offset:offset+140], byteOrder)
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
					return nil, nil, err
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}

				// TODO: waiting for float32 support
				// if namestrs[i].nlng <= 4 {
				// 	values[i] = append(values[i].([]float32), float32(f))
				// } else {
				values[i] = append(values[i].([]float64), f)
				// }

			// CHAR
			case 2:
				s := string(tmp)

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
		rowCounter++
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
func writeXPTv56(df DataFrame, writer io.Writer, byteOrder binary.ByteOrder) error {
	// TODO: implement
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
func readXPTv89(reader io.Reader, maxObservations int, byteOrder binary.ByteOrder, ctx *Context) ([]string, []Series, error) {
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
	switch strings.Split(version, ".")[0] {
	case "8", "9":

	default:
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
		namestrs[i].FromBinary(content[offset:offset+140], byteOrder)
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
			// TODO: waiting for float32 support
			// if namestrs[i].nlng <= 4 {
			// 	values[i] = make([]float32, 0)
			// } else {
			values[i] = make([]float64, 0)
			// }
		case 2:
			values[i] = make([]string, 0)
		default:
			return nil, nil, fmt.Errorf("readXPTv89: invalid variable type '%d'", namestrs[i].ntype)
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
					return nil, nil, err
				}

				if math.IsNaN(f) {
					nulls[i] = append(nulls[i], true)
				} else {
					nulls[i] = append(nulls[i], false)
				}

				// TODO: waiting for float32 support
				// if namestrs[i].nlng <= 4 {
				// 	values[i] = append(values[i].([]float32), float32(f))
				// } else {
				values[i] = append(values[i].([]float64), f)
				// }

			// CHAR
			case 2:
				s := string(tmp)

				nulls[i] = append(nulls[i], false)
				values[i] = append(values[i].([]string), s)
			}
			rowLen += int(namestrs[i].nlng)
		}

		offset += rowLen
		rowCounter++
	}

	series := make([]Series, varsNum)
	for i := 0; i < varsNum; i++ {
		switch t := values[i].(type) {
		// TODO: waiting for float32 support
		// case []float32:
		// 	series[i] = NewSeriesFloat32(t, nulls[i], false, ctx)
		case []float64:
			series[i] = NewSeriesFloat64(t, nulls[i], false, ctx)
		case []string:
			series[i] = NewSeriesString(t, nulls[i], false, ctx)
		}
	}

	return names, series, nil
}

// This functions writes a SAS XPT file (versions 8/9).
func writeXPTv89(df DataFrame, writer io.Writer, byteOrder binary.ByteOrder) error {

	const xptV89Template = "" +
		"HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000  " +
		"SAS     SAS     SASLIB  {{.Vrs}}{{.Ops}}                        {{.SasCreateDt}}" +
		"{{.SasCreateDt}}                                                                " +
		"HEADER RECORD*******MEMBER  HEADER RECORD!!!!!!!000000000000000001600000000140  " +
		"HEADER RECORD*******DSCRPTR HEADER RECORD!!!!!!!000000000000000000000000000000  " +
		"SAS     VALUES  SASDATA {{.Vrs}}{{.Ops}}                        {{.SasCreateDt}}" +
		"{{.SasCreateDt}}                                                                " +
		"HEADER RECORD*******NAMESTR HEADER RECORD!!!!!!!{{.VarsN}}00000000000000000000  "

	const xptV89ObsHeader = "" +
		"HEADER RECORD*******OBS     HEADER RECORD!!!!!!!000000000000000000000000000000  "

	type xptV89TemplateData struct {
		Vrs         string
		Ops         string
		SasCreateDt string
		VarsN       string
	}

	tmpl, err := template.New("xptV89").Parse(xptV89Template)
	if err != nil {
		return err
	}

	err = tmpl.Execute(writer, xptV89TemplateData{
		Vrs:         "9.4     ",
		Ops:         "X64_10HO",
		SasCreateDt: formatDateTimeSAS(time.Now()),
		VarsN:       fmt.Sprintf("%010d", df.NCols()),
	})
	if err != nil {
		return err
	}

	offset := 0
	stringVarLengths := make([]int, df.NCols())

	var series Series
	for i := 0; i < df.NCols(); i++ {
		series = df.At(i)

		namestr := NewNamestrV89()
		namestr.npos = int32(offset)

		switch s := series.(type) {
		case SeriesBool:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case SeriesInt:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case SeriesInt64:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case SeriesFloat64:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		case SeriesString:
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

		case SeriesDuration:
			namestr.ntype = 1
			namestr.nlng = 8
			offset += 8

		default:
			return fmt.Errorf("writeXPTv89: invalid variable type '%v'", series.Type())
		}

		namestr.nvar0 = int16(i + 1)
		copy(namestr.nname[:], []byte(fmt.Sprintf("%-8s", df.NameAt(i))[0:8])) // TODO: check if are repeated names
		// copy(namestr.nlabel[:], []byte(df.NameAt(i))[0:40]) // TODO: add labels to writer

		_, err = writer.Write(namestr.ToBinary(byteOrder))
		if err != nil {
			return err
		}
	}

	// add padding
	if p := ((140 * df.NCols()) % 80); p != 0 {
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
	for i := 0; i < df.NRows(); i++ {
		for j := 0; j < df.NCols(); j++ {
			series = df.At(j)

			switch series.(type) {

			// Numeric types
			case SeriesBool, SeriesInt, SeriesInt64, SeriesFloat64:
				var val float64
				if series.IsNull(i) {
					val = math.NaN()
				} else {
					switch s := series.(type) {
					case SeriesBool:
						val = 0
						if s.Get(i).(bool) {
							val = 1
						}
					case SeriesInt:
						val = float64(s.Get(i).(int))
					case SeriesInt64:
						val = float64(s.Get(i).(int64))
					case SeriesFloat64:
						val = s.Get(i).(float64)
					case SeriesDuration:
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
			case SeriesString:
				val := ""
				if !series.IsNull(i) {
					val = series.Get(i).(string)
				}

				_, err = writer.Write([]byte(fmt.Sprintf("%-*s", stringVarLengths[j], val)))
				if err != nil {
					return err
				}

				offset += stringVarLengths[j]

				// TODO: implement
				// case SeriesTime:
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
