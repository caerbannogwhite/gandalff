package gandalff

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/caerbannogwhite/preludiometa"

	"github.com/tealeg/xlsx"
)

type XlsxReader struct {
	path             string
	sheet            string
	header           int
	rows             int
	guessDataTypeLen int
	nullValues       bool
	schema           *preludiometa.Schema
	ctx              *Context
}

func NewXlsxReader(ctx *Context) *XlsxReader {
	return &XlsxReader{
		path:             "",
		sheet:            "",
		header:           0,
		rows:             -1,
		guessDataTypeLen: XLSX_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
		nullValues:       false,
		schema:           nil,
		ctx:              ctx,
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

func (r *XlsxReader) SetGuessDataTypeLen(guessDataTypeLen int) *XlsxReader {
	r.guessDataTypeLen = guessDataTypeLen
	return r
}

func (r *XlsxReader) SetNullValues(nullValues bool) *XlsxReader {
	r.nullValues = nullValues
	return r
}

func (r *XlsxReader) SetSchema(schema *preludiometa.Schema) *XlsxReader {
	r.schema = schema
	return r
}

func (r *XlsxReader) Read() DataFrame {
	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("XlsxReader: no context specified"), ctx: r.ctx}
	}

	names, series, err := readXlsx(r.path, r.sheet, r.header, r.rows, r.nullValues, r.guessDataTypeLen, r.schema, r.ctx)

	if err != nil {
		return BaseDataFrame{err: err, ctx: r.ctx}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

type xlsxRowReader struct {
	sh    *xlsx.Sheet
	row   int
	cols  int
	cells []*xlsx.Cell
}

func (r *xlsxRowReader) Read() ([]string, error) {
	if r.row >= r.sh.MaxRow {
		return nil, io.EOF
	}

	row := r.sh.Row(r.row)
	r.cells = row.Cells
	r.row++

	values := make([]string, len(r.cells))
	for i, cell := range r.cells {
		values[i] = cell.String()
	}

	if len(values) < r.cols {
		for i := len(values); i < r.cols; i++ {
			values = append(values, "")
		}
	}

	return values, nil
}

func readXlsx(
	path string, sheet string, header, rows int, nullValues bool,
	guessDataTypeLen int, schema *preludiometa.Schema, ctx *Context,
) ([]string, []Series, error) {
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

	names := make([]string, len(sh.Row(header).Cells))
	for i, cell := range sh.Row(header).Cells {
		names[i] = cell.String()
	}

	xlsxRowReader := &xlsxRowReader{
		sh:    sh,
		row:   header + 1,
		cols:  len(names),
		cells: nil,
	}

	series, err := readRowData(xlsxRowReader, nullValues, guessDataTypeLen, sh.MaxRow, schema, ctx)
	if err != nil {
		return nil, nil, err
	}

	return names, series, nil
}

type XlsxWriter struct {
	path      string
	sheet     string
	naText    string
	writer    io.Writer
	dataframe DataFrame
}

func NewXlsxWriter() *XlsxWriter {
	return &XlsxWriter{
		path:      "",
		sheet:     "Sheet1",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *XlsxWriter) SetPath(path string) *XlsxWriter {
	w.path = path
	return w
}

func (w *XlsxWriter) SetSheet(sheet string) *XlsxWriter {
	w.sheet = sheet
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
		return BaseDataFrame{err: fmt.Errorf("XlsxWriter: no dataframe specified"), ctx: w.dataframe.GetContext()}
	}

	if w.dataframe.IsErrored() {
		return w.dataframe
	}

	if w.path != "" {
		// make sure os.O_WRONLY arg is supplies so that file is not opened in read-only mode
		file, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		return BaseDataFrame{err: fmt.Errorf("XlsxWriter: no writer specified"), ctx: w.dataframe.GetContext()}
	}

	err := writeXlsx(w.dataframe, w.writer, w.sheet, w.naText)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeXlsx(dataframe DataFrame, writer io.Writer, sheetName string, naText string) error {

	file := xlsx.NewFile()

	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		return err
	}

	// write header
	row := sheet.AddRow()
	for _, name := range dataframe.Names() {
		cell := row.AddCell()
		cell.Value = name
	}

	// write data
	for i := 0; i < dataframe.NRows(); i++ {
		row := sheet.AddRow()
		for j := range dataframe.Names() {
			cell := row.AddCell()

			switch s := dataframe.At(j).(type) {
			case SeriesBool:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetBool(s.Get(i).(bool))

			case SeriesInt:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetInt(s.Get(i).(int))

			case SeriesInt64:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetInt64(s.Get(i).(int64))

			case SeriesFloat64:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetFloat(s.Get(i).(float64))

			case SeriesString:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.Value = s.Get(i).(string)

			case SeriesTime:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetDateTime(s.Get(i).(time.Time))

			case SeriesDuration:
				if s.IsNull(i) {
					cell.Value = naText
					continue
				}
				cell.SetInt(int(s.Get(i).(time.Duration).Nanoseconds()))

			default:
				cell.Value = naText
			}
		}
	}

	err = file.Write(writer)
	if err != nil {
		return err
	}

	return nil
}
