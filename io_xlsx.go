package gandalff

import (
	"fmt"
	"io"
	"os"

	"github.com/tealeg/xlsx"
)

type XlscReader struct {
	path  string
	sheet string
	ctx   *Context
}

func NewXlscReader(ctx *Context) *XlscReader {
	return &XlscReader{
		path:  "",
		sheet: "",
		ctx:   ctx,
	}
}

func (r *XlscReader) SetPath(path string) *XlscReader {
	r.path = path
	return r
}

func (r *XlscReader) SetSheet(sheet string) *XlscReader {
	r.sheet = sheet
	return r
}

func (r *XlscReader) Read() DataFrame {
	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XlscReader: no context specified")}
	}

	names, series, err := readXlsx(r.path, r.sheet, r.ctx)

	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

func readXlsx(path string, sheet string, ctx *Context) ([]string, []Series, error) {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}

	sh, ok := wb.Sheet[sheet]
	if !ok {
		return nil, nil, fmt.Errorf("Sheet %s not found", sheet)
	}

	// sh.ForEachRow(rowVisitor)
	return nil, nil, nil
}

type XlsxWriter struct {
	path      string
	writer    io.Writer
	dataframe DataFrame
}

func NewXlsxWriter() *XlsxWriter {
	return &XlsxWriter{
		path:      "",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *XlsxWriter) SetPath(path string) *XlsxWriter {
	w.path = path
	return w
}

func (w *XlsxWriter) SetWriter(writer io.Writer) *XlsxWriter {
	w.writer = writer
	return w
}

func (w *XlsxWriter) SetDataFrame(dataframe DataFrame) *XlsxWriter {
	w.dataframe = dataframe
	return w
}

func (w *XlsxWriter) Write() DataFrame {
	if w.dataframe == nil {
		return BaseDataFrame{err: fmt.Errorf("XlsxWriter: no dataframe specified")}
	}

	if w.dataframe.IsErrored() {
		return w.dataframe
	}

	if w.path != "" {
		file, err := os.OpenFile(w.path, os.O_CREATE, 0666)
		if err != nil {
			return BaseDataFrame{err: err}
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		return BaseDataFrame{err: fmt.Errorf("XlsxWriter: no writer specified")}
	}

	err := writeXlsx(w.dataframe, w.writer)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err}
	}

	return w.dataframe
}

func writeXlsx(dataframe DataFrame, writer io.Writer) error {
	return nil
}
