package main

// import "fmt"

func main() {
	number := 10

	var ptr *int // 0x0
	pointer := &number

	if ptr != nil {
		println(*ptr, "Default off")
	}

	foo(pointer)
	println(*pointer)
	println(number)
	println(ptr)
}

func foo(n *int) {
	println(*n, "inside")
	println("Mutate")
	*n = 3
}
