package io

import (
	"os"
	"testing"
)

func Test_IOAll(t *testing.T) {

	df := NewBaseDataFrame(ctx).
		FromXpt().
		SetPath("testdata/xpt_test_mixed.xpt").
		SetVersion(XPT_VERSION_9).
		Read().

		// to JSON
		ToJson().
		SetPath("testdata/test.json").
		Write().

		// From JSON
		FromJson().
		SetPath("testdata/test.json").
		Read().

		// to CSV
		ToCsv().
		SetPath("testdata/test.csv").
		SetNaText("").
		SetDelimiter(',').
		Write().

		// From CSV
		FromCsv().
		SetPath("testdata/test.csv").
		SetDelimiter(',').
		SetHeader(true).
		Read().

		// to Excel
		ToXlsx().
		SetPath("testdata/test.xlsx").
		SetSheet("test").
		SetNaText("").
		Write().

		// From Excel
		FromXlsx().
		SetPath("testdata/test.xlsx").
		SetSheet("test").
		Read()

	os.Remove("testdata/test.json")
	os.Remove("testdata/test.csv")
	os.Remove("testdata/test.xlsx")

	if df.NCols() != 4 {
		t.Errorf("expected 4 columns, got %d", df.NCols())
	}

	if df.NRows() != 7 {
		t.Errorf("expected 7 rows, got %d", df.NRows())
	}

	charvar1 := NewSeries([]string{
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
	}, nil, true, false, ctx)

	charvar2 := NewSeries([]string{
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
	}, nil, true, false, ctx)

	numvar1 := NewSeries([]int64{1, 2, 3, 4, 5, 6, 7}, nil, true, false, ctx)
	numvar2 := NewSeries([]float64{
		1.2345e2,
		6.5432e1,
		1.2345e0,
		6.5432e-1,
		1.2345e-2,
		6.543e-3, // 6.5432e-3
		1.23e-4,  // 1.2345e-4
	}, nil, true, false, ctx)

	if series := df.C("CHARVAR1"); !series.Eq(charvar1).(Bools).All() {
		t.Errorf("expected %s, got %s", charvar1.DataAsString(), series.DataAsString())
	}

	if series := df.C("CHARVAR2"); !series.Eq(charvar2).(Bools).All() {
		t.Errorf("expected %s, got %s", charvar2.DataAsString(), series.DataAsString())
	}

	if series := df.C("NUMVAR1"); !series.Eq(numvar1).(Bools).All() {
		t.Errorf("expected %s, got %s", numvar1.DataAsString(), series.DataAsString())
	}

	if series := df.C("FOO"); !series.Eq(numvar2).(Bools).All() {
		t.Errorf("expected %s, got %s", numvar2.DataAsString(), series.DataAsString())
	}
}
