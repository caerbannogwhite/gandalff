package gandalff

import (
	"fmt"
	"io"
	"os"
	"preludiometa"
)

type MarkDownWriter struct {
	header    bool
	index     bool
	types     bool
	path      string
	naText    string
	writer    io.Writer
	dataframe DataFrame
}

func NewMarkDownWriter() *MarkDownWriter {
	return &MarkDownWriter{
		header:    true,
		index:     false,
		path:      "",
		naText:    NA_TEXT,
		writer:    nil,
		dataframe: nil,
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

func (w *MarkDownWriter) SetNaText(naText string) *MarkDownWriter {
	w.naText = naText
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
	if w.dataframe == nil {
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeMarkDown: no dataframe specified"), ctx: w.dataframe.GetContext()}
		return w.dataframe
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
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeMarkDown: no writer specified"), ctx: w.dataframe.GetContext()}
		return w.dataframe
	}

	err := writeMarkDown(w.dataframe, w.writer, w.header, w.index, w.naText)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeMarkDown(dataframe DataFrame, writer io.Writer, header, index bool, naText string) error {
	buff := ""
	if header {
		buff += "|"
		if index {
			buff += " |"
		}

		for _, col := range dataframe.Names() {
			buff += "**" + col + "**|"
		}
		buff += "\n"
	}

	buff += "|"
	if index {
		buff += "----:|"
	}

	for i := range dataframe.Names() {
		switch dataframe.Types()[i].GetDefaultJustification() {
		case preludiometa.JUSTIFY_LEFT:
			buff += ":----|"
		case preludiometa.JUSTIFY_RIGHT:
			buff += "----:|"
		case preludiometa.JUSTIFY_CENTER:
			buff += ":---:|"
		}
	}
	buff += "\n"

	for i := 0; i < dataframe.NRows(); i++ {
		buff += "|"
		if index {
			buff += fmt.Sprintf("%d|", i)
		}

		for j := 0; j < dataframe.NCols(); j++ {
			if dataframe.SeriesAt(j).IsNull(i) {
				buff += naText + "|"
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
