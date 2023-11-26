package gandalff

import (
	"fmt"
	"io"
	"regexp"
	"strconv"

	"github.com/caerbannogwhite/preludiometa"
)

type typeGuesser struct {
	boolRegex      *regexp.Regexp
	boolTrueRegex  *regexp.Regexp
	boolFalseRegex *regexp.Regexp
	intRegex       *regexp.Regexp
	floatRegex     *regexp.Regexp
}

// Get the regexes for guessing data types
func newTypeGuesser() typeGuesser {
	boolRegex := regexp.MustCompile(`^([Tt]([Rr][Uu][Ee])?)|([Ff]([Aa][Ll][Ss][Ee])?)$`)
	intRegex := regexp.MustCompile(`^[-+]?[0-9]+$`)
	floatRegex := regexp.MustCompile(`^[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?$`)

	boolTrueRegex := regexp.MustCompile(`^[Tt]([Rr][Uu][Ee])?$`)
	boolFalseRegex := regexp.MustCompile(`^[Ff]([Aa][Ll][Ss][Ee])?$`)

	return typeGuesser{boolRegex, boolTrueRegex, boolFalseRegex, intRegex, floatRegex}
}

func (tg typeGuesser) guessType(s string) preludiometa.BaseType {
	if tg.boolRegex.MatchString(s) {
		return preludiometa.BoolType
	} else if tg.intRegex.MatchString(s) {
		return preludiometa.Int64Type
	} else if tg.floatRegex.MatchString(s) {
		return preludiometa.Float64Type
	}
	return preludiometa.StringType
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
	tg := newTypeGuesser()

	// Guess data types
	if schema == nil {
		recordsForGuessing = make([][]string, guessDataTypeLen)

		for i := 0; i < guessDataTypeLen; i++ {
			record, err := reader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
			}
			recordsForGuessing[i] = record

			for j, v := range record {
				if i == 0 {
					dataTypes = append(dataTypes, tg.guessType(v))
				} else {
					if dataTypes[j] == preludiometa.StringType {
						continue
					}
					if tg.guessType(v) != dataTypes[j] {
						dataTypes[j] = preludiometa.StringType
					}
				}
			}
		}
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
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]bool), b)
						}

					case preludiometa.IntType:
						if d, err := strconv.Atoi(v); err != nil {
							nullMasks[i] = append(nullMasks[i], true)
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int), int(d))
						}

					case preludiometa.Int64Type:
						if d, err := strconv.ParseInt(v, 10, 64); err != nil {
							return nil, err
						} else {
							nullMasks[i] = append(nullMasks[i], false)
							values[i] = append(values[i].([]int64), d)
						}

					case preludiometa.Float64Type:
						if f, err := strconv.ParseFloat(v, 64); err != nil {
							return nil, err
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
			}

			for i, v := range record {
				switch dataTypes[i] {
				case preludiometa.BoolType:
					if b, err := tg.atoBool(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]bool), b)
					}

				case preludiometa.IntType:
					if d, err := strconv.Atoi(v); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int), int(d))
					}

				case preludiometa.Int64Type:
					if d, err := strconv.ParseInt(v, 10, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
					} else {
						nullMasks[i] = append(nullMasks[i], false)
						values[i] = append(values[i].([]int64), d)
					}

				case preludiometa.Float64Type:
					if f, err := strconv.ParseFloat(v, 64); err != nil {
						nullMasks[i] = append(nullMasks[i], true)
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
