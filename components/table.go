package components

import "strings"

type Table struct {
	base
}

func (c Table) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.Table, " "),
	}

	return Table{base: base{template: "table", HTMLAttrs: attrs}}
}

func NewTable() Component {
	return NewComponent("table")
}

type Tbody struct {
	base
}

func (c Tbody) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.Tbody, " "),
	}

	return Tbody{base: base{template: "tbody", HTMLAttrs: attrs}}
}

func NewTBody() Component {
	return NewComponent("tbody")
}

type Thead struct {
	base
}

func (c Thead) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.Thead, " "),
	}

	return Thead{base: base{template: "thead", HTMLAttrs: attrs}}
}

func NewTHead() Component {
	return NewComponent("thead")
}

type TD struct {
	base
}

func (c TD) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.TD, " "),
	}

	return TD{base: base{template: "td", HTMLAttrs: attrs}}
}

type TH struct {
	base
}

func (c TH) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.TH, " "),
	}

	return TH{base: base{template: "th", HTMLAttrs: attrs}}
}

type TR struct {
	base
}

func (c TR) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.TR, " "),
	}

	return TR{base: base{template: "tr", HTMLAttrs: attrs}}
}
