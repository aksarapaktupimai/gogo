package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("hello world", "1", "x", 2))
	fmt.Println(strings.Replace("hello World", "1", "x", -1))
}
