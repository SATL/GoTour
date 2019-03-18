package main

import "fmt"

func main_while() {
	sum := 1
	for sum < 1000 {
		sum += sum

	}
	fmt.Println(sum)
}
