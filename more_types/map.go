package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func basic_map() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m)
}

func literal() {
	var m = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}

	fmt.Println(m)
}

func literal_continue() {
	var m = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967}, //vertex ommited
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func mutating_maps() {
	m := make(map[string]int)
	m["Answer"] = 42

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main_map() {
	basic_map()
	literal()
	literal_continue()

	mutating_maps()
}
