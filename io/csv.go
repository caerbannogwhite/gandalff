package io

import (
	"encoding/csv"
	"fmt"
	"os"

	"io"

	"github.com/caerbannogwhite/gandalff"
	"github.com/caerbannogwhite/gandalff/meta"
	"github.com/caerbannogwhite/gandalff/series"
)

type CsvReader struct {
	header           bool
	rows             int
	delimiter        rune
	guessDataTypeLen int
	path             string
	nullValues       bool
	reader           io.Reader
	schema           *meta.Schema
	ctx              *gandalff.Context
}

func NewCsvReader(ctx *gandalff.Context) *CsvReader {
	return &CsvReader{
		header:           gandalff.CSV_READER_DEFAULT_HEADER,
		rows:             -1,
		delimiter:        gandalff.CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: gandalff.CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
		path:             "",
		nullValues:       false,
		reader:           nil,
		schema:           nil,
		ctx:              ctx,
	}
}

func (r *CsvReader) SetHeader(header bool) *CsvReader {
	r.header = header
	return r
}

func (r *CsvReader) SetDelimiter(delimiter rune) *CsvReader {
	r.delimiter = delimiter
	return r
}

func (r *CsvReader) SetGuessDataTypeLen(guessDataTypeLen int) *CsvReader {
	r.guessDataTypeLen = guessDataTypeLen
	return r
}

func (r *CsvReader) SetRows(rows int) *CsvReader {
	r.rows = rows
	return r
}

func (r *CsvReader) SetPath(path string) *CsvReader {
	r.path = path
	return r
}

func (r *CsvReader) SetNullValues(nullValues bool) *CsvReader {
	r.nullValues = nullValues
	return r
}

func (r *CsvReader) SetReader(reader io.Reader) *CsvReader {
	r.reader = reader
	return r
}

func (r *CsvReader) SetSchema(schema *meta.Schema) *CsvReader {
	r.schema = schema
	return r
}

func (r *CsvReader) SetContext(ctx *gandalff.Context) *CsvReader {
	r.ctx = ctx
	return r
}

func (r *CsvReader) Read() (*IoData, error) {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return nil, fmt.Errorf("CsvReader: no reader specified")
	}

	if r.ctx == nil {
		return nil, fmt.Errorf("CsvReader: no context specified")
	}

	names, series, err := readCsv(r.reader, r.delimiter, r.header, r.rows, r.nullValues, r.guessDataTypeLen, r.schema, r.ctx)
	if err != nil {
		return nil, err
	}

	iod := IoData{
		FileMeta: FileMeta{
			FileName: r.path,
			FilePath: r.path,
		},
	}

	for i, name := range names {
		iod.AddSeries(series[i], SeriesMeta{
			Name: name,
		})
	}

	return &iod, nil
}

// ReadCsv reads a CSV file and returns a GDLDataFrame.
func readCsv(
	reader io.Reader, delimiter rune, header bool, rows int, nullValues bool,
	guessDataTypeLen int, schema *meta.Schema, ctx *gandalff.Context,
) ([]string, []series.Series, error) {

	// TODO: Add support for Time and Duration types (defined in a schema)
	// TODO: Optimize null masks (use bit vectors)?
	// TODO: Try to optimize this function by using goroutines: read the rows (like 1000)
	//		and guess the data types in parallel

	if ctx == nil {
		return nil, nil, fmt.Errorf("readCsv: no context specified")
	}

	// Initialize CSV reader
	csvReader := csv.NewReader(reader)
	csvReader.Comma = delimiter
	csvReader.FieldsPerRecord = -1

	// Read header if present
	var names []string
	var err error
	if header {
		names, err = csvReader.Read()
		if err != nil {
			return nil, nil, err
		}
	}

	series, err := readRowData(csvReader, nullValues, guessDataTypeLen, rows, schema, ctx)
	if err != nil {
		return nil, nil, err
	}

	// Generate names if not present
	if !header {
		for i := 0; i < len(series); i++ {
			names = append(names, fmt.Sprintf("Column %d", i+1))
		}
	}

	return names, series, nil
}

type CsvWriter struct {
	delimiter rune
	header    bool
	format    bool // TODO: Implement this
	path      string
	naText    string
	writer    io.Writer
	ioData    *IoData
}

func NewCsvWriter() *CsvWriter {
	return &CsvWriter{
		delimiter: gandalff.CSV_READER_DEFAULT_DELIMITER,
		header:    gandalff.CSV_READER_DEFAULT_HEADER,
		format:    true,
		path:      "",
		naText:    gandalff.NA_TEXT,
		writer:    nil,
		ioData:    nil,
	}
}

func (w *CsvWriter) SetDelimiter(delimiter rune) *CsvWriter {
	w.delimiter = delimiter
	return w
}

func (w *CsvWriter) SetHeader(header bool) *CsvWriter {
	w.header = header
	return w
}

func (w *CsvWriter) SetFormat(format bool) *CsvWriter {
	w.format = format
	return w
}

func (w *CsvWriter) SetPath(path string) *CsvWriter {
	w.path = path
	return w
}

func (w *CsvWriter) SetNaText(naText string) *CsvWriter {
	w.naText = naText
	return w
}

func (w *CsvWriter) SetWriter(writer io.Writer) *CsvWriter {
	w.writer = writer
	return w
}

func (w *CsvWriter) SetIoData(ioData *IoData) *CsvWriter {
	w.ioData = ioData
	return w
}

func (w *CsvWriter) Write() error {
	if w.ioData == nil {
		return fmt.Errorf("CsvWriter: no ioData specified")
	}

	if w.path != "" {
		file, err := os.OpenFile(w.path, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return err
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		return fmt.Errorf("CsvWriter: no writer specified")
	}

	err := writeCsv(w.ioData, w.writer, w.delimiter, w.header, w.format, w.naText)
	if err != nil {
		return err
	}

	return nil
}

func writeCsv(ioData *IoData, writer io.Writer, delimiter rune, header bool, format bool, naText string) error {
	series := make([]series.Series, len(ioData.Series))
	for i := 0; i < len(ioData.Series); i++ {
		series[i] = ioData.Series[i]
	}

	if header {
		for i, name := range ioData.SeriesMeta {
			if i > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}
			fmt.Fprintf(writer, "%s", name)
		}

		fmt.Fprintf(writer, "\n")
	}

	for i := 0; i < ioData.NRows(); i++ {
		for j, s := range series {
			if j > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}

			if s.IsNull(i) {
				fmt.Fprintf(writer, "%s", naText)
			} else {
				fmt.Fprintf(writer, "%s", s.GetAsString(i))
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	return nil
}
