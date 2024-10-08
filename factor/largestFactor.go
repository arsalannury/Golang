package factor

func LargestFactor() int {
	var largeInt int = 600851475143
	var final int

	for {
		final = largeInt / 2
		largeInt = final
		if largeInt == 2 {
			break
		}
	}
	return largeInt
}
