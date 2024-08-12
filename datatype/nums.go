package datatype

import "fmt"

var temperture int8 = -10
var earthAge uint32 = 3000000000

func EarthDetail() {
	fmt.Print("The earth age is -")
	fmt.Print(earthAge)
	fmt.Print(" - and its average temperture is ")
	fmt.Print(temperture)
}