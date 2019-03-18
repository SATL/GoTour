package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	item := make([][]uint8, dy)

	for i := range item {
		item[i] = make([]uint8, dy)
	}

	for i := range item {
		for j := range item[i] {
			item[i][j] = uint8(i * j)
		}
	}
	printSlice(item)

	return item
}

func printSlice(s [][]uint8) {
	fmt.Println(s)
}

func main_exe() {
	pic.Show(Pic)
}
