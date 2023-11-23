package main

import (
	"fmt"
	. "gandalff"
	"strings"
)

const (
	data1 = `
name,age,weight,junior,department,salary band
Alice C,29,75.0,F,HR,4
John Doe,30,80.5,true,IT,2
Bob,31,85.0,T,IT,4
Jane H,25,60.0,false,IT,4
Mary,28,70.0,false,IT,3
Oliver,32,90.0,true,HR,1
Ursula,27,65.0,f,Business,4
Charlie,33,60.0,t,Business,2
`

	data2 = `
department,number of employees,budget
IT,4,100000
HR,2,50000
Business,2,50000
Operations,4,250000
`
)

func Example01() {
	// f, _ := os.Create("test.csv")

	NewBaseDataFrame(NewContext()).
		FromCsv().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		Read().
		Select("department", "age", "weight", "junior").
		GroupBy("department").
		Agg(Min("age"), Max("weight"), Mean("junior"), Count()).
		PrettyPrint(
			NewPrettyPrintParams().
				SetUseLipGloss(true).
				SetWidth(130).
				SetNRows(50)).
		ToXlsx().
		SetPath("../testdata/test.xlsx").
		SetSheet("test").
		Write()

	// ToMarkdown().
	// SetWriter(f).
	// SetHeader(true).
	// SetIndex(false).
	// SetNullString("").
	// Write()

	// ToCsv().
	// SetDelimiter('\t').
	// SetWriter(f).
	// Write()

	// Output:
	// ╭────────────┬────────────┬────────────┬────────────┬────────────╮
	// │ department │        age │     weight │     junior │          n │
	// ├────────────┼────────────┼────────────┼────────────┼────────────┤
	// │     String │    Float64 │    Float64 │    Float64 │      Int64 │
	// ├────────────┼────────────┼────────────┼────────────┼────────────┤
	// │         HR │         29 │         90 │        0.5 │          2 │
	// │         IT │         25 │         85 │        0.5 │          4 │
	// │   Business │         27 │         65 │        0.5 │          2 │
	// ╰────────────┴────────────┴────────────┴────────────┴────────────╯
}

func Example02() {
	ctx := NewContext()
	ppp := NewPrettyPrintParams()

	employees := NewBaseDataFrame(ctx).
		FromCsv().
		SetReader(strings.NewReader(data1)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	departments := NewBaseDataFrame(ctx).
		FromCsv().
		SetReader(strings.NewReader(data2)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	departments.PrettyPrint(ppp)

	employees.Join(LEFT_JOIN, departments, "department").
		PrettyPrint(ppp)
}

// func Example03() {
// 	s := NewSeriesInt64("nums", true, false, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}).
// 		SetNullMask([]bool{true, false, false, true, false, false, true, false, false}).(SeriesInt64).
// 		Mul(NewSeriesInt32("nums2", false, false, []int32{1, 2, 3, 4, 5, 6, 7, 8, 9})).(SeriesInt64).
// 		Add(NewSeriesInt64("scalar", false, false, []int64{1})).
// 		Gt(NewSeriesFloat64("nums", false, false, []float64{20}))

// 	p := NewSeriesString("nums", true, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}, NewStringPool()).
// 		SetNullMask([]bool{true, false, false, true, false, false, true, false, false}).(SeriesString).
// 		Add(NewSeriesString("nums2", false, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}, NewStringPool()))

// 	fmt.Println(s.Data())
// 	fmt.Println(s.GetNullMask())

// 	fmt.Println(p.Data())
// 	fmt.Println(p.GetNullMask())
// }

func Example04() {
	x := `
a,b
1,2
2,2
3,3
3,4
4,4
`

	y := `
a,b
1,2
2,2
2,3
3,3
4,4
`

	ctx := NewContext()
	ppp := NewPrettyPrintParams()

	dfX := NewBaseDataFrame(ctx).
		FromCsv().
		SetReader(strings.NewReader(x)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	dfY := NewBaseDataFrame(ctx).
		FromCsv().
		SetReader(strings.NewReader(y)).
		SetDelimiter(',').
		SetHeader(true).
		Read()

	dfX.Join(INNER_JOIN, dfY, "a", "b").
		PrettyPrint(ppp)

	dfX.Join(LEFT_JOIN, dfY, "a", "b").
		PrettyPrint(ppp)

	dfX.Join(RIGHT_JOIN, dfY, "a", "b").
		PrettyPrint(ppp)

	dfX.Join(OUTER_JOIN, dfY, "a", "b").
		PrettyPrint(ppp)
}

func Example05() {

	NewBaseDataFrame(NewContext()).
		FromXpt().
		// SetPath("../testdata/CDBRFS90.XPT").
		SetPath("../testdata/xpt_test_mixed.xpt").
		SetVersion(XPT_VERSION_9).
		// SetMaxObservations(10).
		Read().
		// Take(100).

		// to SAS XPT
		ToXpt().
		SetPath("../testdata/mixed_out.XPT").
		SetVersion(XPT_VERSION_9).
		Write().

		// to Excel
		ToXlsx().
		SetPath("../testdata/test.xlsx").
		SetSheet("test").
		SetNaText("").
		Write().

		// to HTML
		ToHtml().
		SetPath("../testdata/test.html").
		SetNaText("-").
		SetNewLine("\n").
		SetIndent("  ").
		Write().

		// to JSON
		ToJson().
		SetPath("../testdata/test.json").
		Write().

		// Pretty print
		PrettyPrint(
			NewPrettyPrintParams().
				SetUseLipGloss(true).
				SetWidth(200).
				SetNRows(20))
}

// Excel read and write
func Example06() {
	NewBaseDataFrame(NewContext()).
		FromXlsx().
		SetPath("").
		SetSheet("").
		SetHeader(3).
		Read().
		PrettyPrint(
			NewPrettyPrintParams().
				SetUseLipGloss(true).
				SetWidth(200).
				SetNRows(50))

}

// JSON read and write
func Example07() {
	NewBaseDataFrame(NewContext()).
		FromJson().
		SetPath("../testdata/test.json").
		Read().
		PrettyPrint(
			NewPrettyPrintParams().
				SetUseLipGloss(true).
				SetWidth(200).
				SetNRows(50))

}

func main() {
	// fmt.Println("Example01:")
	// Example01()

	// fmt.Println("Example02:")
	// Example02()

	// fmt.Println("Example03:")
	// Example03()

	// fmt.Println("Example04:")
	// Example04()

	// fmt.Println("Example05:")
	// Example05()

	// fmt.Println("Example06:")
	// Example06()

	fmt.Println("Example07:")
	Example07()

	// f, _ := os.Open("../testdata/test.json")
	// tokens := json.NewDecoder(f)

	// for {
	// 	token, err := tokens.Token()
	// 	fmt.Println(token)

	// 	if err != nil {
	// 		if err.Error() == "EOF" {
	// 			break
	// 		}
	// 	}
	// }
}
