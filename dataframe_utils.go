package gandalff

import "github.com/charmbracelet/lipgloss"

type PrettyPrintParams struct {
	useLipGloss bool
	minColWidth int
	maxColWidth int
	colWidth    int
	width       int
	nrows       int
	indent      string

	styleNames lipgloss.Style
	styleTypes lipgloss.Style
}

func NewPrettyPrintParams() PrettyPrintParams {
	return PrettyPrintParams{
		minColWidth: 10,
		maxColWidth: 20,
		colWidth:    11,
		width:       100,
		nrows:       10,
		indent:      "",

		styleNames: lipgloss.NewStyle().Bold(true),
		styleTypes: lipgloss.NewStyle().Bold(true).Italic(true),
	}
}

func (ppp PrettyPrintParams) SetUseLipGloss(b bool) PrettyPrintParams {
	ppp.useLipGloss = b
	return ppp
}

func (ppp PrettyPrintParams) SetMinColWidth(n int) PrettyPrintParams {
	ppp.minColWidth = n
	return ppp
}

func (ppp PrettyPrintParams) SetMaxColWidth(n int) PrettyPrintParams {
	ppp.maxColWidth = n
	return ppp
}

func (ppp PrettyPrintParams) SetWidth(n int) PrettyPrintParams {
	ppp.width = n
	return ppp
}

func (ppp PrettyPrintParams) SetNRows(n int) PrettyPrintParams {
	ppp.nrows = n
	return ppp
}

func (ppp PrettyPrintParams) SetIndent(s string) PrettyPrintParams {
	ppp.indent = s
	return ppp
}
