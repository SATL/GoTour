package main

import "fmt"

func main_pointer() {
	i, j := 42, 2701

	p := &i         //point of i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
