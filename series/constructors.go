package series

import (
	"fmt"
	"time"

	"github.com/caerbannogwhite/aargh"
	"github.com/caerbannogwhite/aargh/utils"
)

// Build a Series from a generic interface.
// The interface can be a single value, slice of values,
// a nullable value, a slice of nullable values.
// If nullMask is nil then the series is not nullable.
func NewSeries(data interface{}, nullMask []bool, makeCopy bool, memOpt bool, ctx *aargh.Context) Series {
	if ctx == nil {
		return Errors{"NewSeries: context is nil"}
	}

	switch data := data.(type) {
	case nil:
		return NewSeriesNA(1, ctx)

	case bool:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesBool([]bool{data}, nil, false, ctx)
	case []bool:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesBool(data, nullMask, makeCopy, ctx)
	case aargh.NullableBool:
		return NewSeriesBool([]bool{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableBool:
		values := make([]bool, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesBool(values, nulls, false, ctx)

	case int:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesInt([]int{data}, nil, false, ctx)
	case []int:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt(data, nullMask, makeCopy, ctx)
	case aargh.NullableInt:
		return NewSeriesInt([]int{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableInt:
		values := make([]int, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesInt(values, nulls, false, ctx)

	case int64:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesInt64([]int64{data}, nil, false, ctx)
	case []int64:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesInt64(data, nullMask, makeCopy, ctx)
	case aargh.NullableInt64:
		return NewSeriesInt64([]int64{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableInt64:
		values := make([]int64, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesInt64(values, nulls, false, ctx)

	case float64:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesFloat64([]float64{data}, nil, false, ctx)
	case []float64:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesFloat64(data, nullMask, makeCopy, ctx)
	case aargh.NullableFloat64:
		return NewSeriesFloat64([]float64{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableFloat64:
		values := make([]float64, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesFloat64(values, nulls, false, ctx)

	case string:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesString([]string{data}, nil, false, ctx)
	case []string:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesString(data, nullMask, makeCopy, ctx)
	case aargh.NullableString:
		return NewSeriesString([]string{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableString:
		values := make([]string, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesString(values, nulls, false, ctx)

	case time.Time:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesTime([]time.Time{data}, nil, false, ctx)
	case []time.Time:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesTime(data, nullMask, makeCopy, ctx)
	case aargh.NullableTime:
		return NewSeriesTime([]time.Time{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableTime:
		values := make([]time.Time, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesTime(values, nulls, false, ctx)

	case time.Duration:
		if nullMask != nil && len(nullMask) != 1 {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length 1", len(nullMask))}
		}
		return NewSeriesDuration([]time.Duration{data}, nil, false, ctx)
	case []time.Duration:
		if nullMask != nil && len(nullMask) != len(data) {
			return Errors{fmt.Sprintf("NewSeries: null mask length %d does not match data length %d", len(nullMask), len(data))}
		}
		return NewSeriesDuration(data, nullMask, makeCopy, ctx)
	case aargh.NullableDuration:
		return NewSeriesDuration([]time.Duration{data.Value}, []bool{!data.Valid}, false, ctx)
	case []aargh.NullableDuration:
		values := make([]time.Duration, len(data))
		nulls := make([]bool, len(data))
		for i, v := range data {
			values[i] = v.Value
			nulls[i] = !v.Valid
		}
		return NewSeriesDuration(values, nulls, false, ctx)

	default:
		return Errors{fmt.Sprintf("NewSeries: unsupported type %T", data)}
	}
}

// Build an Error Series
func NewSeriesError(err string) Errors {
	return Errors{Msg_: err}
}

// Build an NA Series
func NewSeriesNA(size int, ctx *aargh.Context) NAs {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesNA: context is nil")
	}

	if size < 0 {
		size = 0
	}

	return NAs{size: size, Ctx_: ctx}
}

// Build a Bool Series, if nullMask is nil then the series is not nullable
func NewSeriesBool(data []bool, nullMask []bool, makeCopy bool, ctx *aargh.Context) Bools {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesBool: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]bool, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Bools{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a Int Series, if nullMask is nil then the series is not nullable
func NewSeriesInt(data []int, nullMask []bool, makeCopy bool, ctx *aargh.Context) Ints {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesInt: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Ints{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a Int64 Series, if nullMask is nil then the series is not nullable
func NewSeriesInt64(data []int64, nullMask []bool, makeCopy bool, ctx *aargh.Context) Int64s {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesInt64: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]int64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Int64s{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a Float64 Series, if nullMask is nil then the series is not nullable
func NewSeriesFloat64(data []float64, nullMask []bool, makeCopy bool, ctx *aargh.Context) Float64s {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesFloat64: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Float64s{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a String Series, if nullMask is nil then the series is not nullable
func NewSeriesString(data []string, nullMask []bool, makeCopy bool, ctx *aargh.Context) Strings {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesString: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	actualData := make([]*string, len(data))
	if nullMask != nil {
		for i, v := range data {
			if nullMask[i] {
				actualData[i] = ctx.StringPool.Put(aargh.NA_TEXT)
				continue
			}
			actualData[i] = ctx.StringPool.Put(v)
		}
	} else {
		for i, v := range data {
			actualData[i] = ctx.StringPool.Put(v)
		}
	}

	return Strings{
		IsNullable_: isNullable,
		Data_:       actualData,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a String Series from a slice of pointers to strings, if nullMask is nil then the series is not nullable
func NewSeriesStringFromPtrs(data []*string, nullMask []bool, makeCopy bool, ctx *aargh.Context) Strings {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesStringFromPtrs: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	return Strings{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}

// Build a Time Series, if nullMask is nil then the series is not nullable
func NewSeriesTime(data []time.Time, nullMask []bool, makeCopy bool, ctx *aargh.Context) Times {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesTime: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Time, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Times{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
		timeFormat:  ctx.GetDateTimeFormat(),
	}
}

// Build a Duration Series, if nullMask is nil then the series is not nullable
func NewSeriesDuration(data []time.Duration, nullMask []bool, makeCopy bool, ctx *aargh.Context) Durations {
	if ctx == nil {
		fmt.Println("WARNING: NewSeriesDuration: context is nil")
	}

	var isNullable bool
	var nullMask_ []uint8
	if nullMask != nil {
		isNullable = true
		if len(nullMask) < len(data) {
			nullMask = append(nullMask, make([]bool, len(data)-len(nullMask))...)
		} else if len(nullMask) > len(data) {
			nullMask = nullMask[:len(data)]
		}
		nullMask_ = utils.BinVecFromBools(nullMask)
	} else {
		isNullable = false
		nullMask_ = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]time.Duration, len(data))
		copy(actualData, data)
		data = actualData
	}

	return Durations{
		IsNullable_: isNullable,
		Data_:       data,
		NullMask_:   nullMask_,
		Ctx_:        ctx,
	}
}
