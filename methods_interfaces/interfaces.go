package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func abs() {
	var a Abser

	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	fmt.Println(a.Abs())
	a = &v
	fmt.Println(a.Abs())
	a = &v
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func impl() {
	var i T = T{"Hello"}
	i.M()
}

func values() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func empty_interface() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func assertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*f = i.(float64)
	fmt.Println(f)*/
}

func switches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main_interfaces() {
	abs()
	impl()
	values()
	empty_interface()

	assertions()

	switches(21)
	switches("hello")
	switches(true)
}
