package main

import (
	//"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, el := range strings.Fields(s) {
		_, ok := count[string(el)]
		if !ok {
			count[el] = 1
		} else {
			count[el]++
		}

	}
	return count
}

func main_exercise() {
	//wc.Test(WordCount)
}
