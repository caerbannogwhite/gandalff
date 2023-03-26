package gandalff

import (
	"fmt"
	"typesys"
)

// GDLSeriesFloat64 represents a series of floats.
type GDLSeriesFloat64 struct {
	isNullable bool
	name       string
	data       []float64
	nullMask   []uint8
}

func NewGDLSeriesFloat64(name string, isNullable bool, makeCopy bool, data []float64) GDLSeriesFloat64 {
	var nullMask []uint8
	if isNullable {
		if len(data)%8 == 0 {
			nullMask = make([]uint8, (len(data) >> 3))
		} else {
			nullMask = make([]uint8, (len(data)>>3)+1)
		}
	} else {
		nullMask = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GDLSeriesFloat64{isNullable: isNullable, name: name, data: data, nullMask: nullMask}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GDLSeriesFloat64) Len() int {
	return len(s.data)
}

func (s GDLSeriesFloat64) IsNullable() bool {
	return s.isNullable
}

func (s GDLSeriesFloat64) MakeNullable() GDLSeries {
	if !s.isNullable {
		s.isNullable = true
		if len(s.data)%8 == 0 {
			s.nullMask = make([]uint8, (len(s.data) >> 3))
		} else {
			s.nullMask = make([]uint8, (len(s.data)>>3)+1)
		}
	}
	return s
}

func (s GDLSeriesFloat64) Name() string {
	return s.name
}

func (s GDLSeriesFloat64) Type() typesys.BaseType {
	return typesys.Float64Type
}

func (s GDLSeriesFloat64) HasNull() bool {
	for _, v := range s.nullMask {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GDLSeriesFloat64) NullCount() int {
	count := 0
	for _, v := range s.nullMask {
		for i := 0; i < 8; i++ {
			if v&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s GDLSeriesFloat64) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMask[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GDLSeriesFloat64) SetNull(i int) error {
	if s.isNullable {
		s.nullMask[i/8] |= 1 << uint(i%8)
		return nil
	}
	return fmt.Errorf("GDLSeriesFloat64.SetNull: series is not nullable")
}

func (s GDLSeriesFloat64) GetNullMask() []bool {
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

func (s GDLSeriesFloat64) SetNullMask(mask []bool) error {
	if !s.isNullable {
		return fmt.Errorf("GDLSeriesFloat64.SetNullMask: series is not nullable")
	}

	for k, v := range mask {
		if v {
			s.nullMask[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMask[k/8] &= ^(1 << uint(k%8))
		}
	}

	return nil
}

func (s GDLSeriesFloat64) Get(i int) interface{} {
	return s.data[i]
}

func (s GDLSeriesFloat64) Set(i int, v interface{}) {
	s.data[i] = v.(float64)
}

func (s GDLSeriesFloat64) Append(v interface{}) GDLSeries {
	switch v := v.(type) {
	case float64:
		return s.AppendRaw(v)
	case []float64:
		return s.AppendRaw(v)
	case NullableFloat64:
		return s.AppendNullable(v)
	case []NullableFloat64:
		return s.AppendNullable(v)
	case GDLSeriesFloat64:
		return s.AppendSeries(v)
	case GDLSeriesError:
		return v
	default:
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.Append: invalid type, %T", v)}
	}
}

// Append appends a value or a slice of values to the series.
func (s GDLSeriesFloat64) AppendRaw(v interface{}) GDLSeries {
	if s.isNullable {
		if f, ok := v.(float64); ok {
			s.data = append(s.data, f)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, 0)
			}
		} else if fv, ok := v.([]float64); ok {
			s.data = append(s.data, fv...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
			}
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	} else {
		if f, ok := v.(float64); ok {
			s.data = append(s.data, f)
		} else if fv, ok := v.([]float64); ok {
			s.data = append(s.data, fv...)
		} else {
			return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendRaw: invalid type %T", v)}
		}
	}
	return s
}

// AppendNullable appends a nullable value or a slice of nullable values to the series.
func (s GDLSeriesFloat64) AppendNullable(v interface{}) GDLSeries {
	if !s.isNullable {
		return GDLSeriesError{"GDLSeriesFloat64.AppendNullable: series is not nullable"}
	}

	if f, ok := v.(NullableFloat64); ok {
		s.data = append(s.data, f.Value)
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !f.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}
	} else if fv, ok := v.([]NullableFloat64); ok {
		if len(s.data) > len(s.nullMask)<<8 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for _, f := range fv {
			s.data = append(s.data, f.Value)
			if !f.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}
	} else {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendNullable: invalid type %T", v)}
	}

	return s
}

func (s GDLSeriesFloat64) AppendSeries(other GDLSeries) GDLSeries {
	var ok bool
	var o GDLSeriesFloat64
	if o, ok = other.(GDLSeriesFloat64); !ok {
		return GDLSeriesError{fmt.Sprintf("GDLSeriesFloat64.AppendSeries: invalid type %T", other)}
	}

	if s.isNullable {
		if o.isNullable {
			s.data = append(s.data, o.data...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
			}

			// merge null masks
			sIdx := len(s.data) - len(o.data)
			oIdx := 0
			for _, v := range o.nullMask {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		} else {
			s.data = append(s.data, o.data...)
			if len(s.data) > len(s.nullMask)<<3 {
				s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
			}
		}
	} else {
		if o.isNullable {
			s.data = append(s.data, o.data...)
			if len(s.data)%8 == 0 {
				s.nullMask = make([]uint8, (len(s.data) >> 3))
			} else {
				s.nullMask = make([]uint8, (len(s.data)>>3)+1)
			}
			s.isNullable = true

			// merge null masks
			sIdx := len(s.data) - len(o.data)
			oIdx := 0
			for _, v := range o.nullMask {
				for j := 0; j < 8; j++ {
					if v&(1<<uint(j)) != 0 {
						s.nullMask[sIdx>>3] |= 1 << uint(sIdx%8)
					}
					sIdx++
					oIdx++
				}
			}
		} else {
			s.data = append(s.data, o.data...)
		}
	}

	return s
}

///////////////////////////////		ALL DATA ACCESSORS			/////////////////////////

func (s GDLSeriesFloat64) Data() interface{} {
	return s.data
}

func (s GDLSeriesFloat64) NullableData() interface{} {
	data := make([]NullableFloat64, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat64{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GDLSeriesFloat64) StringData() []string {
	data := make([]string, len(s.data))
	for i, v := range s.data {
		if s.IsNull(i) {
			data[i] = NULL_STRING
		} else {
			data[i] = floatToString(v)
		}
	}
	return data
}

func (s GDLSeriesFloat64) Copy() GDLSeries {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return GDLSeriesFloat64{isNullable: s.isNullable, name: s.name, data: data, nullMask: s.nullMask}
}

///////////////////////////////		SERIES OPERATIONS			/////////////////////////

// FilterByMask returns a new series with elements filtered by the mask.
func (s GDLSeriesFloat64) FilterByMask(mask []bool) GDLSeries {
	data := make([]float64, 0)
	nullMask := make([]uint8, len(s.nullMask))
	for i, v := range mask {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMask[i/8] |= 1 << uint(i%8)
			}
		}
	}
	return GDLSeriesFloat64{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

func (s GDLSeriesFloat64) FilterByIndeces(indices []int) GDLSeries {
	data := make([]float64, len(indices))
	nullMask := make([]uint8, len(s.nullMask))
	for i, v := range indices {
		data[i] = s.data[v]
		if s.isNullable {
			nullMask[i/8] |= 1 << uint(i%8)
		}
	}
	return GDLSeriesFloat64{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

///////////////////////////////		GROUPING OPERATIONS			/////////////////////////

func (s GDLSeriesFloat64) Group() GDLSeriesPartition {
	return nil
}

func (s GDLSeriesFloat64) SubGroup(gp GDLSeriesPartition) GDLSeriesPartition {
	return nil
}