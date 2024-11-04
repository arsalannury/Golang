package factor

// SOLVED BY CHAT_GPT NOT ME

func LargestFactor() int64 {
	var number int64 = 600851475143
	//fmt.Println("Largest prime factor:", largestPrimeFactor(number))
	largestPrimeFactor(number)
	return number
}

func largestPrimeFactor(n int64) int64 {
	var factor int64 = 2
	for factor*factor <= n {
		if n%factor == 0 {
			n /= factor
		} else {
			factor++
		}
	}
	return n
}
