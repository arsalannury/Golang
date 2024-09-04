package main

import (
	"calculator/date"
	"calculator/temp"
	"fmt"
)

func main() {
	const chooseConverter string = "please select your converter function to continue run application"
	var converter string

	fmt.Println(chooseConverter)
	fmt.Println("date converters -------------------------------")
	fmt.Println("1- convert georgian date to jalali date")
	fmt.Println("2- convert jalali date to georgian date")
	fmt.Println("temp converters -------------------------------")
	fmt.Println("3- convert Fahrenheit to Celsius")

	fmt.Scan(&converter)

	if converter == "1" {
		date.GeorgianToJalali()
	} else if converter == "2" {
		date.JalaliToGeorgian()
	} else if converter == "3" {
		temp.FahrenheitCelsius()
	}
}
