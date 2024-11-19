package smallestMultiple

import "fmt"

func SmallestMultiple() {
	var nums [20]int = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	var counter int = 1
	var index int = 0

	for {
		if counter == 300 {
			break
		}

		if index == 19 {
			counter += 1
			index = 0
		}

		if nums[index]%counter == 0 {
			index += 1
			fmt.Println(nums[index] % counter)
		}
	}

}
