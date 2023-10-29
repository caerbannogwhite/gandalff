package gandalff

import (
	"fmt"
	"io"
)

type HtmlWriter struct {
	naText    string
	writer    io.Writer
	dataframe DataFrame
}

func NewHtmlWriter() *HtmlWriter {
	return &HtmlWriter{
		naText:    NA_TEXT,
		writer:    nil,
		dataframe: nil,
	}
}

func (w *HtmlWriter) SetNaText(naText string) *HtmlWriter {
	w.naText = naText
	return w
}

func (w *HtmlWriter) SetWriter(writer io.Writer) *HtmlWriter {
	w.writer = writer
	return w
}

func (w *HtmlWriter) SetDataFrame(dataframe DataFrame) *HtmlWriter {
	w.dataframe = dataframe
	return w
}

func (w *HtmlWriter) Write() DataFrame {
	err := writeHtml(w.dataframe, w.writer, w.naText)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err}
	}

	return w.dataframe
}

func writeHtml(df DataFrame, writer io.Writer, naText string) error {
	series := make([]Series, df.NCols())
	for i := 0; i < df.NCols(); i++ {
		series[i] = df.SeriesAt(i)
	}

	if writer == nil {
		return fmt.Errorf("writeHtml: no writer specified")
	}

	_, err := writer.Write([]byte("<html><head><style>table, th, td {border: 1px solid black;}</style></head><body><table>"))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte("<tr>"))
	if err != nil {
		return err
	}

	for _, name := range df.Names() {
		_, err = writer.Write([]byte(fmt.Sprintf("<th>%s</th>", name)))
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte("</tr>"))
	if err != nil {
		return err
	}

	for i := 0; i < df.NRows(); i++ {
		_, err = writer.Write([]byte("<tr>"))
		if err != nil {
			return err
		}

		for _, s := range series {
			_, err = writer.Write([]byte(fmt.Sprintf("<td>%s</td>", s.GetAsString(i))))
			if err != nil {
				return err
			}
		}

		_, err = writer.Write([]byte("</tr>"))
		if err != nil {
			return err
		}
	}

	_, err = writer.Write([]byte("</table></body></html>"))
	if err != nil {
		return err
	}

	return nil
}
