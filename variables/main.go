package main

import "fmt"

func main() {
	welcomeText := "Get Ready"
	endText := "Game Over"
	score := 10

	fmt.Println(welcomeText)
	score = score + 1
	score += 1
	score++
	score--

	fmt.Println(endText, "your score:", score)

	floatNumber := 0.5
	fmt.Println(floatNumber)

	isSub := true
	fmt.Println(isSub)

	scoreFloat := 10.0

	fmt.Println(score / 3)      // 3
	fmt.Println(scoreFloat / 3) // 3.333333

	intPart := score / 3
	partLeft := score % 3

	fmt.Println("intPart:", intPart)
	fmt.Println("partLeft:", partLeft)

	text := "Hello"
	andText := text + ", World!"
	fmt.Println(andText)

	// Full type variable
	var isZen bool = true // default false
	var cart int          // default 0
	var float float64     // default 0
	var text2 string      // default ''

	fmt.Println(isZen)
	fmt.Println(cart)
	fmt.Println(float)
	fmt.Println(text2)
}
