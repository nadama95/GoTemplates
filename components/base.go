package components

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"strings"
)

//go:embed templates/*.html
var templateFS embed.FS

type ClassListType []string

var ClassList ClassListType

type base struct {
	template string
	Text     string
	Children
	HTMLAttrs
}

func (c base) AddClass(class string) Component {
	classes := strings.Split(class, " ")

	for _, c := range classes {
		ClassList = append(ClassList, c)
	}

	c.HTMLAttrs["class"] += fmt.Sprintf(" %s", class)
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

	tmpls, err := template.ParseFS(templateFS, "templates/*.html")

	if err != nil {
		log.Fatalf("Failed to load template: %s", err)
	}

	err = tmpls.ExecuteTemplate(wr, fmt.Sprintf("%s.html", c.template), templateData{
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
		"a":      A{},
		"button": Button{},
		"div":    Div{},
		"h2":     H2{},
		"li":     LI{},
		"p":      P{},
		"table":  Table{},
		"tbody":  Tbody{},
		"td":     TD{},
		"thead":  Thead{},
		"tr":     TR{},
		"ul":     UL{},
	}

	c, ok := componentList[component]

	if !ok {
		log.Fatalf("'%s' is not a valid component\n", component)
	}

	return c.New()
}
