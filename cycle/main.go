package main

import (
	"fmt"
	// "math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		// randomValue := rand.Intn(6)
		// time.Sleep(2 * time.Second)
		time.Sleep(200 * time.Millisecond)
		fmt.Println("now work with ", i, "...")
	}

}
