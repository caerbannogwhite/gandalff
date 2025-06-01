package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilepickerStyles struct {
	DisabledFile lipgloss.Style
	Selected     lipgloss.Style
}

type Filepicker struct {
	AllowedTypes     []string
	CurrentDirectory string
	Styles           FilepickerStyles
}

func (f Filepicker) Init() tea.Cmd {
	return nil
}

func (f Filepicker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	panic("implement me")
}

func (f Filepicker) View() string {
	panic("implement me")
}

func (f Filepicker) DidSelectFile(msg tea.Msg) (bool, string) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
		return true, ""
	} else {
		return false, ""
	}
}

func (f Filepicker) DidSelectDisabledFile(msg tea.Msg) (bool, string) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
		return true, ""
	} else {
		return false, ""
	}
}
