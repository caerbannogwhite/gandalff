package gandalff

import (
	"fmt"
	"sort"
	"sync"
	"typesys"
)

// SeriesInt32 represents a series of ints.
type SeriesInt32 struct {
	isGrouped  bool
	isNullable bool
	sorted     SeriesSortOrder
	name       string
	data       []int32
	nullMask   []uint8
	pool       *StringPool
	partition  *SeriesInt32Partition
}

// Get the element at index i as a string.
func (s SeriesInt32) GetString(i int) string {
	if s.isNullable && s.IsNull(i) {
		return NULL_STRING
	}
	return intToString(int64(s.data[i]))
}

// Set the element at index i. The value v can be any belonging to types:
// int8, int16, int, int32, int64 and their nullable versions.
func (s SeriesInt32) Set(i int, v any) Series {
	switch val := v.(type) {
	case int8:
		s.data[i] = int32(val)

	case int16:
		s.data[i] = int32(val)

	case int:
		s.data[i] = int32(val)

	case int32:
		s.data[i] = int32(val)

	case NullableInt8:
		s.MakeNullable()
		if v.(NullableInt8).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt16:
		s.MakeNullable()
		if v.(NullableInt16).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	case NullableInt32:
		s.MakeNullable()
		if v.(NullableInt32).Valid {
			s.data[i] = int32(val.Value)
		} else {
			s.data[i] = 0
			s.nullMask[i>>3] |= 1 << uint(i%8)
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Set: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

// Take the elements according to the given interval.
func (s SeriesInt32) Take(params ...int) Series {
	indeces, err := seriesTakePreprocess("SeriesInt32", s.Len(), params...)
	if err != nil {
		return SeriesError{err.Error()}
	}
	return s.filterIntSlice(indeces, false)
}

// Append appends a value or a slice of values to the series.
func (s SeriesInt32) Append(v any) Series {
	switch v := v.(type) {
	case int32:
		s.data = append(s.data, v)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}

	case []int32:
		s.data = append(s.data, v...)
		if s.isNullable && len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask))...)
		}

	case NullableInt32:
		s.data = append(s.data, v.Value)
		s.isNullable = true
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, 0)
		}
		if !v.Valid {
			s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
		}

	case []NullableInt32:
		ssize := len(s.data)
		s.data = append(s.data, make([]int32, len(v))...)
		s.isNullable = true
		if len(s.data) > len(s.nullMask)<<3 {
			s.nullMask = append(s.nullMask, make([]uint8, (len(s.data)>>3)-len(s.nullMask)+1)...)
		}
		for i, b := range v {
			s.data[ssize+i] = b.Value
			if !b.Valid {
				s.nullMask[len(s.data)>>3] |= 1 << uint(len(s.data)%8)
			}
		}

	case SeriesInt32:
		s.isNullable, s.nullMask = __mergeNullMasks(len(s.data), s.isNullable, s.nullMask, len(v.data), v.isNullable, v.nullMask)
		s.data = append(s.data, v.data...)

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Append: invalid type %T", v)}
	}

	s.sorted = SORTED_NONE
	return s
}

////////////////////////			ALL DATA ACCESSORS

func (s SeriesInt32) Int32s() []int32 {
	return s.data
}

func (s SeriesInt32) Data() any {
	return s.data
}

func (s SeriesInt32) DataAsNullable() any {
	data := make([]NullableInt32, len(s.data))
	for i, v := range s.data {
		data[i] = NullableInt32{Valid: !s.IsNull(i), Value: v}
	}
	return data
}

func (s SeriesInt32) DataAsString() []string {
	data := make([]string, len(s.data))
	if s.isNullable {
		for i, v := range s.data {
			if s.IsNull(i) {
				data[i] = NULL_STRING
			} else {
				data[i] = intToString(int64(v))
			}
		}
	} else {
		for i, v := range s.data {
			data[i] = intToString(int64(v))
		}
	}
	return data
}

// Casts the series to a given type.
func (s SeriesInt32) Cast(t typesys.BaseType) Series {
	switch t {
	case typesys.BoolType:
		data := make([]bool, len(s.data))
		for i, v := range s.data {
			data[i] = v != 0
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.Int32Type:
		return s

	case typesys.Int64Type:
		data := make([]int64, len(s.data))
		for i, v := range s.data {
			data[i] = int64(v)
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.Float64Type:
		data := make([]float64, len(s.data))
		for i, v := range s.data {
			data[i] = float64(v)
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case typesys.StringType:
		if s.pool == nil {
			return SeriesError{"SeriesInt32.Cast: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		if s.isNullable {
			for i, v := range s.data {
				if s.IsNull(i) {
					data[i] = s.pool.Put(NULL_STRING)
				} else {
					data[i] = s.pool.Put(intToString(int64(v)))
				}
			}
		} else {
			for i, v := range s.data {
				data[i] = s.pool.Put(intToString(int64(v)))
			}
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	default:
		return SeriesError{fmt.Sprintf("SeriesInt32.Cast: invalid type %s", t.ToString())}
	}
}

func (s SeriesInt32) Copy() Series {
	data := make([]int32, len(s.data))
	copy(data, s.data)
	nullMask := make([]uint8, len(s.nullMask))
	copy(nullMask, s.nullMask)

	return SeriesInt32{isNullable: s.isNullable, name: s.name, data: data, nullMask: nullMask}
}

func (s SeriesInt32) getDataPtr() *[]int32 {
	return &s.data
}

func (s SeriesInt32) Map(f GDLMapFunc) Series {
	if len(s.data) == 0 {
		return s
	}

	v := f(s.Get(0))
	switch v.(type) {
	case bool:

		data := make([]bool, len(s.data))
		chunkLen := len(s.data) / THREADS_NUMBER
		if chunkLen < MINIMUM_PARALLEL_SIZE_2 {
			for i := 0; i < len(s.data); i++ {
				data[i] = f(s.data[i]).(bool)
			}
		} else {
			var wg sync.WaitGroup
			wg.Add(THREADS_NUMBER)

			for n := 0; n < THREADS_NUMBER; n++ {
				start := n * chunkLen
				end := (n + 1) * chunkLen
				if n == THREADS_NUMBER-1 {
					end = len(s.data)
				}

				go func() {
					for i := start; i < end; i++ {
						data[i] = f(s.data[i]).(bool)
					}
					wg.Done()
				}()
			}

			wg.Wait()
		}

		return SeriesBool{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case int32:
		data := make([]int32, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int32)
		}

		s.isGrouped = false
		s.sorted = SORTED_NONE
		s.data = data
		s.partition = nil

		return s

	case int64:
		data := make([]int64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(int64)
		}

		return SeriesInt64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case float64:
		data := make([]float64, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = f(s.data[i]).(float64)
		}

		return SeriesFloat64{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}

	case string:
		if s.pool == nil {
			return SeriesError{"SeriesInt32.Map: StringPool is nil"}
		}

		data := make([]*string, len(s.data))
		for i := 0; i < len(s.data); i++ {
			data[i] = s.pool.Put(f(s.data[i]).(string))
		}

		return SeriesString{
			isGrouped:  false,
			isNullable: s.isNullable,
			sorted:     SORTED_NONE,
			name:       s.name,
			data:       data,
			nullMask:   s.nullMask,
			pool:       s.pool,
			partition:  nil,
		}
	}

	return SeriesError{fmt.Sprintf("SeriesInt32.Map: Unsupported type %T", v)}
}

////////////////////////			GROUPING OPERATIONS

type SeriesInt32Partition struct {
	partition           map[int64][]int
	isDense             bool
	partitionDenseMin   int32
	partitionDense      [][]int
	partitionDenseNulls []int
}

func (gp *SeriesInt32Partition) getSize() int {
	if gp.isDense {
		if gp.partitionDenseNulls != nil && len(gp.partitionDenseNulls) > 0 {
			return len(gp.partitionDense) + 1
		}
		return len(gp.partitionDense)
	}
	return len(gp.partition)
}

func (gp *SeriesInt32Partition) getMap() map[int64][]int {
	if gp.isDense {
		map_ := make(map[int64][]int, len(gp.partitionDense))
		for i, part := range gp.partitionDense {
			map_[int64(i)+int64(gp.partitionDenseMin)] = part
		}

		// Merge the nulls to the map
		if gp.partitionDenseNulls != nil && len(gp.partitionDenseNulls) > 0 {
			nullKey := __series_get_nullkey(map_, HASH_NULL_KEY)
			map_[nullKey] = gp.partitionDenseNulls
		}

		return map_
	}

	return gp.partition
}

func (s SeriesInt32) group() Series {
	var useDenseMap bool
	var min, max int32
	var partition SeriesInt32Partition

	// If the number of elements is small,
	// look for the minimum and maximum values
	if len(s.data) < MINIMUM_PARALLEL_SIZE_2 {
		useDenseMap = true
		max = s.data[0]
		min = s.data[0]
		for _, v := range s.data {
			if v > max {
				max = v
			}
			if v < min {
				min = v
			}
		}
	}

	// If the difference between the maximum and minimum values is acceptable,
	// then we can use a dense map, otherwise we use a sparse map
	if useDenseMap && (max-min >= MINIMUM_PARALLEL_SIZE_1) {
		useDenseMap = false
	}

	// DENSE MAP
	if useDenseMap {
		var nulls []int
		map_ := make([][]int, max-min+1)
		for i := 0; i < len(map_); i++ {
			map_[i] = make([]int, 0, DEFAULT_DENSE_MAP_ARRAY_INITIAL_CAPACITY)
		}

		if s.HasNull() {
			nulls = make([]int, 0, DEFAULT_DENSE_MAP_ARRAY_INITIAL_CAPACITY)
			for i, v := range s.data {
				if s.IsNull(i) {
					nulls = append(nulls, i)
				} else {
					map_[v-min] = append(map_[v-min], i)
				}
			}
		} else {
			for i, v := range s.data {
				map_[v-min] = append(map_[v-min], i)
			}
		}

		partition = SeriesInt32Partition{
			isDense:             true,
			partitionDenseMin:   min,
			partitionDense:      map_,
			partitionDenseNulls: nulls,
		}
	} else

	// SPARSE MAP
	{
		// Define the worker callback
		worker := func(threadNum, start, end int, map_ map[int64][]int) {
			up := end - ((end - start) % 8)
			for i := start; i < up; {
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				i++
			}

			for i := up; i < end; i++ {
				map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
			}
		}

		// Define the worker callback for nulls
		workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
			for i := start; i < end; i++ {
				if s.IsNull(i) {
					(*nulls) = append((*nulls), i)
				} else {
					map_[int64(s.data[i])] = append(map_[int64(s.data[i])], i)
				}
			}
		}

		partition = SeriesInt32Partition{
			isDense: false,
			partition: __series_groupby(
				THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_2, len(s.data), s.HasNull(),
				worker, workerNulls),
		}
	}

	s.isGrouped = true
	s.partition = &partition

	return s
}

func (s SeriesInt32) GroupBy(partition SeriesPartition) Series {
	if partition == nil {
		return s
	}

	// collect all keys
	otherIndeces := partition.getMap()
	keys := make([]int64, len(otherIndeces))
	i := 0
	for k := range otherIndeces {
		keys[i] = k
		i++
	}

	// Define the worker callback
	worker := func(threadNum, start, end int, map_ map[int64][]int) {
		var newHash int64
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				newHash = int64(s.data[index]) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	// Define the worker callback for nulls
	workerNulls := func(threadNum, start, end int, map_ map[int64][]int, nulls *[]int) {
		var newHash int64
		for _, h := range keys[start:end] { // keys is defined outside the function
			for _, index := range otherIndeces[h] { // otherIndeces is defined outside the function
				if s.IsNull(index) {
					newHash = HASH_MAGIC_NUMBER_NULL + (h << 13) + (h >> 4)
				} else {
					newHash = int64(s.data[index]) + HASH_MAGIC_NUMBER + (h << 13) + (h >> 4)
				}
				map_[newHash] = append(map_[newHash], index)
			}
		}
	}

	newPartition := SeriesInt32Partition{
		partition: __series_groupby(
			THREADS_NUMBER, MINIMUM_PARALLEL_SIZE_2, len(keys), s.HasNull(),
			worker, workerNulls),
	}

	s.isGrouped = true
	s.partition = &newPartition

	return s
}

func (s SeriesInt32) UnGroup() Series {
	s.isGrouped = false
	s.partition = nil
	return s
}

func (s SeriesInt32) GetPartition() SeriesPartition {
	return s.partition
}

////////////////////////			SORTING OPERATIONS

func (s SeriesInt32) Less(i, j int) bool {
	if s.isNullable {
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 {
			return false
		}
		if s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			return true
		}
	}

	return s.data[i] < s.data[j]
}

func (s SeriesInt32) equal(i, j int) bool {
	if s.isNullable {
		if (s.nullMask[i>>3] & (1 << uint(i%8))) > 0 {
			return (s.nullMask[j>>3] & (1 << uint(j%8))) > 0
		}
		if (s.nullMask[j>>3] & (1 << uint(j%8))) > 0 {
			return false
		}
	}

	return s.data[i] == s.data[j]
}

func (s SeriesInt32) Swap(i, j int) {
	if s.isNullable {
		// i is null, j is not null
		if s.nullMask[i>>3]&(1<<uint(i%8)) > 0 && s.nullMask[j>>3]&(1<<uint(j%8)) == 0 {
			s.nullMask[i>>3] &= ^(1 << uint(i%8))
			s.nullMask[j>>3] |= 1 << uint(j%8)
		} else

		// i is not null, j is null
		if s.nullMask[i>>3]&(1<<uint(i%8)) == 0 && s.nullMask[j>>3]&(1<<uint(j%8)) > 0 {
			s.nullMask[i>>3] |= 1 << uint(i%8)
			s.nullMask[j>>3] &= ^(1 << uint(j%8))
		}
	}

	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s SeriesInt32) Sort() Series {
	if s.sorted != SORTED_ASC {
		sort.Sort(s)
		s.sorted = SORTED_ASC
	}
	return s
}

func (s SeriesInt32) SortRev() Series {
	if s.sorted != SORTED_DESC {
		sort.Sort(sort.Reverse(s))
		s.sorted = SORTED_DESC
	}
	return s
}
