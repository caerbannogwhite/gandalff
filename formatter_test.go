package gandalff

import (
	"fmt"
	"math"
	"testing"
)

func Test_Format_01(t *testing.T) {

	row1 := []float64{1, 1.1, -1.1, 0.0, math.NaN()}
	row2 := []float64{16e62, 16e-64, -999, 999, 0}
	row3 := []float64{1.0000000000000001e+09, 1.000000000000001e+09, 1.00000000000001e+09, 1.0000000000001e+09, math.NaN()}
	row4 := []float64{1.1e-05, 1.2e-05, 1.3e-05, 1.4e-05, 1.5e-05}
	row5 := []float64{1.1e-08, 1.2e-08, 1.3e-08, 1.4e-08, 1.5e-08}
	row6 := []float64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000, 1000000000000000}

	f := NewNumericFormatter()
	for _, row := range [][]float64{row1, row2, row3, row4, row5} {
		for _, num := range row {
			f.Push(num)
		}
	}

	f = NewNumericFormatter()
	for _, num := range row6 {
		f.Push(num)
	}

	for _, num := range row6 {
		fmt.Println(f.Format(num))
	}

}
