package gandalff

import (
	"fmt"
	"io"
	"os"
	"runtime"
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

///////////////////////////////////////     SAS XPT v5/6     ///////////////////////////////////////

// Technical documentation:
// https://support.sas.com/content/dam/SAS/support/en/technical-papers/record-layout-of-a-sas-version-5-or-6-data-set-in-sas-transport-xport-format.pdf
const FIRST_HEADER_V56 = "HEADER RECORD*******LIBRARY HEADER RECORD!!!!!!!000000000000000000000000000000"

// This functions writes a SAS XPT file (versions 5/6).
func readXPTv56(reader io.Reader, ctx *Context) ([]string, []Series, error) {

	if ctx == nil {
		return nil, nil, fmt.Errorf("readCSV: no context specified")
	}

	var err error
	var buff []byte

	_, err = reader.Read(buff)
	if err != nil {
		return nil, nil, err
	}

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

const FIRST_HEADER_V89 = "HEADER RECORD*******LIBV8 HEADER RECORD!!!!!!!000000000000000000000000000000"

// This functions writes a SAS XPT file (versions 7/8).
func readXPTv89(reader io.Reader, ctx *Context) ([]string, []Series, error) {
	return nil, nil, nil
}

// This functions writes a SAS XPT file (versions 7/8).
func writeXPTv89(path string) error {

	buff := make([]byte, 0)

	buff = append(buff, []byte(fmt.Sprintf(
		"%s%8s%8s%8s%8s%24s%16s%80s",
		FIRST_HEADER_V89,                      // 1-80 		First header record
		"SAS",                                 // 81-88 	SAS
		"SAS",                                 // 89-96 	SAS
		"SASLIB  9.4",                         // 97-104 	SASLIB
		runtime.GOOS,                          // 105-128	OS Name
		"",                                    // 129-152 	24 blanks
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 153-176   Date/time created
		time.Now().Format("ddMMMyy:hh:mm:ss"), // 177-200 	Second header record, date/time modified
	))...)

	// write buff to file
	os.WriteFile(path, buff, 0644)

	return nil
}
