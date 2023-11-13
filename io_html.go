package gandalff

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type HtmlWriter struct {
	path      string
	naText    string
	newLine   string
	indent    string
	writer    io.Writer
	dataframe DataFrame
}

func NewHtmlWriter() *HtmlWriter {
	return &HtmlWriter{
		path:      "",
		naText:    NA_TEXT,
		newLine:   "\n",
		indent:    "\t",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *HtmlWriter) SetPath(path string) *HtmlWriter {
	w.path = path
	return w
}

func (w *HtmlWriter) SetNaText(naText string) *HtmlWriter {
	w.naText = naText
	return w
}

func (w *HtmlWriter) SetNewLine(newLine string) *HtmlWriter {
	w.newLine = newLine
	return w
}

func (w *HtmlWriter) SetIndent(indent string) *HtmlWriter {
	w.indent = indent
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
	if w.dataframe == nil {
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeHtml: no dataframe specified"), ctx: w.dataframe.GetContext()}
		return w.dataframe
	}

	if w.dataframe.IsErrored() {
		return w.dataframe
	}

	if w.path != "" {
		file, err := os.OpenFile(w.path, os.O_CREATE, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
		}
		defer file.Close()
		w.writer = file
	}

	if w.writer == nil {
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeHtml: no writer specified"), ctx: w.dataframe.GetContext()}
		return w.dataframe
	}

	err := writeHtml(w.dataframe, w.writer, w.naText, w.newLine, w.indent)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeHtml(df DataFrame, writer io.Writer, naText, newLine, indent string) error {
	series := make([]Series, df.NCols())
	for i := 0; i < df.NCols(); i++ {
		series[i] = df.SeriesAt(i)
	}

	if writer == nil {
		return fmt.Errorf("writeHtml: no writer specified")
	}

	indent2 := indent + indent
	indent3 := indent2 + indent
	indent4 := indent3 + indent

	_, err := writer.Write([]byte(
		"<html>" + newLine +
			indent + "<head>" + newLine +
			indent + "</head>" + newLine +
			indent + "<body>" + newLine +
			indent2 + "<table>" + newLine))
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(indent3 + "<tr>" + newLine))
	if err != nil {
		return err
	}

	for _, name := range df.Names() {
		_, err = writer.Write([]byte(fmt.Sprintf("%s<th>%s</th>%s", indent4, name, newLine)))
		if err != nil {
			return err
		}
	}

	var rowBuffer strings.Builder
	rowBuffer.WriteString(indent3 + "</tr>" + newLine)

	for i := 0; i < df.NRows(); i++ {

		rowBuffer.WriteString(indent3 + "<tr>" + newLine)
		for _, s := range series {
			if s.IsNull(i) {
				rowBuffer.WriteString(fmt.Sprintf("%s<td>%s</td>%s", indent4, naText, newLine))
			} else {
				rowBuffer.WriteString(fmt.Sprintf("%s<td>%s</td>%s", indent4, s.GetAsString(i), newLine))
			}
		}
		rowBuffer.WriteString(indent3 + "</tr>" + newLine)

	}
	rowBuffer.WriteString(indent2 + "</table>" + newLine + indent + "</body>" + newLine + "</html>" + newLine)

	_, err = writer.Write([]byte(rowBuffer.String()))
	if err != nil {
		return err
	}

	return nil
}
