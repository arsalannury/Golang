package fibonacci

import "fmt"

func Fibonacci() int {
	var current int = 0
	var finalItems = []int{1, 2}

	for i := 0; i <= len(finalItems); i++ {
		fmt.Println(finalItems[1], "fdfsdfsd")
		finalItems[len(finalItems)+i] = finalItems[len(finalItems)-1] + finalItems[len(finalItems)-2]
		if len(finalItems) == 3000 {
			break

		}
	}
	fmt.Println(finalItems)

	return current
}
