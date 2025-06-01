package main

import (
	"fmt"
	"os"
	"os/exec"
	"unicode"

	"github.com/76creates/stickers/flexbox"
	"github.com/76creates/stickers/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}

// type model struct {
// 	filepicker   Filepicker
// 	selectedFile string
// 	quitting     bool
// 	err          error
// }

// type clearErrorMsg struct{}

// func clearErrorAfter(t time.Duration) tea.Cmd {
// 	return tea.Tick(t, func(_ time.Time) tea.Msg {
// 		return clearErrorMsg{}
// 	})
// }

// func (m model) Init() tea.Cmd {
// 	return m.filepicker.Init()
// }

// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "ctrl+c", "q":
// 			m.quitting = true
// 			return m, tea.Quit
// 		}
// 	case clearErrorMsg:
// 		m.err = nil
// 	}

// 	var cmd tea.Cmd
// 	m.filepicker, cmd = m.filepicker.Update(msg)

// 	// Did the user select a file?
// 	if didSelect, path := m.filepicker.DidSelectFile(msg); didSelect {
// 		// Get the path of the selected file.
// 		m.selectedFile = path
// 	}

// 	// Did the user select a disabled file?
// 	// This is only necessary to display an error to the user.
// 	if didSelect, path := m.filepicker.DidSelectDisabledFile(msg); didSelect {
// 		// Let's clear the selectedFile and display an error.
// 		m.err = errors.New(path + " is not valid.")
// 		m.selectedFile = ""
// 		return m, tea.Batch(cmd, clearErrorAfter(2*time.Second))
// 	}

// 	return m, cmd
// }

// func (m model) View() string {
// 	if m.quitting {
// 		return ""
// 	}
// 	var s strings.Builder
// 	s.WriteString("\n  ")
// 	if m.err != nil {
// 		s.WriteString(m.filepicker.Styles.DisabledFile.Render(m.err.Error()))
// 	} else if m.selectedFile == "" {
// 		s.WriteString("Pick a file:")
// 	} else {
// 		s.WriteString("Selected file: " + m.filepicker.Styles.Selected.Render(m.selectedFile))
// 	}
// 	s.WriteString("\n\n" + m.filepicker.View() + "\n")
// 	return s.String()
// }

// func main() {
// 	fp := Filepicker{AllowedTypes: []string{".csv", ".tsv", ".xpt", ".xlsx"}}
// 	fp.CurrentDirectory, _ = os.UserHomeDir()

// 	m := model{
// 		filepicker: fp,
// 	}
// 	tm, _ := tea.NewProgram(&m, tea.WithOutput(os.Stderr)).Run()
// 	mm := tm.(model)
// 	fmt.Println("\n  You selected: " + m.filepicker.Styles.Selected.Render(mm.selectedFile) + "\n")
// }

var selectedValue string = "\nselect something with spacebar or enter"

type model struct {
	table   *table.Table
	infoBox *flexbox.FlexBox
	headers []string
}

func (m *model) Init() tea.Cmd { return nil }

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height - m.infoBox.GetHeight())
		m.infoBox.SetWidth(msg.Width)
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "down":
			m.table.CursorDown()
		case "up":
			m.table.CursorUp()
		case "left":
			m.table.CursorLeft()
		case "right":
			m.table.CursorRight()
		case "ctrl+s":
			x, _ := m.table.GetCursorLocation()
			m.table.OrderByColumn(x)
		case "enter", " ":
			selectedValue = m.table.GetCursorValue()
			m.infoBox.GetRow(0).GetCell(1).SetContent("\nselected cell: " + selectedValue)
		case "backspace":
			m.filterWithStr(msg.String())
		default:
			if len(msg.String()) == 1 {
				r := msg.Runes[0]
				if unicode.IsLetter(r) || unicode.IsDigit(r) {
					m.filterWithStr(msg.String())
				}
			}
		}

	}
	return m, nil
}

func (m *model) filterWithStr(key string) {
	i, s := m.table.GetFilter()
	x, _ := m.table.GetCursorLocation()
	if x != i && key != "backspace" {
		m.table.SetFilter(x, key)
		return
	}
	if key == "backspace" {
		if len(s) == 1 {
			m.table.UnsetFilter()
			return
		} else if len(s) > 1 {
			s = s[0 : len(s)-1]
		} else {
			return
		}
	} else {
		s = s + key
	}
	m.table.SetFilter(i, s)
}

func (m *model) View() string {
	return lipgloss.JoinVertical(lipgloss.Left, m.table.Render(), m.infoBox.Render())
}
