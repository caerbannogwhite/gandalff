package io

import (
	"os"

	"testing"
)

func Test_IOMd_ValidWrite(t *testing.T) {
	df := NewBaseDataFrame(ctx).
		AddSeriesFromFloat64s("a", []float64{1, 2, 3}, nil, false).
		AddSeriesFromStrings("b", []string{"a", "b", "c"}, nil, false).
		ToXlsx().
		SetPath("test.md").
		Write()

	if df.IsErrored() {
		t.Errorf(df.GetError().Error())
	}

	_, err := os.Stat("test.md")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = os.Remove("test.md")
	if err != nil {
		t.Errorf(err.Error())
	}
}
