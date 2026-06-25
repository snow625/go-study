package main

import (
	"fmt"
	"slices"
)

func main() {
	// static sArr
	sArr := [5]int{5, 2, 33, 4, 5}
	sArr[4] = 11 // mutate

	for i := 0; i < len(sArr); i++ {
		sArr[i] = sArr[i] * 2
	}

	// for index, el := range sArr {
	// 	fmt.Println(index, el)
	// }

	// fmt.Println(sArr)
	// _______________________________________

	// slice
	slArr := []int{5, 2, 33, 4, 5}

	slices.Reverse(slArr)
	slArr = append(slArr, 1)

	// fmt.Println(slArr)
	// fmt.Println(slArr[len(slArr)-1])

	// ___________________________________
	// intSlice := make([]int, 0)
	intSlice := make([]int, 0, 5)
	// intSlice = append(intSlice, 1, 3, 4, 5)
	intSlice = append(intSlice, 2)
	// fmt.Println(intSlice)
	// fmt.Println(len(intSlice))
	// fmt.Println(cap(intSlice))
	// _____________________________________

	// MAP

	// weather := map[int]int{
	// 	1: 3,
	// 	2: 4,
	// 	3: 5,
	// }

	weather := make(map[int]int, 10)

	// s, isOk := weather[4]
	// fmt.Println(s, isOk)

	for key, value := range weather {
		fmt.Println(value)
		if value == 4 {
			continue
		}
		weather[key] = value * 2
	}
	fmt.Println(weather)

}
