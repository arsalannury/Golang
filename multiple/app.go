package multiple

func MultiplesThreeAndFive() []int {

	var multiples []int = []int{}

	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples = append(multiples, i)
		}
	}
	return multiples
}
