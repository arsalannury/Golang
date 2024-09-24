package main

import (
	"example.com/essential/loops"
	"fmt"
)

func main() {
	const chooseProgram = "enter the program that you want run: "

	var program string

	fmt.Print(chooseProgram)
	fmt.Scan(&program)

	if program == "counter" {
		loops.Counter()
	} else if program == "multiple" {
		loops.Multiple()
	} else {
		fmt.Println("the program selected is invalid")
	}

}
