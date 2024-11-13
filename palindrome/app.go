package palindrome

import (
	"math"
)

func Palindrome() int {

	var variable int = 800
	var sortSlice []int
	var revSlice []int
	var palindrome int

	for counter := 800; counter <= 999; counter++ {
		var res = variable * counter
		variable += 1
		sortSlice = append(sortSlice, res)
		revSlice = append(revSlice, reverseInt(res))
	}

	for item := 0; item <= len(sortSlice)-1; item++ {
		if sortSlice[item] == revSlice[item] {
			//fmt.Println(sortSlice[item])
			palindrome = sortSlice[item]
			break
		}
	}

	return palindrome
}

func reverseInt(x int) int {
	x = int(math.Abs(float64(x)))

	var reversedDigit int

	for x > 0 {
		lastDigit := x % 10
		reversedDigit = reversedDigit*10 + lastDigit

		x = x / 10
	}

	return reversedDigit
}
