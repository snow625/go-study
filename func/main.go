package main

import "fmt"

func main() {
	fmt.Println("Hello")

	result := hello(3)

	println(result)
}

func hello(x int) int {
	fmt.Println("Hello func")
	return x + 5
}
