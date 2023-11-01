package gandalff

import (
	"fmt"
	"io"
	"os"
)

type XlscReader struct {
	path   string
	reader io.Reader
	ctx    *Context
}

func NewXlscReader(ctx *Context) *XlscReader {
	return &XlscReader{
		path:   "",
		reader: nil,
		ctx:    ctx,
	}
}

func (r *XlscReader) SetPath(path string) *XlscReader {
	r.path = path
	return r
}

func (r *XlscReader) SetReader(reader io.Reader) *XlscReader {
	r.reader = reader
	return r
}

func (r *XlscReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("XlscReader: no reader specified")}
	}

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XlscReader: no context specified")}
	}

	names, series, err := readXlsx(r.reader, r.ctx)

	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

func readXlsx(reader io.Reader, ctx *Context) ([]string, []Series, error) {
	return nil, nil, nil
}

type XlsxWriter struct {
	path   string
	writer io.Writer
	ctx    *Context
}
