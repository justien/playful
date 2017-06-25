package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// func random(min, max int) {
// 	rand.Seed(time.Now().Unix())
// 	return rand.Intn(max-min) + min
// }
//
// func sombre_wrapper(n int) func() int {
// 	x := 0
// 	return func() int {
// 		x++
// 		gemsum := x + n
// 		return gemsum
// 	}
// }

func main() {
	fmt.Println()

	fmt.Println("It is the beach at night by a cold ocean.")
	fmt.Println("The sand is mixed with a glittering gray and black rutile.")
	fmt.Println("Bright stars glitter against dark sky.")
	fmt.Println("Moonlight, softened by a low diaphanous fog, falls into shadowy foam-crested waves.")

	fmt.Println()

	fmt.Println("The sounds are:")
	fmt.Println("\t water sliding back over grit to the sea")
	fmt.Println("\t barely cresting waves collapsing upon themselves")
	fmt.Println("\t feet pushing and displacing sand")
	fmt.Println("\t wind a soft force around body and head")

	fmt.Println()

}

func wrapper(n int) func() int {
	x := 0
	return func() int {
		x++
		gemsum := x + n
		return gemsum
	}
}
