package main

import "fmt"

type VertexS struct {
	X int
	Y int
}

func main_structs() {
	v1 := VertexS{1, 2}
	var v2 = VertexS{3, 4}
	fmt.Println(VertexS{1, 2})
	fmt.Println(v1)
	v2.X = 4
	p := &v1
	p.X = 1e9
	fmt.Println(v1)
	fmt.Println(v2)

	var v3 = VertexS{X: 12}
	fmt.Println(v3)
}
