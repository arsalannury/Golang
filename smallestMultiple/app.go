package smallestMultiple

import "fmt"

func Run() {
	nums := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	i := 0
	result := []int{}

	for true {
		i += 1
		for _, value := range nums {
			var divided = i % value

			if divided == 0 {
				result = append(result, divided)
			}

			if value == len(nums) && len(result) < len(nums) {
				result = []int{}
			}

		}
		if len(result) == len(nums) {
			fmt.Println(i)
			return
		}
	}

}
