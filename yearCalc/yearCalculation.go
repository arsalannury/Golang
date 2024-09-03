package yearCalc

import (
	"fmt"
	"time"
)

func YearCalculation() {
	const askBornYear string = "enter your year's birthday"
	var currentYear int = time.Now().Year()
	var calcYear int

	fmt.Println(askBornYear)
	fmt.Scan(&calcYear)
	fmt.Println("you are ", currentYear-calcYear, " years old")
}
