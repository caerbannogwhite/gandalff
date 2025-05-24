package io

import (
	"encoding/binary"
	"os"

	"math"

	"testing"
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

func Test_IOXpt_Overflow(t *testing.T) {
	ibm := SasFloat{}
	err := ibm.FromIeee(math.Pow(16, 63), binary.BigEndian)
	if err == nil || err.Error() != "cannot store magnitude more than ~ 16 ** 63 as IBM-format" {
		t.FailNow()
	}
}

func Test_IOXpt_Underflow(t *testing.T) {
	ibm := SasFloat{}
	err := ibm.FromIeee(math.Pow(16, -66), binary.BigEndian)
	if err == nil || err.Error() != "cannot store magnitude less than ~ 16 ** -65 as IBM-format" {
		t.FailNow()
	}
}

func Test_IOXpt_Nan(t *testing.T) {
	res, err := Roundtrip(math.NaN())
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	if !math.IsNaN(res) {
		t.Errorf("%f != %f (actual)\n", res, math.NaN())
		t.FailNow()
	}
}

func Test_IOXpt_SpecialMissingValues(t *testing.T) {

	// From A to Z
	for i := byte('A'); i <= byte('Z'); i++ {
		v := math.Float64frombits(binary.BigEndian.Uint64([]byte{i, 0, 0, 0, 0, 0, 0, 0}))
		res, err := Roundtrip(v)
		if err != nil {
			t.Errorf(err.Error())
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
		t.Errorf(err.Error())
		t.FailNow()
	}

	if v != res || !IsIbmSpecialMissingValue(v) || !IsIbmSpecialMissingValue(res) {
		t.FailNow()
	}
}

func Test_IOXpt_Zero(t *testing.T) {
	res, err := Roundtrip(0)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}

	if math.Abs(res-0) > DELTA {
		t.Errorf("%f != %f (actual)\n", res, 0.0)
		t.FailNow()
	}
}

func Test_IOXpt_SmallMagnitudeIntegers(t *testing.T) {
	for i := -1000; i < 1000; i++ {
		res, err := Roundtrip(float64(i))
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-float64(i)) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, float64(i))
			t.FailNow()
		}
	}
}

func Test_IOXpt_LargeMagnitudeFloats(t *testing.T) {
	n := int(1e9)
	for i := n; i < n+100; i++ {
		res, err := Roundtrip(float64(i))
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-float64(i)) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, float64(i))
			t.FailNow()
		}
	}
}

func Test_IOXpt_LargeMagnitudeFloatsWithFraction(t *testing.T) {
	offset := 1e9
	for i := 0; i < 100; i++ {
		x := (float64(i) / 1e9) + offset
		res, err := Roundtrip(x)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-x) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, x)
			t.FailNow()
		}
	}
}

func Test_IOXpt_SmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e3
		res, err := Roundtrip(v)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_IOXpt_VerySmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e6
		res, err := Roundtrip(v)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_IOXpt_VeryVerySmallMagnitudeFloats(t *testing.T) {
	for i := -20; i < 20; i++ {
		v := float64(i) / 1.0e9
		res, err := Roundtrip(v)
		if err != nil {
			t.Errorf(err.Error())
			t.FailNow()
		}

		if math.Abs(res-v) > DELTA {
			t.Errorf("%f != %f (actual)\n", res, v)
			t.FailNow()
		}
	}
}

func Test_IOXpt_ValidWrite(t *testing.T) {
	df := NewBaseDataFrame(ctx).
		AddSeriesFromFloat64s("a", []float64{1, 2, 3}, nil, false).
		AddSeriesFromStrings("b", []string{"a", "b", "c"}, nil, false).
		ToXpt().
		SetPath("test.xpt").
		Write()

	if df.IsErrored() {
		t.Errorf(df.GetError().Error())
	}

	_, err := os.Stat("test.xpt")
	if err != nil {
		t.Errorf(err.Error())
	}

	err = os.Remove("test.xpt")
	if err != nil {
		t.Errorf(err.Error())
	}
}
