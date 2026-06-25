package main

import "fmt"

func main() {
	fmt.Println("Hello1")
	hello()
	defer func() {
		println("Defer func")
	}()
	fmt.Println("Hello2")

}

func hello() {
	fmt.Println("start Hello func")
	defer func() {
		println("Defer Hello")
	}()
	fmt.Println("End Hello func")
}
