package temp

import "fmt"

func FahrenheitCelsius() {
	const inputMessage string = "Please enter Fahrenheit temperature"
	const formula float64 = 1.8 + 32
	var temperature float64

	fmt.Println(inputMessage)
	fmt.Scan(&temperature)
	fmt.Println("result -> ", temperature*formula)

}
