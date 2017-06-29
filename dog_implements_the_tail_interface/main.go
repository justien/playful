package main

import (
	"fmt"
	"strings"
)

type dog struct {
	enthusiasm int
	name       string
}

type tail interface {
	wag() float64
}

// DOG IMPLEMENTS THE TAIL INTERFACE
func (d dog) wag() int {
	return d.enthusiasm
}

func happy(d dog) {
	a := make([]string, d.enthusiasm)
	a[0] = "     "

	for i := 0; i <= d.enthusiasm; i++ {
		a = append(a, " ")
		fmt.Print(strings.Trim(fmt.Sprint(a), "[]"))
		fmt.Println("Tail is wagging!")
	}
}

func main() {
	thisdog := dog{4, "Bluey"}
	fmt.Println(thisdog.name, "- how keen are you today???")
	happy(thisdog)
	fmt.Println("You are truly the best dog,", thisdog.name, "!!")
}
