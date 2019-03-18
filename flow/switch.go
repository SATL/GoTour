package main

import (
	"fmt"
	"runtime"
	"time"
)

func whenIsSaturday() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")

	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away")
	}
}

func whatMyOS() {
	fmt.Print("Go runt on ")

	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")

	default:
		fmt.Printf("%s. \n", os)

	}
}

func switchNoCondition() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
}

func main_switch() {
	whatMyOS()
	whenIsSaturday()
	switchNoCondition()
}
