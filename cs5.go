package main

import "fmt"

const greeting = "Hello World"

func main() {
	fmt.Println(greeting)
	greeting = "Hello Walt"
	fmt.Println(greeting)
}
