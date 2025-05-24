package io

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/caerbannogwhite/gandalff/series"
)

func Test_IOAll(t *testing.T) {
	var err error
	var iod *IoData

	iod, err = FromXpt(ctx).
		SetPath(filepath.Join(testDataFolder, "xpt_test_mixed.xpt")).
		SetVersion(XPT_VERSION_9).
		Read()

	if err != nil {
		t.Errorf(err.Error())
	}

	// to JSON
	err = iod.ToJson().
		SetPath(filepath.Join(testDataFolder, "test.json")).
		Write()

	// From JSON
	iod, err = FromJson(ctx).
		SetPath(filepath.Join(testDataFolder, "test.json")).
		Read()

	if err != nil {
		t.Errorf(err.Error())
	}

	// to CSV
	err = iod.ToCsv().
		SetPath(filepath.Join(testDataFolder, "test.csv")).
		SetNaText("").
		SetDelimiter(',').
		Write()

	// From CSV
	iod, err = FromCsv(ctx).
		SetPath(filepath.Join(testDataFolder, "test.csv")).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	if err != nil {
		t.Errorf(err.Error())
	}

	// to Excel
	err = iod.ToXlsx().
		SetPath(filepath.Join(testDataFolder, "test.xlsx")).
		SetSheet("test").
		SetNaText("").
		Write()

	// From Excel
	iod, err = FromXlsx(ctx).
		SetPath(filepath.Join(testDataFolder, "test.xlsx")).
		SetSheet("test").
		Read()

	if err != nil {
		t.Errorf(err.Error())
	}

	os.Remove(filepath.Join(testDataFolder, "test.json"))
	os.Remove(filepath.Join(testDataFolder, "test.csv"))
	os.Remove(filepath.Join(testDataFolder, "test.xlsx"))

	if iod.NCols() != 4 {
		t.Errorf("expected 4 columns, got %d", iod.NCols())
	}

	if iod.NRows() != 7 {
		t.Errorf("expected 7 rows, got %d", iod.NRows())
	}

	charvar1 := series.NewSeriesString([]string{
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
		"wbiwbui749",
		"abcdefghij",
	}, nil, false, ctx)

	charvar2 := series.NewSeriesString([]string{
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
		"nionione983203jnfui2",
		"abcdefghijklmnopqrst",
	}, nil, false, ctx)

	numvar1 := series.NewSeriesInt64([]int64{1, 2, 3, 4, 5, 6, 7}, nil, false, ctx)
	numvar2 := series.NewSeriesFloat64([]float64{
		1.2345e2,
		6.5432e1,
		1.2345e0,
		6.5432e-1,
		1.2345e-2,
		6.543e-3, // 6.5432e-3
		1.23e-4,  // 1.2345e-4
	}, nil, false, ctx)

	if _series := iod.ByName("CHARVAR1"); !_series.Eq(charvar1).(series.Bools).All() {
		t.Errorf("expected %s, got %s", charvar1.DataAsString(), _series.DataAsString())
	}

	if _series := iod.ByName("CHARVAR2"); !_series.Eq(charvar2).(series.Bools).All() {
		t.Errorf("expected %s, got %s", charvar2.DataAsString(), _series.DataAsString())
	}

	if _series := iod.ByName("NUMVAR1"); !_series.Eq(numvar1).(series.Bools).All() {
		t.Errorf("expected %s, got %s", numvar1.DataAsString(), _series.DataAsString())
	}

	if _series := iod.ByName("FOO"); !_series.Eq(numvar2).(series.Bools).All() {
		t.Errorf("expected %s, got %s", numvar2.DataAsString(), _series.DataAsString())
	}
}
