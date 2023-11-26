package gandalff

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type HtmlWriter struct {
	path       string
	naText     string
	newLine    string
	indent     string
	writer     io.Writer
	dataframe  DataFrame
	datatables bool
}

func NewHtmlWriter() *HtmlWriter {
	return &HtmlWriter{
		path:       "",
		naText:     NA_TEXT,
		newLine:    "\n",
		indent:     "\t",
		writer:     nil,
		dataframe:  nil,
		datatables: false,
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

func (w *HtmlWriter) SetDatatables(datatables bool) *HtmlWriter {
	w.datatables = datatables
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

	err := writeHtml(w.dataframe, w.writer, w.naText, w.newLine, w.indent, w.datatables)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeHtml(df DataFrame, writer io.Writer, naText, newLine, indent string, datatables bool) error {
	head := ""
	if datatables {
		head = `
        <link href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/5.3.0/css/bootstrap.min.css" rel="stylesheet">
        <link href="https://cdn.datatables.net/1.13.8/css/dataTables.bootstrap5.min.css" rel="stylesheet">
        <link href="https://cdn.datatables.net/buttons/2.4.2/css/buttons.bootstrap5.min.css" rel="stylesheet">
        <link href="https://cdn.datatables.net/fixedcolumns/4.3.0/css/fixedColumns.bootstrap5.min.css" rel="stylesheet">
        <link href="https://cdn.datatables.net/fixedheader/3.4.0/css/fixedHeader.bootstrap5.min.css" rel="stylesheet">
        
        <script src="https://code.jquery.com/jquery-3.7.0.min.js"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/5.3.0/js/bootstrap.bundle.min.js"></script>
        <script src="https://cdn.datatables.net/1.13.8/js/jquery.dataTables.min.js"></script>
        <script src="https://cdn.datatables.net/1.13.8/js/dataTables.bootstrap5.min.js"></script>
        <script src="https://cdn.datatables.net/buttons/2.4.2/js/dataTables.buttons.min.js"></script>
        <script src="https://cdn.datatables.net/buttons/2.4.2/js/buttons.bootstrap5.min.js"></script>
        <script src="https://cdn.datatables.net/fixedcolumns/4.3.0/js/dataTables.fixedColumns.min.js"></script>
        <script src="https://cdn.datatables.net/fixedheader/3.4.0/js/dataTables.fixedHeader.min.js"></script>
		<script>
			$(document).ready(function () {
				$("#example").DataTable({
					scrollX: true,
					scrollY: true,
					scrollCollapse: false,
					paging: true,
				});
			});
		</script>
		`
	}

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
	indent5 := indent4 + indent

	tableHTML := "<table>"
	if datatables {
		tableHTML = "<table id=\"example\" class=\"display\" style=\"width:100%\">"
	}

	_, err := writer.Write([]byte(
		"<html>" + newLine +
			indent + "<head>" + newLine + head + newLine +
			indent + "</head>" + newLine +
			indent + "<body>" + newLine +
			indent2 + tableHTML + newLine +
			indent3 + "<thead>" + newLine))
	if err != nil {
		return err
	}

	// Table header
	_, err = writer.Write([]byte(indent3 + "<tr>" + newLine))
	if err != nil {
		return err
	}

	for _, name := range df.Names() {
		_, err = writer.Write([]byte(fmt.Sprintf("%s<th>%s</th>%s", indent5, name, newLine)))
		if err != nil {
			return err
		}
	}

	var rowBuffer strings.Builder
	rowBuffer.WriteString(indent3 + "</tr>" + newLine)
	rowBuffer.WriteString(indent3 + "</thead>" + newLine + indent3 + "<tbody>" + newLine)

	// Table body
	for i := 0; i < df.NRows(); i++ {

		rowBuffer.WriteString(indent3 + "<tr>" + newLine)
		for _, s := range series {
			if s.IsNull(i) {
				rowBuffer.WriteString(fmt.Sprintf("%s<td>%s</td>%s", indent5, naText, newLine))
			} else {
				rowBuffer.WriteString(fmt.Sprintf("%s<td>%s</td>%s", indent5, s.GetAsString(i), newLine))
			}
		}
		rowBuffer.WriteString(indent3 + "</tr>" + newLine)

	}
	rowBuffer.WriteString("" +
		indent3 + "</tbody>" + newLine +
		indent2 + "</table>" + newLine +
		indent + "</body>" + newLine + "</html>" + newLine)

	_, err = writer.Write([]byte(rowBuffer.String()))
	if err != nil {
		return err
	}

	return nil
}
