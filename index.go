package main

import (
	"core/fibonacci"
	"core/multiple"
	"fmt"
)

func main() {
	var programSelected int

	fmt.Println("Please select an algorithm")
	fmt.Println("------------------------------------")
	fmt.Println("1: fibonacci")
	fmt.Println("2: multiple of three and five")
	fmt.Println("------------------------------------")
	fmt.Print("run: ")
	fmt.Scan(&programSelected)

	if programSelected == 1 {
		fmt.Print(fibonacci.Fibonacci())
	} else if programSelected == 2 {
		fmt.Print(multiple.MultiplesThreeAndFive())
	} else {
		fmt.Println("invalid Algorithm")
	}
}
