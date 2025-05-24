package io

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/caerbannogwhite/preludiometa"
)

type JsonReader struct {
	path   string
	reader io.Reader
	schema *preludiometa.Schema
	ctx    *Context
}

func NewJsonReader(ctx *Context) *JsonReader {
	return &JsonReader{
		path:   "",
		reader: nil,
		schema: nil,
		ctx:    ctx,
	}
}

func (r *JsonReader) SetPath(path string) *JsonReader {
	r.path = path
	return r
}

func (r *JsonReader) SetReader(reader io.Reader) *JsonReader {
	r.reader = reader
	return r
}

func (r *JsonReader) SetSchema(schema *preludiometa.Schema) *JsonReader {
	r.schema = schema
	return r
}

func (r *JsonReader) Read() DataFrame {
	if r.path != "" {
		file, err := os.OpenFile(r.path, os.O_RDONLY, 0666)
		if err != nil {
			return BaseDataFrame{err: err, ctx: r.ctx}
		}
		defer file.Close()
		r.reader = file
	}

	if r.reader == nil {
		return BaseDataFrame{err: fmt.Errorf("JsonReader: no reader specified"), ctx: r.ctx}
	}

	if r.ctx == nil {
		return BaseDataFrame{err: fmt.Errorf("JsonReader: no context specified"), ctx: r.ctx}
	}

	names, series, err := readJson(r.reader, r.schema, r.ctx)
	if err != nil {
		return BaseDataFrame{err: err, ctx: r.ctx}
	}

	df := NewBaseDataFrame(r.ctx)
	for i, name := range names {
		df = df.AddSeries(name, series[i])
	}

	return df
}

func readJson(reader io.Reader, schema *preludiometa.Schema, ctx *Context) ([]string, []Series, error) {

	tokens := json.NewDecoder(reader)

	token, err := tokens.Token()
	if err != nil || token != json.Delim('{') {
		return nil, nil, fmt.Errorf("readJson: invalid json, expected '{' at the beginning")
	}

	var names []string
	var series []Series

	// read by column
	for {

		// read column name or end of json
		token, err := tokens.Token()
		if err != nil {
			return nil, nil, fmt.Errorf("readJson: invalid json, expected column name or '}'")
		}

		switch t := token.(type) {
		case json.Delim:
			if t == json.Delim('}') {
				return names, series, nil
			} else {
				return nil, nil, fmt.Errorf("readJson: invalid json, expected column name or '}'")
			}

		case string:
			names = append(names, t)

		default:
			return nil, nil, fmt.Errorf("readJson: invalid json, expected column name or '}'")
		}

		// read array
		token, err = tokens.Token()
		if err != nil {
			return nil, nil, fmt.Errorf("readJson: invalid json, column '%s' expected array begin, got '%v'", names[len(names)-1], token)
		}

		if token != json.Delim('{') {
			return nil, nil, fmt.Errorf("readJson: invalid json, column '%s' expected array begin, got '%v'", names[len(names)-1], token)
		}

		// read column values
		var type_ preludiometa.BaseType
		var boolValues []bool
		var floatValues []float64
		var stringValues []string

		for tokens.More() {

			// read index
			token, err := tokens.Token()
			if err != nil {
				return nil, nil, fmt.Errorf("readJson: invalid json")
			}

			token, err = tokens.Token()
			if err != nil {
				return nil, nil, fmt.Errorf("readJson: invalid json")
			}

			switch t := token.(type) {
			case bool:
				type_ = preludiometa.BoolType
				boolValues = append(boolValues, t)

			case float64:
				type_ = preludiometa.Float64Type
				floatValues = append(floatValues, t)

			case string:
				type_ = preludiometa.StringType
				stringValues = append(stringValues, t)

			default:
				return nil, nil, fmt.Errorf("readJson: invalid json")
			}
		}

		// read end of array
		token, err = tokens.Token()
		if err != nil {
			return nil, nil, fmt.Errorf("readJson: invalid json, column '%s' expected array end, got '%v'", names[len(names)-1], token)
		}

		if token != json.Delim('}') {
			return nil, nil, fmt.Errorf("readJson: invalid json, column '%s' expected array end, got '%v'", names[len(names)-1], token)
		}

		// create series
		switch type_ {
		case preludiometa.BoolType:
			series = append(series, NewSeriesBool(boolValues, nil, false, ctx))
		case preludiometa.Float64Type:
			series = append(series, NewSeriesFloat64(floatValues, nil, false, ctx))
		case preludiometa.StringType:
			series = append(series, NewSeriesString(stringValues, nil, false, ctx))
		}
	}
}

type JsonWriter struct {
	path      string
	newLine   string
	indent    string
	writer    io.Writer
	dataframe DataFrame
}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{
		path:      "",
		newLine:   "\n",
		indent:    "\t",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *JsonWriter) SetPath(path string) *JsonWriter {
	w.path = path
	return w
}

func (w *JsonWriter) SetNewLine(newLine string) *JsonWriter {
	w.newLine = newLine
	return w
}

func (w *JsonWriter) SetIndent(indent string) *JsonWriter {
	w.indent = indent
	return w
}

func (w *JsonWriter) SetWriter(writer io.Writer) *JsonWriter {
	w.writer = writer
	return w
}

func (w *JsonWriter) SetDataFrame(dataframe DataFrame) *JsonWriter {
	w.dataframe = dataframe
	return w
}

func (w *JsonWriter) Write() DataFrame {
	if w.dataframe == nil {
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeJson: no dataframe specified"), ctx: w.dataframe.GetContext()}
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
		w.dataframe = BaseDataFrame{err: fmt.Errorf("writeJson: no writer specified"), ctx: w.dataframe.GetContext()}
		return w.dataframe
	}

	err := writeJson(w.dataframe, w.writer, w.newLine, w.indent)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeJson(dataframe DataFrame, writer io.Writer, newLine, indent string) error {
	indent2 := indent + indent

	writer.Write([]byte("{\n"))
	for i, name := range dataframe.Names() {
		writer.Write([]byte(fmt.Sprintf("%s\"%s\": {%s", indent2, name, newLine)))

		series := dataframe.At(i)
		switch ser := series.(type) {
		case SeriesBool:
			for j, b := range ser.Bools() {
				if series.IsNull(j) {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": null", indent2, j)))
				} else {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": %v", indent2, j, b)))
				}

				if j < series.Len()-1 {
					writer.Write([]byte(","))
				}
				writer.Write([]byte(newLine))
			}

		case SeriesInt:
			for j, n := range ser.Ints() {
				if series.IsNull(j) {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": null", indent2, j)))
				} else {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": %d", indent2, j, n)))
				}

				if j < series.Len()-1 {
					writer.Write([]byte(","))
				}
				writer.Write([]byte(newLine))
			}

		case SeriesInt64:
			for j, n := range ser.Int64s() {
				if series.IsNull(j) {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": null", indent2, j)))
				} else {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": %d", indent2, j, n)))
				}

				if j < series.Len()-1 {
					writer.Write([]byte(","))
				}
				writer.Write([]byte(newLine))
			}

		case SeriesFloat64:
			for j, f := range ser.Float64s() {
				if series.IsNull(j) {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": null", indent2, j)))
				} else {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": %f", indent2, j, f)))
				}

				if j < series.Len()-1 {
					writer.Write([]byte(","))
				}
				writer.Write([]byte(newLine))
			}

		case SeriesString:
			for j, s := range ser.Strings() {
				if series.IsNull(j) {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": null", indent2, j)))
				} else {
					writer.Write([]byte(fmt.Sprintf("%s\"%d\": \"%s\"", indent2, j, s)))
				}

				if j < series.Len()-1 {
					writer.Write([]byte(","))
				}
				writer.Write([]byte(newLine))
			}
		}

		writer.Write([]byte(indent + "}"))
		if i < dataframe.NCols()-1 {
			writer.Write([]byte(","))
		}
		writer.Write([]byte(newLine))
	}

	writer.Write([]byte("}" + newLine))
	return nil
}
