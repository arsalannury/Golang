package fibonacci

import "fmt"

func Fibonacci() int {
	var current int = 0
	var finalItems = [30]int{1, 2}

	fmt.Println(len(finalItems))
	for i := 2; i <= len(finalItems); i++ {
		fmt.Println(finalItems[i-2]+finalItems[i-1], "i1")

		finalItems[i] = finalItems[i-2] + finalItems[i-1]
		if finalItems[len(finalItems)-1] != 0 {
			break
		}
	}
	return current
}
