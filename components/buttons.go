package components

type Button struct {
	base
}

func (c Button) New() Component {
	attrs := map[string]string{
		"class": "",
	}
	return Button{base: base{template: "button", HTMLAttrs: attrs}}
}

func NewBtn() Component {
	return NewComponent("button")
}
