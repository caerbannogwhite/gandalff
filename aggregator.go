package gandalff

import (
	"fmt"
)

type aggregatorBuilder struct {
	df          BaseDataFrame
	removeNAs   bool
	aggregators []Aggregator
}

func (ab aggregatorBuilder) RemoveNAs(b bool) aggregatorBuilder {
	ab.removeNAs = b
	return ab
}

func (ab aggregatorBuilder) Run() DataFrame {
	df := ab.df
	if df.err != nil {
		return ab.df
	}

	if len(ab.aggregators) == 0 {
		return df
	}

	// CHECK: aggregators must have unique names and names must be valid
	aggNames := make(map[string]bool)
	for _, agg := range ab.aggregators {
		if aggNames[agg.GetNewName()] {
			df.err = fmt.Errorf("BaseDataFrame.Agg: aggregator names must be unique")
			return df
		}
		aggNames[agg.GetNewName()] = true

		// CASE: aggregator count has a default name
		if df.__series(agg.GetColName()) == nil {
			if agg_, ok := agg.(internalAggregator); ok && agg_.type_ == AGGREGATE_COUNT {
			} else {
				df.err = fmt.Errorf("BaseDataFrame.Agg: series \"%s\" not found", agg.GetColName())
				return df
			}
		}

		if ab.removeNAs {
			agg.RemoveNAs(true)
		}
	}

	var result DataFrame
	if df.isGrouped {
		var indeces [][]int
		var flatIndeces []int
		result, indeces, flatIndeces, _ = df.groupHelper()

		groupsNum := len(indeces)

		var series Series
		for _, agg := range ab.aggregators {
			series = df.__series(agg.GetColName())

			// INTERNAL AGGREGATORS
			if agg_, ok := agg.(internalAggregator); ok {
				switch agg_.type_ {
				case AGGREGATE_COUNT:
					counts := make([]int64, groupsNum)
					for i, group := range indeces {
						counts[i] = int64(len(group))
					}
					result = result.AddSeries(agg_.newName, NewSeriesInt64(counts, nil, false, df.ctx))

				case AGGREGATE_SUM:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_sum(dataF64, flatIndeces, groupsNum, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MIN:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_min(dataF64, flatIndeces, groupsNum, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MAX:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_max(dataF64, flatIndeces, groupsNum, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MEAN:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_mean(dataF64, flatIndeces, groupsNum, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_STD:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_std(dataF64, flatIndeces, groupsNum, agg_.removeNAs), nil, false, df.ctx))
				}
			} else
			// CUSTOM AGGREGATORS
			{
				dataF64 := __gdl_stats_preprocess(series)
				for i, idx := range flatIndeces {
					agg.Reduce(idx, nil, dataF64[i], false)
				}
			}
		}

		// var wg sync.WaitGroup
		// wg.Add(THREADS_NUMBER)

		// buffer := make(chan __stats_thread_data)
		// for i := 0; i < THREADS_NUMBER; i++ {
		// 	go __stats_worker(&wg, buffer)
		// }

		// for _, agg := range aggregators {
		// 	series := df.__series(agg.name)

		// 	resultData := make([]float64, len(*indeces))
		// 	result = result.AddSeries(agg.name, NewSeriesFloat64(resultData, nil, false, df.ctx))
		// 	for gi, group := range *indeces {
		// 		buffer <- __stats_thread_data{
		// 			op:      agg.type_,
		// 			gi:      gi,
		// 			indeces: group,
		// 			series:  series,
		// 			res:     resultData,
		// 		}
		// 	}
		// }

		// close(buffer)
		// wg.Wait()

	} else {
		result = NewBaseDataFrame(df.ctx)

		var series Series
		for _, agg := range ab.aggregators {
			series = df.__series(agg.GetColName())

			// INTERNAL AGGREGATORS
			if agg_, ok := agg.(internalAggregator); ok {
				switch agg_.type_ {
				case AGGREGATE_COUNT:
					result = result.AddSeries(agg_.newName, NewSeriesInt64([]int64{int64(df.NRows())}, nil, false, df.ctx))

				case AGGREGATE_SUM:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_sum(dataF64, nil, 1, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MIN:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_min(dataF64, nil, 1, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MAX:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_max(dataF64, nil, 1, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_MEAN:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_mean(dataF64, nil, 1, agg_.removeNAs), nil, false, df.ctx))

				case AGGREGATE_STD:
					dataF64 := __gdl_stats_preprocess(series)
					result = result.AddSeries(agg_.newName, NewSeriesFloat64(__gdl_std(dataF64, nil, 1, agg_.removeNAs), nil, false, df.ctx))
				}
			} else
			// CUSTOM AGGREGATORS
			{
			}
		}
	}

	return result
}

type internalAggregatorType int8

const (
	AGGREGATE_COUNT internalAggregatorType = iota
	AGGREGATE_SUM
	AGGREGATE_MEAN
	AGGREGATE_MEDIAN
	AGGREGATE_MIN
	AGGREGATE_MAX
	AGGREGATE_STD
)

const DEFAULT_COUNT_NAME = "n"

type Aggregator interface {
	RemoveNAs(bool) Aggregator
	GetColName() string
	NewName(string) Aggregator
	GetNewName() string
	Reduce(group int, result, value interface{}, isNA bool)
}

type internalAggregator struct {
	removeNAs bool
	colName   string
	newName   string
	type_     internalAggregatorType
}

func (agg internalAggregator) RemoveNAs(b bool) Aggregator {
	agg.removeNAs = b
	return agg
}

func (agg internalAggregator) GetColName() string {
	return agg.colName
}

func (agg internalAggregator) NewName(name string) Aggregator {
	agg.newName = name
	return agg
}

func (agg internalAggregator) GetNewName() string {
	return agg.newName
}

func (agg internalAggregator) Reduce(group int, result, value interface{}, isNA bool) {

}

func Count() Aggregator {
	return internalAggregator{false, DEFAULT_COUNT_NAME, DEFAULT_COUNT_NAME, AGGREGATE_COUNT}
}

func Sum(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("sum(%s)", colName), AGGREGATE_SUM}
}

func Mean(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("mean(%s)", colName), AGGREGATE_MEAN}
}

func Median(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("median(%s)", colName), AGGREGATE_MEDIAN}
}

func Min(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("min(%s)", colName), AGGREGATE_MIN}
}

func Max(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("max(%s)", colName), AGGREGATE_MAX}
}

func Std(colName string) Aggregator {
	return internalAggregator{false, colName, fmt.Sprintf("std(%s)", colName), AGGREGATE_STD}
}
