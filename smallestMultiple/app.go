package smallestMultiple

import "fmt"

func SmallestMultiple() {
	nums := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	i := 2518
	result := []int{}

	for true {
		i = i + 2
		for _, value := range nums {
			var divided = i % value
			fmt.Println(value, i, len(result), "value, i,len(result)")
			if divided == 0 {
				result = append(result, divided)
			}

			if value == len(nums) && len(result) < len(nums) {
				result = []int{}
			}
		}
		//fmt.Println(len(result), len(nums), "result-nums")
		if len(result) == len(nums) {
			fmt.Println(i)
			return
		}
	}

}
