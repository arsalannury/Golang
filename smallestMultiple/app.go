// BY CHAT GPT

package smallestMultiple

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func Run() {
	result := 1

	for i := 2; i <= 20; i++ {
		result = lcm(result, i)
	}

	fmt.Println(result)
}
