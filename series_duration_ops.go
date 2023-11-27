package gandalff

import (
	"fmt"
	"time"
)

func (s SeriesDuration) Mul(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Div(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Mod(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Exp(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return SeriesError{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Add(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesString:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[0])
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[0].String() + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[0])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.stringPool.Put(s.data[i].String() + *o.data[i])
						}
						return SeriesString{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesTime:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = o.data[0].Add(s.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.data[0].Add(s.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.data[0].Add(s.data[0])
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.data[0].Add(s.data[0])
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[0])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[0])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[0].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[0].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[0].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[0].Add(s.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[i])
						}
						return SeriesTime{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Time, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.data[i].Add(s.data[i])
						}
						return SeriesTime{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] + o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] + o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] + o.data[0]
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Sub(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] - o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] - o.data[0]
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] - o.data[0]
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesDuration{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]time.Duration, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return SeriesDuration{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Eq(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] == o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Ne(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] != o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case SeriesNA:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			} else {
				resultSize := o.Len()
				return SeriesNA{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return SeriesNA{size: resultSize}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Gt(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] > o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Ge(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] >= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Lt(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] < o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s SeriesDuration) Le(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return SeriesError{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case SeriesDuration:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] <= o.data[0]
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return SeriesBool{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return SeriesError{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
