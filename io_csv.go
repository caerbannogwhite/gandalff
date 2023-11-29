package gandalff

import (
	"encoding/csv"
	"fmt"
	"os"

	"io"

	"preludiometa"
)

type CsvReader struct {
	header           bool
	delimiter        rune
	guessDataTypeLen int
	path             string
	nullValues       bool
	reader           io.Reader
	schema           *preludiometa.Schema
	ctx              *Context
}

func NewCsvReader(ctx *Context) *CsvReader {
	return &CsvReader{
		header:           CSV_READER_DEFAULT_HEADER,
		delimiter:        CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
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

func (r *CsvReader) SetSchema(schema *preludiometa.Schema) *CsvReader {
	r.schema = schema
	return r
}

func (r *CsvReader) SetContext(ctx *Context) *CsvReader {
	r.ctx = ctx
	return r
}

func (r *CsvReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: r.ctx}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("CsvReader: no reader specified"), ctx: r.ctx}
	}

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("CsvReader: no context specified"), ctx: r.ctx}
	}

	names, series, err := readCsv(r.reader, r.delimiter, r.header, r.nullValues, r.guessDataTypeLen, r.schema, r.ctx)
	if err != nil {
		return BaseDataFrame{err: err, ctx: r.ctx}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

// ReadCsv reads a CSV file and returns a GDLDataFrame.
func readCsv(reader io.Reader, delimiter rune, header bool, nullValues bool, guessDataTypeLen int, schema *preludiometa.Schema, ctx *Context) ([]string, []Series, error) {

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

	series, err := readRowData(csvReader, nullValues, guessDataTypeLen, schema, ctx)
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
	dataframe DataFrame
}

func NewCsvWriter() *CsvWriter {
	return &CsvWriter{
		delimiter: CSV_READER_DEFAULT_DELIMITER,
		header:    CSV_READER_DEFAULT_HEADER,
		format:    true,
		path:      "",
		naText:    NA_TEXT,
		writer:    nil,
		dataframe: nil,
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

func (w *CsvWriter) SetDataFrame(dataframe DataFrame) *CsvWriter {
	w.dataframe = dataframe
	return w
}

func (w *CsvWriter) Write() DataFrame {
	if w.dataframe == nil {
		return BaseDataFrame{err: fmt.Errorf("CsvWriter: no dataframe specified"), ctx: w.dataframe.GetContext()}
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
		return BaseDataFrame{err: fmt.Errorf("CsvWriter: no writer specified"), ctx: w.dataframe.GetContext()}
	}

	err := writeCsv(w.dataframe, w.writer, w.delimiter, w.header, w.format, w.naText)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeCsv(df DataFrame, writer io.Writer, delimiter rune, header bool, format bool, naText string) error {
	series := make([]Series, df.NCols())
	for i := 0; i < df.NCols(); i++ {
		series[i] = df.SeriesAt(i)
	}

	if header {
		for i, name := range df.Names() {
			if i > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}
			fmt.Fprintf(writer, "%s", name)
		}

		fmt.Fprintf(writer, "\n")
	}

	for i := 0; i < df.NRows(); i++ {
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
