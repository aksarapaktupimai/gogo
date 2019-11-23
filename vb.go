package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("morning")
	b.WriteString(" ")
	b.WriteString("night")
	fmt.Println(b.String())
}
