package series

import (
	"fmt"
	"math"
)

// Not performs logical NOT operation on series
func (s Bools) Not() Series {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = !s.data[i]
	}

	return s
}

func (s Bools) All() bool {
	if s.isNullable {
		for i := 0; i < len(s.data); i++ {
			if s.nullMask[i>>3]&(1<<uint(i%8)) == 0 && !s.data[i] {
				return false
			}
		}

		return true
	} else {
		for i := 0; i < len(s.data); i++ {
			if !s.data[i] {
				return false
			}
		}

		return true
	}
}

func (s Bools) Any() bool {
	if s.isNullable {
		for i := 0; i < len(s.data); i++ {
			if s.nullMask[i>>3]&(1<<uint(i%8)) == 0 && s.data[i] {
				return true
			}
		}

		return false
	} else {
		for i := 0; i < len(s.data); i++ {
			if s.data[i] {
				return true
			}
		}

		return false
	}
}

func (s Bools) And(other any) Series {
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
						result[0] = s.data[0] && o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] && o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] && o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] && o.data[0]
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
							result[i] = s.data[0] && o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
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
							result[i] = s.data[0] && o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] && o.data[i]
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
							result[i] = s.data[i] && o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[0]
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
							result[i] = s.data[i] && o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
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
							result[i] = s.data[i] && o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] && o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot AND %s and %s", s.Type().ToString(), o.Type().ToString())}
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
			return Errors{fmt.Sprintf("Cannot AND %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot AND %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Bools) Or(other any) Series {
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
						result[0] = s.data[0] || o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = s.data[0] || o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = s.data[0] || o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = s.data[0] || o.data[0]
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
							result[i] = s.data[0] || o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
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
							result[i] = s.data[0] || o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[0] || o.data[i]
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
							result[i] = s.data[i] || o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[0]
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
							result[i] = s.data[i] || o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
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
							result[i] = s.data[i] || o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = s.data[i] || o.data[i]
						}
						return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				}
			}
			return Errors{fmt.Sprintf("Cannot OR %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	case NAs:
		if s.Len() == 1 {
			if o.Len() == 1 {
				resultSize := o.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, s.nullMask[0] == 1)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				result[0] = s.data[0]
				return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			} else {
				resultSize := o.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, s.nullMask[0] == 1)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[0]
				}
				return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			}
		} else {
			if o.Len() == 1 {
				resultSize := s.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, false)
					copy(resultNullMask, s.nullMask)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[i]
				}
				return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			} else if s.Len() == o.Len() {
				resultSize := s.Len()
				result := make([]bool, resultSize)
				var resultNullMask []uint8
				if s.isNullable {
					resultNullMask = __binVecInit(resultSize, false)
					copy(resultNullMask, s.nullMask)
				} else {
					resultNullMask = make([]uint8, 0)
				}
				for i := 0; i < resultSize; i++ {
					result[i] = s.data[i]
				}
				return Bools{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
			}
			return Errors{fmt.Sprintf("Cannot OR %s and %s", s.Type().ToString(), o.Type().ToString())}
		}
	default:
		return Errors{fmt.Sprintf("Cannot OR %s and %s", s.Type().ToString(), o.Type().ToString())}
	}

}

func (s Bools) Mul(other any) Series {
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
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] && o.data[0] {
							result[0] = 1
						}
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
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
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
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] && o.data[i] {
								result[i] = 1
							}
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
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[0] {
								result[i] = 1
							}
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
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
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
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] && o.data[i] {
								result[i] = 1
							}
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
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
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
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
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
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
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
						}
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
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
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
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
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						if s.data[0] {
							result[0] = o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						if s.data[0] {
							result[0] = o.data[0]
						}
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
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
							if s.data[0] {
								result[i] = o.data[i]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[0] {
								result[i] = o.data[i]
							}
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
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[0]
							}
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
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
							if s.data[i] {
								result[i] = o.data[i]
							}
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							if s.data[i] {
								result[i] = o.data[i]
							}
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

func (s Bools) Div(other any) Series {
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
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 / b2
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 / b2
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 / b2
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / float64(o.data[0])
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[0])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / float64(o.data[i])
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 / o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 / o.data[i]
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

func (s Bools) Mod(other any) Series {
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
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = math.Mod(b1, b2)
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = math.Mod(b1, b2)
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[0]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = math.Mod(b1, float64(o.data[i]))
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

func (s Bools) Exp(other any) Series {
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
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						b2 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = int64(math.Pow(b1, b2))
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
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
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							b2 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = int64(math.Pow(b1, b2))
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = int64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = int64(math.Pow(b1, float64(o.data[i])))
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = float64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[0])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = float64(math.Pow(b1, float64(o.data[i])))
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

func (s Bools) Add(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 + b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 + b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 + b2
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 + o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 + o.data[i]
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
						result[0] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						result[0] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						result[0] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[0])
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						result[0] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[0])
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
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[i])
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
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[0]) + *o.data[i])
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
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[0])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[0])
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
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[i])
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
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[i])
						}
						return Strings{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]*string, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							result[i] = o.ctx.StringPool.Put(boolToString(s.data[i]) + *o.data[i])
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

func (s Bools) Sub(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 - b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 - b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 - b2
						}
						return Int64s{isNullable: false, nullMask: resultNullMask, data: result, ctx: s.ctx}
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Ints{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Int64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]int64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 - o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
						}
						return Float64s{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]float64, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 - o.data[i]
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

func (s Bools) Eq(other any) Series {
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

func (s Bools) Ne(other any) Series {
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

func (s Bools) Gt(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 > b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 > b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 > b2
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 > o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 > o.data[i]
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

func (s Bools) Ge(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 >= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 >= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 >= b2
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 >= o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 >= o.data[i]
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

func (s Bools) Lt(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 < b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 < b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 < b2
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 < o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 < o.data[i]
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

func (s Bools) Le(other any) Series {
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
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						b2 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						if o.data[0] {
							b2 = 1
						}
						result[0] = b1 <= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[0] {
								b2 = 1
							}
							result[i] = b1 <= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
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
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							b2 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							if o.data[i] {
								b2 = 1
							}
							result[i] = b1 <= b2
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
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := int64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := int64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						b1 := float64(0)
						if s.data[0] {
							b1 = 1
						}
						result[0] = b1 <= o.data[0]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, s.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := o.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[0] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					}
				} else {
					if o.isNullable {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, o.nullMask[0] == 1)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[0]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(resultSize, false)
						copy(resultNullMask, s.nullMask)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
						}
						return Bools{isNullable: true, nullMask: resultNullMask, data: result, ctx: s.ctx}
					} else {
						resultSize := s.Len()
						result := make([]bool, resultSize)
						resultNullMask := __binVecInit(0, false)
						for i := 0; i < resultSize; i++ {
							b1 := float64(0)
							if s.data[i] {
								b1 = 1
							}
							result[i] = b1 <= o.data[i]
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
