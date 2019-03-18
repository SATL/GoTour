package main

import (
	"fmt"
)

func hello() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func stacked() {
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func main() {
	defer hello()

	stacked()
}
