package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkRec(t, ch)
	close(ch)
}

func WalkRec(t *tree.Tree, ch chan int) {
	if t != nil {
		WalkRec(t.Left, ch)
		ch <- t.Value
		WalkRec(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		a, oka := <-ch1
		b, okb := <-ch2

		if oka && okb {
			if a != b {
				return false
			}
		} else if oka != okb {
			return false
		}

		if !oka && !okb {
			return true
		}
	}
}

func main_exec_binary_tree() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for r := range ch {
		fmt.Println(r)
	}
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
