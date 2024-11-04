package components

import "strings"

type A struct {
	base
}

func (c A) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.A, " "),
	}
	return A{base: base{template: "a", HTMLAttrs: attrs}}
}

func NewA() Component {
	return NewComponent("a")
}
