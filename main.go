package main

import (
	"fmt"

	"github.com/nadama95/gotemplates/components"
)

func main() {
	c := components.NewComponent("div")

	fmt.Println(c.Render())
}
