package main

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

func main() {
	t := components.NewComponent("div")
	fmt.Println(t.Render())
}
