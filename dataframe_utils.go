package gandalff

type PrettyPrintParams struct {
	minColWidth int
	maxColWidth int
	colWidth    int
	width       int
	nrows       int
	indent      string
}

func NewPrettyPrintParams() PrettyPrintParams {
	return PrettyPrintParams{
		minColWidth: 10,
		maxColWidth: 20,
		colWidth:    15,
		width:       80,
		nrows:       10,
		indent:      "",
	}
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
