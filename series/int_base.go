package series

import (
	"fmt"
	"time"

	"github.com/caerbannogwhite/gandalff"
	"github.com/caerbannogwhite/gandalff/meta"
)

func (s Ints) printInfo() {
	fmt.Println("Ints")
	fmt.Println("==========")
	fmt.Println("IsNullable:", s.isNullable)
	fmt.Println("Sorted:    ", s.sorted)
	fmt.Println("Data:      ", s.data)
	fmt.Println("NullMask:  ", s.nullMask)
	fmt.Println("Partition: ", s.partition)
	fmt.Println("Context:   ", s.ctx)
}

////////////////////////			BASIC ACCESSORS

// Return the context of the series.
func (s Ints) GetContext() *gandalff.Context {
	return s.ctx
}

// Return the number of elements in the series.
func (s Ints) Len() int {
	return len(s.data)
}

// Return the type of the series.
func (s Ints) Type() meta.BaseType {
	return meta.IntType
}

// Return the type and cardinality of the series.
func (s Ints) TypeCard() meta.BaseTypeCard {
	return meta.BaseTypeCard{Base: meta.IntType, Card: s.Len()}
}

// Return if the series is grouped.
func (s Ints) IsGrouped() bool {
	return s.partition != nil
}

// Return if the series admits null values.
func (s Ints) IsNullable() bool {
	return s.isNullable
}

// Return if the series is sorted.
func (s Ints) IsSorted() gandalff.SeriesSortOrder {
	return s.sorted
}

// Return if the series is error.
func (s Ints) IsError() bool {
	return false
}

// Return the error message of the series.
func (s Ints) GetError() string {
	return ""
}

// Return the partition of the series.
func (s Ints) GetPartition() SeriesPartition {
	return s.partition
}

// Return if the series has null values.
func (s Ints) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

// Return the number of null values in the series.
func (s Ints) NullCount() int {
	count := 0
	for _, x := range s.nullMask {
		for ; x != 0; x >>= 1 {
			count += int(x & 1)
		}
	}
	return count
}

// Return if the element at index i is null.
func (s Ints) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i>>3]&(1<<uint(i%8)) != 0
	}
	return false
}

// Return the null mask of the series.
func (s Ints) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

// Set the null mask of the series.
func (s Ints) SetNullMask(mask []bool) Series {
	if s.partition != nil {
		return Errors{"Ints.SetNullMask: cannot set values on a grouped series"}
	}

	if s.isNullable {
		for k, v := range mask {
			if v {
				s.nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				s.nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}
		return s
	} else {
		nullMask := __binVecInit(len(s.data), false)
		for k, v := range mask {
			if v {
				nullMask[k>>3] |= 1 << uint(k%8)
			} else {
				nullMask[k>>3] &= ^(1 << uint(k%8))
			}
		}

		s.isNullable = true
		s.nullMask = nullMask

		return s
	}
}

// Make the series nullable.
func (s Ints) MakeNullable() Series {
	if !s.isNullable {
		s.isNullable = true
		s.nullMask = __binVecInit(len(s.data), false)
	}
	return s
}

// Make the series non-nullable.
func (s Ints) MakeNonNullable() Series {
	if s.isNullable {
		s.isNullable = false
		s.nullMask = make([]uint8, 0)
	}
	return s
}

// Get the element at index i.
func (s Ints) Get(i int) any {
	return s.data[i]
}

// Append appends a value or a slice of values to the series.
func (s Ints) Append(v any) Series {
	if s.partition != nil {
		return Errors{"Ints.Append: cannot append values to a grouped series"}
	}

	switch v := v.(type) {
	case nil:
		s.data = append(s.data, 0)
		s = s.MakeNullable().(Ints)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		s.nullMask[(len(s.data)-1)>>3] |= 1 << uint8((len(s.data)-1)%8)

	case NAs:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, v.Len(), true, __binVecInit(v.Len(), true))
		s.data = append(s.data, make([]int, v.Len())...)

	case int:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []int:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case gandalff.NullableInt:
		s.data = append(s.data, v.Value)
		s = s.MakeNullable().(Ints)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[(len(s.data)-1)>>3] |= 1 << uint8((len(s.data)-1)%8)
		}

	case []gandalff.NullableInt:
		ssize := len(s.data)
		s.data = append(s.data, make([]int, len(v))...)
		s = s.MakeNullable().(Ints)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			s.data[ssize+i] = b.Value
			if !b.Valid {
				s.nullMask[(ssize+i)>>3] |= 1 << uint8((ssize+i)%8)
			}
		}

	case Ints:
		if s.ctx != v.ctx {
			return Errors{"Ints.Append: cannot append Ints from different contexts"}
		}

		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, v.data...)

	default:
		return Errors{fmt.Sprintf("Ints.Append: invalid type %T", v)}
	}

	s.sorted = gandalff.SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s Ints) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("Ints", s.Len(), params...)
	if err != nil {
		return Errors{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Return the elements of the series as a slice.
func (s Ints) Data() any {
	return s.data
}

// Copy the series.
func (s Ints) Copy() Series {
	data := make([]int, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return Ints{
		isNullable: s.isNullable,
		sorted:     s.sorted,
		data:       data,
		nullMask:   nullMask,
		partition:  s.partition,
		ctx:        s.ctx,
	}
}

func (s Ints) getData() []int {
	return s.data
}

// Ungroup the series.
func (s Ints) UnGroup() Series {
	s.partition = nil
	return s
}

////////////////////////			FILTER OPERATIONS

// Filters out the elements by the given mask.
// Mask can be Bools, Ints, bool slice or a int slice.
func (s Ints) Filter(mask any) Series {
	switch mask := mask.(type) {
	case Bools:
		return s.filterBoolSlice(mask.data)
	case Ints:
		return s.filterIntSlice(mask.data, true)
	case []bool:
		return s.filterBoolSlice(mask)
	case []int:
		return s.filterIntSlice(mask, true)
	default:
		return Errors{fmt.Sprintf("Ints.Filter: invalid type %T", mask)}
	}
}

func (s Ints) filterBoolSlice(mask []bool) Series {
	if len(mask) != len(s.data) {
		return Errors{fmt.Sprintf("Ints.Filter: mask length (%d) does not match series length (%d)", len(mask), len(s.data))}
	}

	elementCount := 0
	for _, v := range mask {
		if v {
			elementCount++
		}
	}

	var data []int
	var nullMask []uint8

	data = make([]int, elementCount)

	if s.isNullable {
		nullMask = __binVecInit(elementCount, false)
		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				if srcIdx%8 > dstIdx%8 {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
				} else {
					nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
				}
				dstIdx++
			}
		}
	} else {
		nullMask = make([]uint8, 0)
		dstIdx := 0
		for srcIdx, v := range mask {
			if v {
				data[dstIdx] = s.data[srcIdx]
				dstIdx++
			}
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

func (s Ints) filterIntSlice(indexes []int, check bool) Series {
	if len(indexes) == 0 {
		s.data = make([]int, 0)
		s.nullMask = make([]uint8, 0)
		return s
	}

	// check if indexes are in range
	if check {
		for _, v := range indexes {
			if v < 0 || v >= len(s.data) {
				return Errors{fmt.Sprintf("Ints.Filter: index %d is out of range", v)}
			}
		}
	}

	var data []int
	var nullMask []uint8

	size := len(indexes)
	data = make([]int, size)

	if s.isNullable {
		nullMask = __binVecInit(size, false)
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
			if srcIdx%8 > dstIdx%8 {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) >> uint(srcIdx%8-dstIdx%8))
			} else {
				nullMask[dstIdx>>3] |= ((s.nullMask[srcIdx>>3] & (1 << uint(srcIdx%8))) << uint(dstIdx%8-srcIdx%8))
			}
		}
	} else {
		nullMask = make([]uint8, 0)
		for dstIdx, srcIdx := range indexes {
			data[dstIdx] = s.data[srcIdx]
		}
	}

	s.data = data
	s.nullMask = nullMask

	return s
}

// Apply the given function to each element of the series.
func (s Ints) Map(f gandalff.MapFunc) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:
		data := make([]bool, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(bool)
		}

		return Bools{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int:
		data := make([]int, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int)
		}

		return Ints{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64:
		data := make([]int64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int64)
		}

		return Int64s{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64:
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(float64)
		}

		return Float64s{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string:
		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = s.ctx.StringPool.Put(f(s.data[i]).(string))
		}

		return Strings{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Time:
		data := make([]time.Time, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(time.Time)
		}

		return Times{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Duration:
		data := make([]time.Duration, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(time.Duration)
		}

		return Durations{
			isNullable: s.isNullable,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   s.nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return Errors{fmt.Sprintf("Ints.Map: Unsupported type %T", v)}
	}
}

// Apply the given function to each element of the series.
func (s Ints) MapNull(f gandalff.MapFuncNull) Series {
	if len(s.data) == 0 {
		return s
	}

	if !s.isNullable {
		return Errors{"Ints.MapNull: series is not nullable"}
	}

	v, isNull := f(s.Get(0), s.IsNull(0))
	switch v.(type) {
	case bool:
		data := make([]bool, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(bool)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Bools{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int:
		data := make([]int, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(int)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Ints{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case int64:
		data := make([]int64, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(int64)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Int64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case float64:
		data := make([]float64, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(float64)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Float64s{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case string:
		data := make([]*string, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = s.ctx.StringPool.Put(v.(string))
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Strings{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Time:
		data := make([]time.Time, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(time.Time)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Times{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	case time.Duration:
		data := make([]time.Duration, len(s.data))
		nullMask := make([]uint8, len(s.nullMask))
		for i := 0; i < len(s.data); i++ {
			v, isNull = f(s.data[i], s.IsNull(i))
			data[i] = v.(time.Duration)
			if isNull {
				nullMask[i>>3] |= 1 << uint(i%8)
			}
		}

		return Durations{
			isNullable: true,
			sorted:     gandalff.SORTED_NONE,
			data:       data,
			nullMask:   nullMask,
			partition:  nil,
			ctx:        s.ctx,
		}

	default:
		return Errors{fmt.Sprintf("Ints.MapNull: Unsupported type %T", v)}
	}
}
