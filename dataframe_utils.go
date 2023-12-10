package gandalff

import "github.com/charmbracelet/lipgloss"

type PPrintParams struct {
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

func NewPPrintParams() PPrintParams {
	return PPrintParams{
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

func (ppp PPrintParams) SetIndent(s string) PPrintParams {
	ppp.indent = s
	return ppp
}
