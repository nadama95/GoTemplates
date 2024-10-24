package components

type Div struct {
	base
}

func (c Div) New() Component {
	attrs := map[string]string{
		"class": "",
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
		"class": "",
	}
	return P{base: base{template: "p", HTMLAttrs: attrs}}
}

func NewP() Component {
	return NewComponent("p")
}
