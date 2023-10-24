package gandalff

import (
	"io"
	"preludiometa"
)

type MarkDownWriter struct {
	header     bool
	index      bool
	types      bool
	path       string
	nullString string
	writer     io.Writer
	dataframe  DataFrame
}

func NewMarkDownWriter() *MarkDownWriter {
	return &MarkDownWriter{
		header:     true,
		index:      false,
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

func (w *MarkDownWriter) SetIndex(index bool) *MarkDownWriter {
	w.index = index
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

	buff := ""

	if header {
		buff += "|"
		for _, col := range dataframe.Names() {
			buff += "**" + col + "**|"
		}
		buff += "\n"
	}

	buff += "|"
	for i, _ := range dataframe.Names() {
		switch preludiometa.GetDefaultJustification(dataframe.Types()[i]) {
		case preludiometa.JUSTIFY_LEFT:
			buff += ":---"
		case preludiometa.JUSTIFY_RIGHT:
			buff += "---:"
		case preludiometa.JUSTIFY_CENTER:
			buff += ":---:"
		}
		buff += "|"
	}
	buff += "\n"

	for i := 0; i < dataframe.NRows(); i++ {
		buff += "|"
		for j := 0; j < dataframe.NCols(); j++ {
			if dataframe.SeriesAt(j).IsNull(i) {
				buff += nullString + "|"
			} else {
				buff += dataframe.SeriesAt(j).GetAsString(i) + "|"
			}
		}
		buff += "\n"
	}

	_, err := writer.Write([]byte(buff))
	if err != nil {
		return err
	}

	return nil
}
