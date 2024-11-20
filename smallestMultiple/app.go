package smallestMultiple

import (
	"strings"
)

func SmallestMultiple() {
	var nums [20]int = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var counter int = 1
	var index int = 0
	var passed []bool

	for {
		if len(passed) == len(nums) {
			break
		}

		if index == len(nums) {
			counter += 1
			index = 0
			passed = nil
		}

		if !strings.Contains(string(rune(float64(counter/nums[index]))), ".") {
			passed = append(passed, true)
		}

		index += 1
	}

}
