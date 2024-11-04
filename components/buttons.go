package components

import "strings"

type Button struct {
	base
}

func (c Button) New() Component {
	attrs := map[string]string{
		"class": strings.Join(theme.Button, " "),
	}
	return Button{base: base{template: "button", HTMLAttrs: attrs}}
}

func NewBtn() Component {
	return NewComponent("button")
}
