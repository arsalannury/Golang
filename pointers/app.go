package pointers

import "fmt"

func Pointer() {
	var years [3]int16 = [3]int16{1400, 1401, 1402}

	yearsLogger(&years)
}

func yearsLogger(years *[3]int16) {
	for index := 0; index < len(*years); index++ {
		fmt.Println(years[index], index)
	}
}

// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

func Addresses() {
	var text string = "Text String"
	var pointText *string = &text
	fmt.Println(pointText, "pointText")
	address(pointText)
}

func address(pointText *string) {
	fmt.Println(pointText)
}
