## GANDALFF: Golang, ANother DAta Library For Fun 🧙‍♂️

Or, for short, GDL: Golang Data Library

### What is it?

Gandalff is a library for data wrangling in Go.
The goal is to provide a simple and efficient API for data manipulation in Go,
similar to Pandas or Polars in Python, and Dplyr in R.
It supports nullable types: null data is optimized for memory usage.

### Examples

```go
package main

import (
	"strings"

	gandalff "github.com/caerbannogwhite/gandalff"
)

func main() {
	data1 := `
name,age,weight,junior,department,salary band
Alice C,29,75.0,F,HR,4
John Doe,30,80.5,true,IT,2
Bob,31,85.0,F,IT,4
Jane H,25,60.0,false,IT,4
Mary,28,70.0,false,IT,3
Oliver,32,90.0,true,HR,1
Ursula,27,65.0,f,Business,4
Charlie,33,60.0,t,Business,2
Megan,26,55.0,F,IT,3
`

	gandalff.NewBaseDataFrame(gandalff.NewContext()).
		FromCSV().
		SetReader(strings.NewReader(data1)).
		Read().
		Select("department", "age", "weight", "junior").
		GroupBy("department").
		Agg(gandalff.Min("age"), gandalff.Max("weight"), gandalff.Mean("junior"), gandalff.Count()).
		PrettyPrint(
      gandalff.NewPrettyPrintParams().
			  SetUseLipGloss(true))
}

// Output:
// ╭────────────┬─────────┬─────────┬─────────┬───────╮
// │ department │ age     │ weight  │ junior  │ n     │
// ├────────────┼─────────┼─────────┼─────────┼───────┤
// │ String     │ Float64 │ Float64 │ Float64 │ Int64 │
// ├────────────┼─────────┼─────────┼─────────┼───────┤
// │ HR         │   29.00 │   90.00 │  0.5000 │ 2.000 │
// │ IT         │   25.00 │   85.00 │  0.5000 │ 4.000 │
// │ Business   │   27.00 │   65.00 │  0.5000 │ 2.000 │
// ╰────────────┴─────────┴─────────┴─────────┴───────╯
```

### Community

You can join the [Gandalff community on Discord](https://discord.gg/vPv5bhXY).

### Supported data types

The data types not checked are not yet supported, but might be in the future.

- [x] Bool
- [ ] Bool (memory optimized, not fully implemented yet)
- [ ] Int16
- [x] Int
- [x] Int64
- [ ] Float32
- [x] Float64
- [ ] Complex64
- [ ] Complex128
- [x] String
- [x] Time
- [x] Duration

### Supported operations for Series

- [x] Filter

  - [x] filter by bool slice
  - [x] filter by int slice
  - [x] filter by bool series
  - [x] filter by int series

- [x] Group

  - [x] Group (with nulls)
  - [x] SubGroup (with nulls)

- [x] Map
- [x] Sort

  - [x] Sort (with nulls)
  - [x] SortRev (with nulls)

- [x] Take

### Supported operations for DataFrame

- [x] Agg
- [x] Filter
- [x] GroupBy
- [ ] Join

  - [x] Inner
  - [x] Left
  - [x] Right
  - [x] Outer
  - [ ] Inner with nulls
  - [ ] Left with nulls
  - [ ] Right with nulls
  - [ ] Outer with nulls

- [ ] Map
- [x] OrderBy
- [x] Select
- [x] Take
- [ ] Pivot
- [ ] Stack/Append

### Supported stats functions

- [x] Count
- [x] Sum
- [x] Mean
- [ ] Median
- [x] Min
- [x] Max
- [x] StdDev
- [ ] Variance
- [ ] Quantile

### TODO

- [ ] Improve dataframe PrettyPrint: add parameters, optimize data display, use lipgloss.
- [ ] Implement string factors.
- [ ] SeriesTime: set time format.
- [ ] Implement `Set(i []int, v []any) Series`.
- [ ] Add `Slice(i []int) Series` (using filter?).
- [ ] Implement memory optimized Bool series with uint64.
- [ ] Use uint64 for null mask.
- [ ] Implement chunked series.
- [ ] Implement JSON reader and writer.
- [ ] Implement Parquet reader and writer.
- [ ] Implement SPSS reader and writer.
- [ ] Implement SAS7BDAT reader and writer (https://cran.r-project.org/web/packages/sas7bdat/vignettes/sas7bdat.pdf)
