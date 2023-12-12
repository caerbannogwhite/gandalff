package gandalff

import (
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"

	"github.com/caerbannogwhite/preludiometa"
)

type typeBucket struct {
	nullCount   int
	boolCount   int
	intCount    int
	floatCount  int
	stringCount int
}

// Get the most common type in the bucket and whether it is the only type
func (tb *typeBucket) getMostCommonType() (preludiometa.BaseType, bool) {
	if tb.boolCount > tb.intCount && tb.boolCount > tb.floatCount && tb.boolCount > tb.stringCount {
		return preludiometa.BoolType, tb.nullCount+tb.intCount+tb.floatCount+tb.stringCount == 0
	} else if tb.intCount > tb.floatCount && tb.intCount > tb.stringCount {
		return preludiometa.Int64Type, tb.nullCount+tb.boolCount+tb.floatCount+tb.stringCount == 0
	} else if tb.floatCount > tb.stringCount {
		return preludiometa.Float64Type, tb.nullCount+tb.boolCount+tb.intCount+tb.stringCount == 0
	}
	return preludiometa.StringType, tb.nullCount+tb.boolCount+tb.intCount+tb.floatCount == 0
}

type typeGuesser struct {
	nullValues     bool
	nullRegex      *regexp.Regexp
	boolRegex      *regexp.Regexp
	boolTrueRegex  *regexp.Regexp
	boolFalseRegex *regexp.Regexp
	intRegex       *regexp.Regexp
	floatRegex     *regexp.Regexp

	// For each column, count the number of values that match each type
	typeBuckets []typeBucket
}

// Get the regexes for guessing data types
func newTypeGuesser(nullValues bool) typeGuesser {
	return typeGuesser{
		nullValues,
		regexp.MustCompile(`^([Nn][Uu][Ll][Ll])$|^([Nn][Aa][Nn]?)$|^([Nn]/[Aa])$|^$`),
		regexp.MustCompile(`^([Tt]([Rr][Uu][Ee])?)|([Ff]([Aa][Ll][Ss][Ee])?)$`),
		regexp.MustCompile(`^[Tt]([Rr][Uu][Ee])?$`),
		regexp.MustCompile(`^[Ff]([Aa][Ll][Ss][Ee])?$`),
		regexp.MustCompile(`^[-+]?[0-9]+$`),
		regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`),
		nil,
	}
}

func (tg *typeGuesser) setLength(length int) {
	tg.typeBuckets = make([]typeBucket, length)
}

func (tg *typeGuesser) guessType(record string) preludiometa.BaseType {
	if tg.boolRegex.MatchString(record) {
		return preludiometa.BoolType
	} else if tg.intRegex.MatchString(record) {
		return preludiometa.Int64Type
	} else if tg.floatRegex.MatchString(record) {
		return preludiometa.Float64Type
	}
	return preludiometa.StringType
}

func (tg *typeGuesser) guessTypes(records []string) {
	for i, v := range records {
		if tg.boolRegex.MatchString(v) {
			tg.typeBuckets[i].boolCount++
		} else if tg.intRegex.MatchString(v) {
			tg.typeBuckets[i].intCount++
		} else if tg.floatRegex.MatchString(v) {
			tg.typeBuckets[i].floatCount++
		} else {
			tg.typeBuckets[i].stringCount++
		}
	}
}

func (tg *typeGuesser) guessTypesNulls(records []string) {
	for i, v := range records {
		if tg.boolRegex.MatchString(v) {
			tg.typeBuckets[i].boolCount++
		} else if tg.intRegex.MatchString(v) {
			tg.typeBuckets[i].intCount++
		} else if tg.floatRegex.MatchString(v) {
			tg.typeBuckets[i].floatCount++
		} else if tg.nullRegex.MatchString(v) {
			tg.typeBuckets[i].nullCount++
		} else {
			tg.typeBuckets[i].stringCount++
		}
	}
}

func (tg typeGuesser) getTypes() []preludiometa.BaseType {
	types := make([]preludiometa.BaseType, len(tg.typeBuckets))
	if tg.nullValues {
		for i, v := range tg.typeBuckets {
			types[i], _ = v.getMostCommonType()
		}
	} else {
		var onlyType bool
		for i, v := range tg.typeBuckets {
			types[i], onlyType = v.getMostCommonType()
			if !onlyType {
				types[i] = preludiometa.StringType
			}
		}
	}
	return types
}

func (tg typeGuesser) atoBool(s string) (bool, error) {
	if tg.boolTrueRegex.MatchString(s) {
		return true, nil
	} else if tg.boolFalseRegex.MatchString(s) {
		return false, nil
	}
	return false, fmt.Errorf("cannot convert \"%s\" to bool", s)
}

type RowDataProvider interface {
	Read() ([]string, error)
}

func readRowData(reader RowDataProvider, nullValues bool, guessDataTypeLen int, schema *preludiometa.Schema, ctx *Context) ([]Series, error) {
	var dataTypes []preludiometa.BaseType
	var recordsForGuessing [][]string

	// Initialize TypeGuesser
	tg := newTypeGuesser(nullValues)

	// Guess data types
	if schema == nil {
		recordsForGuessing = make([][]string, guessDataTypeLen)

		// Read first record to get length
		record, err := reader.Read()
		if err != nil && err != io.EOF {
			return nil, err
		}
		recordsForGuessing[0] = record

		tg.setLength(len(record))

		if nullValues {
			tg.guessTypesNulls(record)
			for i := 1; i < guessDataTypeLen; i++ {
				record, err := reader.Read()
				if err != nil {
					if err == io.EOF {
						break
					}
					return nil, err
				}
				recordsForGuessing[i] = record
				tg.guessTypesNulls(record)
			}
		} else {
			tg.guessTypes(record)
			for i := 1; i < guessDataTypeLen; i++ {
				record, err := reader.Read()
				if err != nil {
					if err == io.EOF {
						break
					}
					return nil, err
				}
				recordsForGuessing[i] = record
				tg.guessTypes(record)
			}
		}
		dataTypes = tg.getTypes()
	} else

	// Use schema
	{
		dataTypes = schema.GetDataTypes()
	}

	nullMasks := make([][]bool, len(dataTypes))
	if nullValues {
		for i := range nullMasks {
			nullMasks[i] = make([]bool, 0)
		}
	}

	values := make([]interface{}, len(dataTypes))
	for i := range values {
		switch dataTypes[i] {
		case preludiometa.BoolType:
			values[i] = make([]bool, 0)
		case preludiometa.IntType:
			values[i] = make([]int, 0)
		case preludiometa.Int64Type:
			values[i] = make([]int64, 0)
		case preludiometa.Float64Type:
			values[i] = make([]float64, 0)
		case preludiometa.StringType:
			values[i] = make([]*string, 0)
		}
	}

	// If no schema: add records for guessing to values
	if schema == nil {
		if nullValues {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case preludiometa.BoolType:
						if b, err := tg.atoBool(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
							values[i] = append(values[i].([]bool), false)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]bool), b)
						}

					case preludiometa.IntType:
						if d, err := strconv.Atoi(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
							values[i] = append(values[i].([]int), 0)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int), int(d))
						}

					case preludiometa.Int64Type:
						if d, err := strconv.ParseInt(v, 10, 64); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
							values[i] = append(values[i].([]int64), 0)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int64), d)
						}

					case preludiometa.Float64Type:
						if f, err := strconv.ParseFloat(v, 64); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
							values[i] = append(values[i].([]float64), math.NaN())
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]float64), f)
						}

					case preludiometa.StringType:
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
					}
				}
			}
		} else {
			for _, record := range recordsForGuessing {
				for i, v := range record {
					switch dataTypes[i] {
					case preludiometa.BoolType:
						b, err := tg.atoBool(v)
						if err != nil {
							return nil, err
						}
						values[i] = append(values[i].([]bool), b)

					case preludiometa.IntType:
						d, err := strconv.Atoi(v)
						if err != nil {
							return nil, err
						}
						values[i] = append(values[i].([]int), int(d))

					case preludiometa.Int64Type:
						d, err := strconv.ParseInt(v, 10, 64)
						if err != nil {
							return nil, err
						}
						values[i] = append(values[i].([]int64), d)

					case preludiometa.Float64Type:
						f, err := strconv.ParseFloat(v, 64)
						if err != nil {
							return nil, err
						}
						values[i] = append(values[i].([]float64), f)

					case preludiometa.StringType:
						values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
					}
				}
			}
		}
	}

	if nullValues {
		for {
			record, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err
			}

			for i, v := range record {
				switch dataTypes[i] {
				case preludiometa.BoolType:
					if b, err := tg.atoBool(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
						values[i] = append(values[i].([]bool), false)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]bool), b)
					}

				case preludiometa.IntType:
					if d, err := strconv.Atoi(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
						values[i] = append(values[i].([]int), 0)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int), int(d))
					}

				case preludiometa.Int64Type:
					if d, err := strconv.ParseInt(v, 10, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
						values[i] = append(values[i].([]int64), 0)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int64), d)
					}

				case preludiometa.Float64Type:
					if f, err := strconv.ParseFloat(v, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
						values[i] = append(values[i].([]float64), math.NaN())
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]float64), f)
					}

				case preludiometa.StringType:
					nullMasks[i] = append(nullMasks[i], false)
					values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
				}
			}
		}
	} else {
		for {
			record, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err
			}

			for i, v := range record {
				switch dataTypes[i] {
				case preludiometa.BoolType:
					b, err := tg.atoBool(v)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]bool), b)

				case preludiometa.IntType:
					d, err := strconv.Atoi(v)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]int), int(d))

				case preludiometa.Int64Type:
					d, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]int64), d)

				case preludiometa.Float64Type:
					f, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return nil, err
					}
					values[i] = append(values[i].([]float64), f)

				case preludiometa.StringType:
					values[i] = append(values[i].([]*string), ctx.stringPool.Put(v))
				}
			}
		}
	}

	// Create series
	series := make([]Series, len(dataTypes))
	for i := range dataTypes {
		switch dataTypes[i] {
		case preludiometa.BoolType:
			series[i] = NewSeriesBool(values[i].([]bool), nullMasks[i], false, ctx)

		case preludiometa.IntType:
			series[i] = NewSeriesInt(values[i].([]int), nullMasks[i], false, ctx)

		case preludiometa.Int64Type:
			series[i] = NewSeriesInt64(values[i].([]int64), nullMasks[i], false, ctx)

		case preludiometa.Float64Type:
			series[i] = NewSeriesFloat64(values[i].([]float64), nullMasks[i], false, ctx)

		case preludiometa.StringType:
			series[i] = SeriesString{
				isNullable: nullValues,
				data:       values[i].([]*string),
				nullMask:   __binVecFromBools(nullMasks[i]),
				ctx:        ctx,
			}
		}
	}

	return series, nil
}
