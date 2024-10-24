package components

type H2 struct {
	base
}

func (c H2) New() Component {
	attrs := map[string]string{
		"class": "",
	}
	return H2{base: base{template: "h2", HTMLAttrs: attrs}}
}

func NewH2() Component {
	return NewComponent("h2")
}