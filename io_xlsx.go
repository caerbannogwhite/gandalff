package gandalff

import (
	"fmt"
	"io"
	"os"

	"github.com/tealeg/xlsx"
)

type XlsxReader struct {
	path   string
	sheet  string
	header int
	rows   int
	ctx    *Context
}

func NewXlsxReader(ctx *Context) *XlsxReader {
	return &XlsxReader{
		path:   "",
		sheet:  "",
		header: 0,
		rows:   -1,
		ctx:    ctx,
	}
}

func (r *XlsxReader) SetPath(path string) *XlsxReader {
	r.path = path
	return r
}

func (r *XlsxReader) SetSheet(sheet string) *XlsxReader {
	r.sheet = sheet
	return r
}

func (r *XlsxReader) SetHeader(header int) *XlsxReader {
	r.header = header
	return r
}

func (r *XlsxReader) SetRows(rows int) *XlsxReader {
	r.rows = rows
	return r
}

func (r *XlsxReader) Read() DataFrame {
	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XlsxReader: no context specified")}
	}

	names, series, err := readXlsx(r.path, r.sheet, r.header, r.rows, r.ctx)

	if err != nil {
		return BaseDataFrame{err: err}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

func readXlsx(path string, sheet string, header int, rows int, ctx *Context) ([]string, []Series, error) {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		return nil, nil, err
	}

	sh, ok := wb.Sheet[sheet]
	if !ok {
		return nil, nil, fmt.Errorf("Sheet %s not found", sheet)
	}

	if rows < 0 {
		rows = sh.MaxRow
	}

	for i := header; i < header+rows; i++ {
		row := sh.Row(i)
		for _, cell := range row.Cells {
			fmt.Printf("%s\n", cell.String())
		}
	}

	return nil, nil, nil
}

type XlsxWriter struct {
	path      string
	naText    string
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

func (w *XlsxWriter) SetNaText(naText string) *XlsxWriter {
	w.naText = naText
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
