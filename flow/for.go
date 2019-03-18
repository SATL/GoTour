package main

import "fmt"

func main_for() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//skip declaration

	sum2 := 1
	for sum2 < 1090 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}
