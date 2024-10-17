package components

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
)

type base struct {
	template string
	Text     string
	Children
	HTMLAttrs
}

func (c base) AddClass(class string) Component {
	c.HTMLAttrs["class"] += class
	return c
}

func (c base) AddChild(child Component) Component {
	c.Children = append(c.Children, child)
	return c
}

func (c base) New() Component {
	attrs := map[string]string{
		"class": "",
	}
	return base{HTMLAttrs: attrs}
}

func (c base) Execute(wr io.Writer) {
	fmt.Println(os.Getwd())
	filename := fmt.Sprintf("templates/%s.html", c.template)
	tmpl, err := template.ParseFiles(filename)

	if err != nil {
		log.Fatal("Failed to load template: ", err)
	}

	err = tmpl.Execute(wr, templateData{
		Attrs:    c.HTMLAttrs.convert(),
		Children: c.Children.convert(),
		Text:     c.Text,
	})

	if err != nil {
		log.Fatal("Failed to execute template: ", err)
	}
}

func (c base) Render() template.HTML {
	buf := new(bytes.Buffer)
	c.Execute(buf)

	return template.HTML(buf.String())
}

func (c base) SetAttr(name, value string) Component {
	c.HTMLAttrs[name] = value
	return c
}

func (c base) SetText(value string) Component {
	c.Text = value
	return c
}

func (c base) OverrideClasses(classes string) Component {
	c.HTMLAttrs["class"] = classes
	return c
}

type templateData struct {
	Attrs    []template.HTMLAttr
	Children []template.HTML
	Text     string
}

type Children []Component

func (c Children) convert() []template.HTML {
	var result []template.HTML
	for _, child := range c {
		result = append(result, child.Render())
	}

	return result
}

type HTMLAttrs map[string]string

func (a HTMLAttrs) convert() []template.HTMLAttr {
	var result []template.HTMLAttr

	for k, v := range a {
		if v != "" {
			result = append(result, template.HTMLAttr(fmt.Sprintf("%s='%s'", k, v)))
		}
	}

	return result
}

type Component interface {
	AddClass(string) Component
	AddChild(Component) Component
	Execute(io.Writer)
	New() Component
	Render() template.HTML
	SetAttr(string, string) Component
	SetText(string) Component
	OverrideClasses(string) Component
}

func NewComponent(component string) Component {
	componentList := map[string]Component{
		"button": Button{},
		"div":    Div{},
	}

	c, ok := componentList[component]

	if !ok {
		log.Fatalf("'%s' is not a valid component\n", component)
	}

	return c.New()
}
