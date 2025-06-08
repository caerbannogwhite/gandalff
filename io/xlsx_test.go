package io

import (
	"os"

	"testing"

	"github.com/caerbannogwhite/aargh/series"
)

func Test_IoXlsx_ValidWrite(t *testing.T) {
	iod := NewIoData(ctx)

	iod.AddSeries(series.NewSeriesFloat64([]float64{1, 2, 3}, nil, false, ctx), SeriesMeta{Name: "a"})
	iod.AddSeries(series.NewSeriesString([]string{"a", "b", "c"}, nil, false, ctx), SeriesMeta{Name: "b"})

	err := iod.ToXlsx().
		SetPath("test.xlsx").
		Write()

	if err != nil {
		t.Errorf(err.Error())
	}

	_, err = os.Stat("test.xlsx")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = os.Remove("test.xlsx")
	if err != nil {
		t.Errorf(err.Error())
	}
}
