package gandalff

import (
	"fmt"
	"io"
	"os"
	"preludiometa"
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
	return nil, nil, nil
}

type JsonWriter struct {
	path      string
	writer    io.Writer
	dataframe DataFrame
}

func NewJsonWriter() *JsonWriter {
	return &JsonWriter{
		path:      "",
		writer:    nil,
		dataframe: nil,
	}
}

func (w *JsonWriter) SetPath(path string) *JsonWriter {
	w.path = path
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
		file, err := os.OpenFile(w.path, os.O_CREATE, 0666)
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

	err := writeJson(w.dataframe, w.writer)
	if err != nil {
		w.dataframe = BaseDataFrame{err: err, ctx: w.dataframe.GetContext()}
	}

	return w.dataframe
}

func writeJson(dataframe DataFrame, writer io.Writer) error {
	// TODO
	return nil
}
