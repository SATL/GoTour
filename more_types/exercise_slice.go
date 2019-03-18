package main

import (
	//"golang.org/x/tour/pic"
	"fmt"
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
	printSliceE(item)

	return item
}

func printSliceE(s [][]uint8) {
	fmt.Println(s)
}

func main_exec() {
	//pic.Show(Pic)
}
