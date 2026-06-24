package main

import "fmt"

func main() {
	score := 8
	/*
		score++
		score++
	*/

	/*
		if score >= 10 {
			if score > 15 {
				fmt.Println("You super!")
			} else {
				fmt.Println("You Best!")
			}
		} else {
			fmt.Println("You need to be better...")
		}
	*/

	if score > 15 {
		fmt.Println("You super!")
	} else if score == 10 {
		fmt.Println("Not Bad!")
	} else if score == 10 || score == 9 {
		fmt.Println("Not Bad!")
	} else if score >= 10 {
		fmt.Println("You Best!")
	} else {
		fmt.Println("You need to be better...")
	}

}
