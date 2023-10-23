package markdown

import (
	"io"

	"github.com/caerbannogwhite/gandalff"
)

type MarkDownWriter struct {
	header     bool
	types      bool
	path       string
	nullString string
	writer     io.Writer
	dataframe  gandalff.DataFrame
}

func NewMarkDownWriter() *MarkDownWriter {
	return &MarkDownWriter{
		header:     gandalff.CSV_READER_DEFAULT_HEADER,
		path:       "",
		nullString: gandalff.NULL_STRING,
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

func (w *MarkDownWriter) SetDataFrame(dataframe gandalff.DataFrame) *MarkDownWriter {
	w.dataframe = dataframe
	return w
}

func (w *MarkDownWriter) Write() gandalff.DataFrame {
	err := writeMarkDown(w.dataframe, w.writer, w.header, w.nullString)
	if err != nil {
		w.dataframe = gandalff.BaseDataFrame{err: err}
	}

	return w.dataframe
}

func writeMarkDown(dataframe gandalff.DataFrame, writer io.Writer, header bool, nullString string) error {
}
