package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("-morning-night-", "-"))
	fmt.Println(strings.Trim("+morning night-", "-+"))
}
