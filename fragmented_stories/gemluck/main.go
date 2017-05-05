package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	fmt.Println()

	gemstash := random(1, 20)

	switch {
	case gemstash < 10:
		fmt.Println("Looking for gems in your clothing.")
		fmt.Printf("You find %d", gemstash)
		fmt.Println(".")
	case gemstash >= 10:
		fmt.Println("In your pocket, you find", gemstash, "gems.")
		fmt.Println("Two gems to rub together. Better than yesterday.")
	}
	fmt.Println()
}
