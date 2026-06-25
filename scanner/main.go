package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Please insert ")
	scanner := bufio.NewScanner(os.Stdin)

	// ok := scanner.Scan()
	// if !ok {
	// 	fmt.Println("Error")
	// 	return
	// }

	if ok := scanner.Scan(); !ok {
		fmt.Println("Error")
		return
	}
	// value := strings.Split(scanner.Text(), " ")
	value := scanner.Text()

	fields := strings.Fields(value)

	if len(fields) == 0 {
		return
	}

	fmt.Println("you have inserted:", fields)
}
