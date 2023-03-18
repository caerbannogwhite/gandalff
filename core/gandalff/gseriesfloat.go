package gandalff

// GSeriesFloat represents a series of floats.
type GSeriesFloat struct {
	isNullable bool
	name       string
	data       []float64
	nullMap    []uint8
}

func NewGSeriesFloat(name string, isNullable bool, makeCopy bool, data []float64) GSeriesFloat {
	var nullMap []uint8
	if isNullable {
		nullMap = make([]uint8, len(data)/8+1)
	} else {
		nullMap = make([]uint8, 0)
	}

	if makeCopy {
		actualData := make([]float64, len(data))
		copy(actualData, data)
		data = actualData
	}

	return GSeriesFloat{isNullable: isNullable, name: name, data: data, nullMap: nullMap}
}

///////////////////////////////		BASIC ACCESSORS			/////////////////////////////////

func (s GSeriesFloat) Len() int {
	return len(s.data)
}

func (s GSeriesFloat) IsNullable() bool {
	return s.isNullable
}

func (s GSeriesFloat) Name() string {
	return s.name
}

func (s GSeriesFloat) Type() GSeriesType {
	return FloatType
}

func (s GSeriesFloat) HasNull() bool {
	for _, v := range s.nullMap {
		if v != 0 {
			return true
		}
	}
	return false
}

func (s GSeriesFloat) NullCount() int {
	count := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8; i++ {
			if v&(1<<uint(i)) != 0 {
				count++
			}
		}
	}
	return count
}

func (s GSeriesFloat) IsNull(i int) bool {
	if s.isNullable {
		return s.nullMap[i/8]&(1<<uint(i%8)) != 0
	}
	return false
}

func (s GSeriesFloat) SetNull(i int) {
	if s.isNullable {
		s.nullMap[i/8] |= 1 << uint(i%8)
	}
}

func (s GSeriesFloat) GetNullMask() []bool {
	mask := make([]bool, len(s.data))
	idx := 0
	for _, v := range s.nullMap {
		for i := 0; i < 8 && idx < len(s.data); i++ {
			mask[idx] = v&(1<<uint(i)) != 0
			idx++
		}
	}
	return mask
}

func (s GSeriesFloat) SetNullMask(mask []bool) {
	for k, v := range mask {
		if v {
			s.nullMap[k/8] |= 1 << uint(k%8)
		} else {
			s.nullMap[k/8] &= ^(1 << uint(k%8))
		}
	}
}

func (s GSeriesFloat) Get(i int) interface{} {
	return s.data[i]
}

func (s GSeriesFloat) Set(i int, v interface{}) {
	s.data[i] = v.(float64)
}

///////////////////////////////		ALL DATA ACCESSORS			/////////////////////////

func (s GSeriesFloat) Data() interface{} {
	return s.data
}

func (s GSeriesFloat) NullableData() interface{} {
	data := make([]NullableFloat, len(s.data))
	for i, v := range s.data {
		data[i] = NullableFloat{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s GSeriesFloat) StringData() []string {
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

func (s GSeriesFloat) Copy() GSeries {
	data := make([]float64, len(s.data))
	copy(data, s.data)
	nullMap := make([]uint8, len(s.nullMap))
	copy(nullMap, s.nullMap)

	return GSeriesFloat{isNullable: s.isNullable, name: s.name, data: data, nullMap: s.nullMap}
}

///////////////////////////////		SERIES OPERATIONS			/////////////////////////

func (s GSeriesFloat) Filter(mask []bool) GSeries {
	data := make([]float64, 0)
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range mask {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}
	}
	return GSeriesFloat{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

func (s GSeriesFloat) FilterInPlace(mask []bool) {
	data := make([]float64, 0)
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range mask {
		if v {
			data = append(data, s.data[i])
			if s.isNullable {
				nullMap[i/8] |= 1 << uint(i%8)
			}
		}
	}
	s.data = data
	s.nullMap = nullMap
}

func (s GSeriesFloat) FilterByIndex(index []int) GSeries {
	data := make([]float64, len(index))
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range index {
		data[i] = s.data[v]
		if s.isNullable {
			nullMap[i/8] |= 1 << uint(i%8)
		}
	}
	return GSeriesFloat{isNullable: s.isNullable, name: s.name, data: data, nullMap: nullMap}
}

func (s GSeriesFloat) FilterByIndexInPlace(index []int) {
	data := make([]float64, len(index))
	nullMap := make([]uint8, len(s.nullMap))
	for i, v := range index {
		data[i] = s.data[v]
		if s.isNullable {
			nullMap[i/8] |= 1 << uint(i%8)
		}
	}
	s.data = data
	s.nullMap = nullMap
}