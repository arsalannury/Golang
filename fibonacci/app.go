package fibonacci

func Fibonacci() [30]int {
	var finalItems = [30]int{1, 2}

	for i := 2; i <= len(finalItems); i++ {
		finalItems[i] = finalItems[i-2] + finalItems[i-1]
		if finalItems[len(finalItems)-1] != 0 {
			break
		}
	}
	return finalItems
}
