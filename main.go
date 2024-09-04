package main

import (
	"calculator/date"
	"fmt"
)

func main() {
	const chooseConverter string = "please select your converter function to continue run application"
	var converter string

	fmt.Println(chooseConverter)
	fmt.Println("1- convert georgian date to jalali date")
	fmt.Println("2- convert jalali date to georgian date")

	fmt.Scan(&converter)

	if converter == "1" {
		date.GeorgianToJalali()
	} else if converter == "2" {
		date.JalaliToGeorgian()
	}
}
