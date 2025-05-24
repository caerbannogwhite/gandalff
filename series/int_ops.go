package series

import (
	"fmt"
	"math"
)

func (s Ints) Neg() Series {
	for i, v := range s.data {
		s.data[i] = -v
	}
	return s
}

func (s Ints) And(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return Errors{fmt.Sprintf("Cannot AND %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Or(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	default:
		return Errors{fmt.Sprintf("Cannot OR %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Mul(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						if o.data[0] {
							result[0] = s.data[0]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[0]
							}
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if o.data[0] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if o.data[i] {
								result[i] = s.data[i]
							}
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] * o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] * o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] * o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] * o.data[0]
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] * o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[0]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] * o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) * o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) * o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) * o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) * o.data[0]
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) * o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[0]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) * o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) * o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) * o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) * o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) * o.data[0]
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) * o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[0]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) * o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case NAs:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return NAs{size: resultSize}
			} else {
				resultSize := o.Len()
				return NAs{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return NAs{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return NAs{size: resultSize}
			}
			return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot multiply %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Div(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = float64(s.data[0]) / b2
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[0]) / b2
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = float64(s.data[i]) / b2
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) / float64(o.data[0])
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / float64(o.data[i])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[0])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / float64(o.data[i])
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) / o.data[0]
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) / o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[0]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) / o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot divide %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Mod(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(float64(s.data[0]), b2)
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[0]), b2)
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(float64(s.data[i]), b2)
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = math.Mod(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Mod(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot use modulo %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Exp(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := float64(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(float64(s.data[0]), b2))
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[0]), b2))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := float64(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(float64(s.data[i]), b2))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(math.Pow(float64(s.data[0]), float64(o.data[0])))
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[0]), float64(o.data[i])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[0])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(math.Pow(float64(s.data[i]), float64(o.data[i])))
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = math.Pow(float64(s.data[0]), float64(o.data[0]))
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[0]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[0]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = math.Pow(float64(s.data[i]), float64(o.data[i]))
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot use exponentiation %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Add(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] + b2
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] + b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] + b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] + o.data[0]
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] + o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[0]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] + o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) + o.data[0]
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) + o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[0]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) + o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) + o.data[0]
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) + o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[0]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) + o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Strings:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[0])
						return Strings{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[0])) + *o.data[i])
						}
						return Strings{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[0])
						}
						return Strings{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(intToString(int64(s.data[i])) + *o.data[i])
						}
						return Strings{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case NAs:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return NAs{size: resultSize}
			} else {
				resultSize := o.Len()
				return NAs{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return NAs{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return NAs{size: resultSize}
			}
			return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot sum %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Sub(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] - b2
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] - b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] - b2
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] - o.data[0]
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] - o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[0]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] - o.data[i]
						}
						return Ints{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) - o.data[0]
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) - o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[0]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) - o.data[i]
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) - o.data[0]
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) - o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
		} else {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVS(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[0]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			} else if s.Len() == o.Len() {
				if s.isNullable {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrVV(s.nullMask, o.nullMask, resultNullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) - o.data[i]
						}
						return Float64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot subtract %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Eq(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] == o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) == o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) == o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) == o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) == o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case NAs:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return NAs{size: resultSize}
			} else {
				resultSize := o.Len()
				return NAs{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return NAs{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return NAs{size: resultSize}
			}
			return Errors{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for equality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Ne(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] != o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) != o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) != o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) != o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) != o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case NAs:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				return NAs{size: resultSize}
			} else {
				resultSize := o.Len()
				return NAs{size: resultSize}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				return NAs{size: resultSize}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				return NAs{size: resultSize}
			}
			return Errors{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for inequality %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Gt(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] > b2
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] > b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] > b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] > o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) > o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) > o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) > o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for greater than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Ge(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] >= b2
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] >= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] >= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] >= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) >= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) >= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) >= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for greater than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Lt(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] < b2
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] < b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] < b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] < o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) < o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) < o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) < o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for less than %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Ints) Le(other any) Series {
	var otherSeries Series
	if _, ok := other.(Series); ok {
		otherSeries = other.(Series)
	} else {
		otherSeries = NewSeries(other, nil, false, false, s.ctx)
	}
	if s.ctx != otherSeries.GetContext() {
		return Errors{fmt.Sprintf("Cannot operate on series with different contexts: %v and %v", s.ctx, otherSeries.GetContext())}
	}
	switch o := otherSeries.(type) {
	case Bools:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b2 := int(0)
						if o.data[0] {
							b2 = 1
						}
						result[0] = s.data[0] <= b2
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[0] <= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[0] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b2 := int(0)
							if o.data[i] {
								b2 = 1
							}
							result[i] = s.data[i] <= b2
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Ints:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = s.data[0] <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] <= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Int64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = int64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = int64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = int64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = int64(s.data[0]) <= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = int64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = int64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case Float64s:
		if s.Len() == 1 {
			if o.Len() == 1 {
				if s.isNullable {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						__binVecOrSS(s.nullMask, o.nullMask, resultNullMask)
						result[0] = float64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = float64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = float64(s.data[0]) <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = float64(s.data[0]) <= o.data[0]
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[0]) <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[0]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, o.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = float64(s.data[i]) <= o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot compare for less than or equal to %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}
