package series

import (
	"fmt"
	"time"

	"github.com/caerbannogwhite/gandalff"
	"github.com/caerbannogwhite/gandalff/meta"
)

// NAs represents a series with no data.
type NAs struct {
	size      int
	partition *SeriesNAPartition
	ctx       *gandalff.Context
}

func (s NAs) printInfo() {}

// Return the context of the series.
func (s NAs) GetContext() *gandalff.Context {
	return s.ctx
}

// Returns the length of the series.
func (s NAs) Len() int {
	return s.size
}

// Returns if the series is grouped.
func (s NAs) IsGrouped() bool {
	return s.partition != nil
}

// Returns if the series admits null values.
func (s NAs) IsNullable() bool {
	return true
}

func (s NAs) IsSorted() gandalff.SeriesSortOrder {
	return gandalff.SORTED_ASC
}

// Returns if the series is error.
func (s NAs) IsError() bool {
	return false
}

// Returns the error message of the series.
func (s NAs) GetError() string {
	return ""
}

// Makes the series nullable.
func (s NAs) MakeNullable() Series {
	return s
}

// Make the series non-nullable.
func (s NAs) MakeNonNullable() Series {
	return s
}

// Returns the type of the series.
func (s NAs) Type() meta.BaseType {
	return meta.NullType
}

// Returns the type and cardinality of the series.
func (s NAs) TypeCard() meta.BaseTypeCard {
	return meta.BaseTypeCard{Base: meta.NullType, Card: s.Len()}
}

// Returns if the series has null values.
func (s NAs) HasNull() bool {
	return true
}

// Returns the number of null values in the series.
func (s NAs) NullCount() int {
	return s.size
}

// Returns if the element at index i is null.
func (s NAs) IsNull(i int) bool {
	return true
}

// Returns the null mask of the series.
func (s NAs) GetNullMask() []bool {
	nullMask := make([]bool, s.size)
	for i := 0; i < s.size; i++ {
		nullMask[i] = true
	}
	return nullMask
}

// Sets the null mask of the series.
func (s NAs) SetNullMask(mask []bool) Series {
	return s
}

// Get the element at index i.
func (s NAs) Get(i int) any {
	return nil
}

func (s NAs) GetAsString(i int) string {
	return gandalff.NA_TEXT
}

// Set the element at index i.
func (s NAs) Set(i int, v any) Series {
	return s
}

// Take the elements according to the given interval.
func (s NAs) Take(params ...int) Series {
	return s
}

// Append elements to the series.
func (s NAs) Append(v any) Series {
	var nullMask []byte
	switch v := v.(type) {
	case nil:
		s.size++
		return s

	case NAs:
		s.size += v.size
		return s

	case bool, gandalff.NullableBool, []bool, []gandalff.NullableBool, Bools:
		var data []bool
		switch v := v.(type) {
		case bool:
			data = make([]bool, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case gandalff.NullableBool:
			data = make([]bool, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []bool:
			data = append(make([]bool, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []gandalff.NullableBool:
			data = make([]bool, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case Bools:
			data = append(make([]bool, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return Bools{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int, gandalff.NullableInt, []int, []gandalff.NullableInt, Ints:
		var data []int
		switch v := v.(type) {
		case int:
			data = make([]int, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case gandalff.NullableInt:
			data = make([]int, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []int:
			data = append(make([]int, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []gandalff.NullableInt:
			data = make([]int, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case Ints:
			data = append(make([]int, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return Ints{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64, gandalff.NullableInt64, []int64, []gandalff.NullableInt64, Int64s:
		var data []int64
		switch v := v.(type) {
		case int64:
			data = make([]int64, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case gandalff.NullableInt64:
			data = make([]int64, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []int64:
			data = append(make([]int64, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []gandalff.NullableInt64:
			data = make([]int64, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case Int64s:
			data = append(make([]int64, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return Int64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64, gandalff.NullableFloat64, []float64, []gandalff.NullableFloat64, Float64s:
		var data []float64
		switch v := v.(type) {
		case float64:
			data = make([]float64, s.size+1)
			data[s.size] = v
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case gandalff.NullableFloat64:
			data = make([]float64, s.size+1)
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data[s.size] = v.Value
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			}

		case []float64:
			data = append(make([]float64, s.size), v...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []gandalff.NullableFloat64:
			data = make([]float64, s.size+len(v))
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = v.Value
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case Float64s:
			data = append(make([]float64, s.size), v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return Float64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string, gandalff.NullableString, []string, []gandalff.NullableString, Strings:
		data := make([]*string, s.size)
		for i := 0; i < s.size; i++ {
			data[i] = s.ctx.StringPool.Put(gandalff.NA_TEXT)
		}

		switch v := v.(type) {
		case string:
			data = append(data, s.ctx.StringPool.Put(v))
			nullMask = __binVecInit(s.size+1, true)
			nullMask[s.size>>3] &= ^(1 << uint(s.size%8))

		case gandalff.NullableString:
			nullMask = __binVecInit(s.size+1, true)
			if v.Valid {
				data = append(data, s.ctx.StringPool.Put(v.Value))
				nullMask[s.size>>3] &= ^(1 << uint(s.size%8))
			} else {
				data = append(data, s.ctx.StringPool.Put(gandalff.NA_TEXT))
			}

		case []string:
			data = append(data, make([]*string, len(v))...)
			for i, v := range v {
				data[s.size+i] = s.ctx.StringPool.Put(v)
			}
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), false, make([]uint8, 0))

		case []gandalff.NullableString:
			data = append(data, make([]*string, len(v))...)
			nullMask = __binVecInit(len(v), false)
			for i, v := range v {
				if v.Valid {
					data[s.size+i] = s.ctx.StringPool.Put(v.Value)
				} else {
					nullMask[i>>3] |= 1 << uint(i%8)
					data[s.size+i] = s.ctx.StringPool.Put(gandalff.NA_TEXT)
				}
			}

			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), len(v), true, nullMask)

		case Strings:
			data = append(data, v.data...)
			_, nullMask = __mergeNullMasks(s.size, true, __binVecInit(s.size, true), v.Len(), v.IsNullable(), v.nullMask)
		}

		return Strings{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return Errors{fmt.Sprintf("NAs.Append: invalid type %T", v)}
	}
}

// All-data accessors.

// Returns the actual data of the series.
func (s NAs) Data() any {
	return make([]bool, s.size)
}

// Returns the nullable data of the series.
func (s NAs) DataAsNullable() any {
	return make([]gandalff.NullableBool, s.size)
}

// Returns the data of the series as a slice of strings.
func (s NAs) DataAsString() []string {
	data := make([]string, s.size)
	for i := 0; i < s.size; i++ {
		data[i] = gandalff.NA_TEXT
	}
	return data
}

// Casts the series to a given type.
func (s NAs) Cast(t meta.BaseType) Series {
	switch t {
	case meta.NullType:
		return s

	case meta.BoolType:
		return Bools{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]bool, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.IntType:
		return Ints{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]int, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.Int64Type:
		return Int64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]int64, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.Float64Type:
		return Float64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]float64, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.StringType:
		return Strings{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]*string, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.TimeType:
		return Times{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]time.Time, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	case meta.DurationType:
		return Durations{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       make([]time.Duration, s.size),
			nullMask:   __binVecInit(s.size, true),
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return Errors{fmt.Sprintf("NAs.Cast: invalid type %s", t.ToString())}
	}
}

// Copies the series.
func (s NAs) Copy() Series {
	return s
}

// Series operations.

// Filters out the elements by the given mask.
// Mask can be a bool series, a slice of bools or a slice of ints.
func (s NAs) Filter(mask any) Series {
	switch mask := mask.(type) {
	case Bools:
		return s.filterBool(mask)
	case []bool:
		return s.filterBoolSlice(mask)
	case []int:
		return s.filterIntSlice(mask, true)
	default:
		return Errors{fmt.Sprintf("NAs.Filter: invalid type %T", mask)}
	}
}

func (s NAs) filterBool(mask Bools) Series {
	elementCount := 0
	for _, v := range mask.data {
		if v {
			elementCount++
		}
	}

	s.size = elementCount
	return s
}

func (s NAs) filterBoolSlice(mask []bool) Series {
	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	s.size = elementCount
	return s
}

func (s NAs) filterIntSlice(indexes []int, check bool) Series {
	// check if indexes are in range
	if check {
		for _, v := range indexes {
			if v < 0 || v >= s.size {
				return Errors{fmt.Sprintf("NAs.Filter: index %d is out of range", v)}
			}
		}
	}

	s.size = len(indexes)
	return s
}

func (s NAs) Map(f gandalff.MapFunc) Series {
	return s
}

func (s NAs) MapNull(f gandalff.MapFuncNull) Series {
	return s
}

type SeriesNAPartition struct {
	partition map[int64][]int
}

func (gp *SeriesNAPartition) getSize() int {
	return len(gp.partition)
}

func (gp *SeriesNAPartition) getMap() map[int64][]int {
	return gp.partition
}

// Group the elements in the series.
func (s NAs) group() Series {
	return s
}

func (s NAs) GroupBy(gp SeriesPartition) Series {
	return s
}

func (s NAs) UnGroup() Series {
	return s
}

func (s NAs) GetPartition() SeriesPartition {
	return s.partition
}

// Sort interface.
func (s NAs) Less(i, j int) bool {
	return false
}

func (s NAs) equal(i, j int) bool {
	return false
}

func (s NAs) Swap(i, j int) {}

func (s NAs) Sort() Series {
	return s
}

func (s NAs) SortRev() Series {
	return s
}
