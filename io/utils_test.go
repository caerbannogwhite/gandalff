package io

import (
	"encoding/binary"
	"math"
	"testing"
	"time"
)

const DELTA = 10e-16

func Roundtrip(n float64) (float64, error) {
	ibm := SasFloat{}
	err := ibm.FromIeee(n, binary.BigEndian)
	if err != nil {
		return 0, err
	}

	ieee, err := ibm.ToIeee(binary.BigEndian)
	if err != nil {
		return 0, err
	}

	u := math.Float64bits(ieee)
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, u)

	return ieee, nil
}

func Test_IoXpt_Overflow(t *testing.T) {
	ibm := SasFloat{}
	err := ibm.FromIeee(math.Pow(16, 63), binary.BigEndian)
	if err == nil || err.Error() != "cannot store magnitude more than ~ 16 ** 63 as IBM-format" {
		t.FailNow()
	}
}

func Test_IoXpt_Underflow(t *testing.T) {
	ibm := SasFloat{}
	err := ibm.FromIeee(math.Pow(16, -66), binary.BigEndian)
	if err == nil || err.Error() != "cannot store magnitude less than ~ 16 ** -65 as IBM-format" {
		t.FailNow()
	}
}

func Test_IoXpt_Nan(t *testing.T) {
	res, err := Roundtrip(math.NaN())
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if !math.IsNaN(res) {
		t.Errorf("%f != %f (actual)\n", res, math.NaN())
		t.FailNow()
	}
}

func Test_IoXpt_SpecialMissingValues(t *testing.T) {

	// From A to Z
	for i := byte('A'); i <= byte('Z'); i++ {
		v := math.Float64frombits(binary.BigEndian.Uint64([]byte{i, 0, 0, 0, 0, 0, 0, 0}))
		res, err := Roundtrip(v)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if v != res || !IsIbmSpecialMissingValue(v) || !IsIbmSpecialMissingValue(res) {
			t.FailNow()
		}
	}

	// Underscore
	v := math.Float64frombits(binary.BigEndian.Uint64([]byte{byte('_'), 0, 0, 0, 0, 0, 0, 0}))
	res, err := Roundtrip(v)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if v != res || !IsIbmSpecialMissingValue(v) || !IsIbmSpecialMissingValue(res) {
		t.FailNow()
	}
}

func Test_IoXpt_Zero(t *testing.T) {
	res, err := Roundtrip(0)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}

	if math.Abs(res-0) > DELTA {
		t.Errorf("%f != %f (actual)\n", res, 0.0)
		t.FailNow()
	}
}

func Test_IoXpt_SmallMagnitudeIntegers(t *testing.T) {
	for i := -1000; i < 1000; i++ {
		res, err := Roundtrip(float64(i))
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-float64(i)) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, float64(i))
			t.FailNow()
		}
	}
}

func Test_IoXpt_LargeMagnitudeFloats(t *testing.T) {
	n := int(1e9)
	for i := n; i < n+100; i++ {
		res, err := Roundtrip(float64(i))
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-float64(i)) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, float64(i))
			t.FailNow()
		}
	}
}

func Test_IoXpt_LargeMagnitudeFloatsWithFraction(t *testing.T) {
	offset := 1e9
	for i := 0; i < 100; i++ {
		x := (float64(i) / 1e9) + offset
		res, err := Roundtrip(x)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-x) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, x)
			t.FailNow()
		}
	}
}

func Test_IoXpt_SmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e3
		res, err := Roundtrip(v)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_IoXpt_VerySmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e6
		res, err := Roundtrip(v)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_IoXpt_VeryVerySmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e9
		res, err := Roundtrip(v)
		if err != nil {
			t.Error(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_SasNumericToDate(t *testing.T) {

	// 2024-11-20 00:00:00 +0000 UTC 23700
	date := sasNumericToDate(23700)
	if date.Year() != 2024 || date.Month() != 11 || date.Day() != 20 {
		t.Errorf("%v != %v (actual)\n", date, time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC))
		t.FailNow()
	}

	// 2024-03-18 00:00:00 +0000 UTC 23453
	date = sasNumericToDate(23453)
	if date.Year() != 2024 || date.Month() != 3 || date.Day() != 18 {
		t.Errorf("%v != %v (actual)\n", date, time.Date(2024, 3, 18, 0, 0, 0, 0, time.UTC))
		t.FailNow()
	}

	// 2025-02-21 00:00:00 +0000 UTC 23793
	date = sasNumericToDate(23793)
	if date.Year() != 2025 || date.Month() != 2 || date.Day() != 21 {
		t.Errorf("%v != %v (actual)\n", date, time.Date(2025, 2, 21, 0, 0, 0, 0, time.UTC))
		t.FailNow()
	}
}

func Test_SasDateToNumeric(t *testing.T) {
	// 2024-11-20 00:00:00 +0000 UTC 23700
	date := time.Date(2024, 11, 20, 0, 0, 0, 0, time.UTC)
	numeric := sasDateToNumeric(date)
	if numeric != 23700 {
		t.Errorf("%d != %d (actual)\n", numeric, 23700)
		t.FailNow()
	}

	// 2024-03-18 00:00:00 +0000 UTC 23453
	date = time.Date(2024, 3, 18, 0, 0, 0, 0, time.UTC)
	numeric = sasDateToNumeric(date)
	if numeric != 23453 {
		t.Errorf("%d != %d (actual)\n", numeric, 23453)
		t.FailNow()
	}

	// 2025-02-21 00:00:00 +0000 UTC 23793
	date = time.Date(2025, 2, 21, 0, 0, 0, 0, time.UTC)
	numeric = sasDateToNumeric(date)
	if numeric != 23793 {
		t.Errorf("%d != %d (actual)\n", numeric, 23793)
		t.FailNow()
	}
}
