package gandalff

import (
	"os"

	"testing"
)

func Test_IOJson_ValidWrite(t *testing.T) {
	df := NewBaseDataFrame(ctx).
		AddSeriesFromFloat64s("a", []float64{1, 2, 3}, nil, false).
		AddSeriesFromStrings("b", []string{"a", "b", "c"}, nil, false).
		ToXlsx().
		SetPath("test.json").
		Write()

	if df.IsErrored() {
		t.Errorf(df.GetError().Error())
	}

	_, err := os.Stat("test.json")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = os.Remove("test.json")
	if err != nil {
		t.Errorf(err.Error())
	}
}
