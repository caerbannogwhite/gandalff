package io

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"io"

	"github.com/caerbannogwhite/aargh"
	"github.com/caerbannogwhite/aargh/meta"
	"github.com/caerbannogwhite/aargh/series"
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
	ctx              *aargh.Context
}

func NewCsvReader(ctx *aargh.Context) *CsvReader {
	return &CsvReader{
		header:           aargh.CSV_READER_DEFAULT_HEADER,
		rows:             -1,
		delimiter:        aargh.CSV_READER_DEFAULT_DELIMITER,
		guessDataTypeLen: aargh.CSV_READER_DEFAULT_GUESS_DATA_TYPE_LEN,
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

func (r *CsvReader) SetContext(ctx *aargh.Context) *CsvReader {
	r.ctx = ctx
	return r
}

func (r *CsvReader) Read() *IoData {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return &IoData{Error: err}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return &IoData{Error: fmt.Errorf("CsvReader: no reader specified")}
	}

	if r.ctx == nil {
		return &IoData{Error: fmt.Errorf("CsvReader: no context specified")}
	}

	names, series, err := readCsv(r.reader, r.delimiter, r.header, r.rows, r.nullValues, r.guessDataTypeLen, r.schema, r.ctx)
	if err != nil {
		return &IoData{Error: err}
	}

	iod := IoData{
		FileMeta: FileMeta{
			FileName: r.path,
			FilePath: r.path,
		},
		ctx: r.ctx,
	}

	for i, name := range names {
		iod.AddSeries(series[i], SeriesMeta{
			Name: name,
		})
	}

	return &iod
}

// ReadCsv reads a CSV file and returns a GDLDataFrame.
func readCsv(
	reader io.Reader, delimiter rune, header bool, rows int, nullValues bool,
	guessDataTypeLen int, schema *meta.Schema, ctx *aargh.Context,
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

type CsvQuotingType int

const (
	CsvQuotingNone CsvQuotingType = iota
	CsvQuotingAll
	CsvQuotingNeeded
	CsvQuotingNonNumeric
)

type CsvWriter struct {
	delimiter rune
	header    bool
	format    bool // TODO: Implement this
	path      string
	naText    string
	eol       string
	quote     string
	quoting   CsvQuotingType
	writer    io.Writer
	ioData    *IoData
}

func NewCsvWriter() *CsvWriter {
	return &CsvWriter{
		delimiter: aargh.CSV_READER_DEFAULT_DELIMITER,
		header:    aargh.CSV_READER_DEFAULT_HEADER,
		format:    true,
		path:      "",
		naText:    aargh.NA_TEXT,
		eol:       aargh.EOL,
		quote:     aargh.QUOTE,
		quoting:   CsvQuotingNeeded,
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

func (w *CsvWriter) SetEol(eol string) *CsvWriter {
	w.eol = eol
	return w
}

func (w *CsvWriter) SetQuote(quote string) *CsvWriter {
	w.quote = quote
	return w
}

func (w *CsvWriter) SetQuoting(quoting CsvQuotingType) *CsvWriter {
	w.quoting = quoting
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

	if w.ioData.Error != nil {
		return w.ioData.Error
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

	err := writeCsv(w.ioData, w.writer, w.delimiter, w.header, w.format, w.naText, w.eol, w.quote, w.quoting)
	if err != nil {
		return err
	}

	return nil
}

func writeCsv(ioData *IoData, writer io.Writer, delimiter rune, header bool, format bool, naText string, eol string, quote string, quoting CsvQuotingType) error {
	if header {
		for i, meta := range ioData.SeriesMeta {
			if i > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}
			fmt.Fprintf(writer, "%s", meta.Name)
		}

		fmt.Fprintf(writer, "%s", eol)
	}

	for i := 0; i < ioData.NRows(); i++ {
		for j, s := range ioData.Series {
			if j > 0 {
				fmt.Fprintf(writer, "%c", delimiter)
			}

			if s.IsNull(i) {
				fmt.Fprintf(writer, "%s", naText)
			} else {
				switch quoting {
				case CsvQuotingNone:
					fmt.Fprintf(writer, "%s", s.GetAsString(i))

				case CsvQuotingAll:
					fmt.Fprintf(writer, "%s", quote+s.GetAsString(i)+quote)

				case CsvQuotingNeeded:
					str := s.GetAsString(i)
					if strings.Contains(str, quote) || strings.Contains(str, string(delimiter)) || strings.Contains(str, eol) {
						fmt.Fprintf(writer, "%s", quote+str+quote)
					} else {
						fmt.Fprintf(writer, "%s", str)
					}

				case CsvQuotingNonNumeric:
					if s.Type() == meta.Float64Type || s.Type() == meta.Int64Type || s.Type() == meta.IntType {
						fmt.Fprintf(writer, "%s", s.GetAsString(i))
					} else {
						fmt.Fprintf(writer, "%s", quote+s.GetAsString(i)+quote)
					}

				default:
					return fmt.Errorf("writeCsv: invalid quoting type")
				}
			}
		}

		fmt.Fprintf(writer, "%s", eol)
	}

	return nil
}
