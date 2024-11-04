package components

type LI struct {
	base
}

func (c LI) New() Component {
	attrs := map[string]string{
		"class": "",
	}
	return LI{base: base{template: "li", HTMLAttrs: attrs}}
}

func NewLI() Component {
	return NewComponent("li")
}

type UL struct {
	base
}

func (c UL) New() Component {
	attrs := map[string]string{
		"class": "",
	}
	return UL{base: base{template: "ul", HTMLAttrs: attrs}}
}

func NewUL() Component {
	return NewComponent("ul")
}
