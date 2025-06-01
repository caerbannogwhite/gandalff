package io

import (
	"fmt"
	"time"

	"github.com/caerbannogwhite/gandalff"
	"github.com/caerbannogwhite/gandalff/meta"
	"github.com/caerbannogwhite/gandalff/series"
)

type IoData struct {
	ctx *gandalff.Context

	FileMeta   FileMeta
	SeriesMeta []SeriesMeta
	Series     []series.Series
	Error      error
}

type FileMeta struct {
	FileName       string
	FilePath       string
	Label          string
	Created        time.Time
	LastModified   time.Time
	SasLibVersion  string
	SasDataVersion string
	SasOs          string
}

func (fileMeta *FileMeta) ToString() string {
	return fmt.Sprintf("FileMeta{FileName: %s, FilePath: %s, Label: %s, Created: %s, LastModified: %s, SasLibVersion: %s, SasDataVersion: %s, SasOs: %s}", fileMeta.FileName, fileMeta.FilePath, fileMeta.Label, fileMeta.Created, fileMeta.LastModified, fileMeta.SasLibVersion, fileMeta.SasDataVersion, fileMeta.SasOs)
}

func (fileMeta *FileMeta) PrettyPrint() {
	fmt.Println("File Meta")
	fmt.Println("--------------------------------")
	fmt.Println("File Name: ", fileMeta.FileName)
	fmt.Println("File Path: ", fileMeta.FilePath)
	fmt.Println("Label: ", fileMeta.Label)
	fmt.Println("Created: ", fileMeta.Created)
	fmt.Println("Last Modified: ", fileMeta.LastModified)
	fmt.Println("Sas Lib Version: ", fileMeta.SasLibVersion)
	fmt.Println("Sas Data Version: ", fileMeta.SasDataVersion)
	fmt.Println("Sas Os: ", fileMeta.SasOs)
}

type SeriesMeta struct {
	Format      string
	Label       string
	Length      int
	KeySequence int
	Name        string
	Type        meta.BaseType
}

func (iod *IoData) AddSeries(series series.Series, meta SeriesMeta) {
	iod.Series = append(iod.Series, series)
	iod.SeriesMeta = append(iod.SeriesMeta, meta)
}

func (iod *IoData) At(i int) series.Series {
	return iod.Series[i]
}

func (iod *IoData) ByName(name string) series.Series {
	for i, meta := range iod.SeriesMeta {
		if meta.Name == name {
			return iod.Series[i]
		}
	}
	return nil
}

func (iod *IoData) SeriesMetaAt(i int) SeriesMeta {
	return iod.SeriesMeta[i]
}

func (iod *IoData) Types() []meta.BaseType {
	types := make([]meta.BaseType, len(iod.Series))
	for i, series := range iod.Series {
		types[i] = series.Type()
	}
	return types
}

func (iod *IoData) GetContext() *gandalff.Context {
	return iod.ctx
}

func (iod *IoData) NRows() int {
	if len(iod.Series) >= 1 {
		return iod.Series[0].Len()
	}
	return 0
}

func (iod *IoData) NCols() int {
	return len(iod.Series)
}

func (iod *IoData) ToCsv() *CsvWriter {
	return NewCsvWriter().SetIoData(iod)
}

func (iod *IoData) ToHtml() *HtmlWriter {
	return NewHtmlWriter().SetIoData(iod)
}

func (iod *IoData) ToJson() *JsonWriter {
	return NewJsonWriter().SetIoData(iod)
}

func (iod *IoData) ToMarkdown() *MarkDownWriter {
	return NewMarkDownWriter().SetIoData(iod)
}

func (iod *IoData) ToXlsx() *XlsxWriter {
	return NewXlsxWriter().SetIoData(iod)
}

func (iod *IoData) ToXpt() *XptWriter {
	return NewXptWriter().SetIoData(iod)
}

func NewIoData(ctx *gandalff.Context) *IoData {
	return &IoData{
		ctx:        ctx,
		Series:     make([]series.Series, 0),
		SeriesMeta: make([]SeriesMeta, 0),
		FileMeta:   FileMeta{},
	}
}

func FromCsv(ctx *gandalff.Context) *CsvReader {
	return NewCsvReader(ctx)
}

func FromJson(ctx *gandalff.Context) *JsonReader {
	return NewJsonReader(ctx)
}

func FromXlsx(ctx *gandalff.Context) *XlsxReader {
	return NewXlsxReader(ctx)
}

func FromXpt(ctx *gandalff.Context) *XptReader {
	return NewXptReader(ctx)
}
