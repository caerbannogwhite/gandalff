package dataframe

import "github.com/charmbracelet/lipgloss"

type PPrintParams struct {
	index       bool
	useLipGloss bool
	minColWidth int
	maxColWidth int
	colWidth    int
	width       int
	nrows       int
	tailLen     int
	indent      string

	styleNames lipgloss.Style
	styleTypes lipgloss.Style
}

func NewPPrintParams() PPrintParams {
	return PPrintParams{
		index:       true,
		minColWidth: 10,
		maxColWidth: 20,
		colWidth:    11,
		width:       100,
		nrows:       10,
		tailLen:     3,
		indent:      "",

		styleNames: lipgloss.NewStyle().Bold(true),
		styleTypes: lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.Color("241")),
	}
}

func (ppp PPrintParams) SetIndex(b bool) PPrintParams {
	ppp.index = b
	return ppp
}

func (ppp PPrintParams) SetUseLipGloss(b bool) PPrintParams {
	ppp.useLipGloss = b
	return ppp
}

func (ppp PPrintParams) SetMinColWidth(n int) PPrintParams {
	ppp.minColWidth = n
	return ppp
}

func (ppp PPrintParams) SetMaxColWidth(n int) PPrintParams {
	ppp.maxColWidth = n
	return ppp
}

func (ppp PPrintParams) SetWidth(n int) PPrintParams {
	ppp.width = n
	return ppp
}

func (ppp PPrintParams) SetNRows(n int) PPrintParams {
	ppp.nrows = n
	return ppp
}

func (ppp PPrintParams) SetTailLen(n int) PPrintParams {
	ppp.tailLen = n
	return ppp
}

func (ppp PPrintParams) SetIndent(s string) PPrintParams {
	ppp.indent = s
	return ppp
}
