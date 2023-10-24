package gandalff

import (
	"io"
)

type MarkDownWriter struct {
	header     bool
	types      bool
	path       string
	nullString string
	writer     io.Writer
	dataframe  DataFrame
}

func NewMarkDownWriter() *MarkDownWriter {
	return &MarkDownWriter{
		header:     CSV_READER_DEFAULT_HEADER,
		path:       "",
		nullString: NULL_STRING,
		writer:     nil,
		dataframe:  nil,
	}
}

func (w *MarkDownWriter) SetHeader(header bool) *MarkDownWriter {
	w.header = header
	return w
}

func (w *MarkDownWriter) SetPath(path string) *MarkDownWriter {
	w.path = path
	return w
}

func (w *MarkDownWriter) SetNullString(nullString string) *MarkDownWriter {
	w.nullString = nullString
	return w
}

func (w *MarkDownWriter) SetWriter(writer io.Writer) *MarkDownWriter {
	w.writer = writer
	return w
}

func (w *MarkDownWriter) SetDataFrame(dataframe DataFrame) *MarkDownWriter {
	w.dataframe = dataframe
	return w
}

func (w *MarkDownWriter) Write() DataFrame {
	err := writeMarkDown(w.dataframe, w.writer, w.header, w.nullString)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err}
	}

	return w.dataframe
}

func writeMarkDown(dataframe DataFrame, writer io.Writer, header bool, nullString string) error {
	return nil
}
