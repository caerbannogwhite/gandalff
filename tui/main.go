package main

import (
	"fmt"
	"os"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	"github.com/charmbracelet/lipgloss"
	"github.com/gocarina/gocsv"

	tea "github.com/charmbracelet/bubbletea"
)

func run() {
	// read in CSV data
	f, err := os.Open("sample.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	type SampleData struct {
		ID         int    `csv:"id"`
		FirstName  string `csv:"First Name"`
		LastName   string `csv:"Last Name"`
		Age        int    `csv:"Age"`
		Occupation string `csv:"Occupation"`
	}
	var sampleData []*SampleData

	if err := gocsv.UnmarshalFile(f, &sampleData); err != nil {
		panic(err)
	}

	headers := []string{"id", "First Name", "Last Name", "Age", "Occupation"}
	ratio := []int{1, 10, 10, 5, 10}
	minSize := []int{4, 5, 5, 2, 5}

	var s string
	var i int
	types := []any{i, s, s, i, s}

	m := model{
		table:   table.NewTable(0, 0, headers),
		infoBox: flexbox.New(0, 0).SetHeight(7),
		headers: headers,
	}
	// set types
	_, err = m.table.SetTypes(types...)
	if err != nil {
		panic(err)
	}
	// setup dimensions
	m.table.SetRatio(ratio).SetMinWidth(minSize)
	// set style passing
	m.table.SetStylePassing(true)
	// add rows
	// with multi type table we have to convert our rows to []any first which is a bit of a pain
	var orderedRows [][]any
	for _, row := range sampleData {
		orderedRows = append(orderedRows, []any{
			row.ID, row.FirstName, row.LastName, row.Age, row.Occupation,
		})
	}
	m.table.MustAddRows(orderedRows)

	// setup info box
	infoText := `
use the arrows to navigate
ctrl+s: sort by current column
alphanumerics: filter column
enter, spacebar: get column value
ctrl+c: quit
`
	r1 := m.infoBox.NewRow()
	r1.AddCells(
		flexbox.NewCell(1, 1).
			SetID("info").
			SetContent(infoText),
		flexbox.NewCell(1, 1).
			SetID("info").
			SetContent(selectedValue).
			SetStyle(lipgloss.NewStyle().Bold(true)),
	)
	m.infoBox.AddRows([]*flexbox.Row{r1})

	p := tea.NewProgram(&m, tea.WithAltScreen())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func main() {

}
