package components

import "strings"

type Div struct {
	base
}

func (c Div) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.Div, " "),
	}
	return Div{base: base{template: "div", HTMLAttrs: attrs}}
}

func NewDiv() Component {
	return NewComponent("div")
}

type P struct {
	base
}

func (c P) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.P, " "),
	}
	return P{base: base{template: "p", HTMLAttrs: attrs}}
}

func NewP() Component {
	return NewComponent("p")
}
