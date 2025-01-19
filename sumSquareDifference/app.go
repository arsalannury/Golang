package sumSquareDifference

import (
	"errors"
	"fmt"
	"math"
)

func Sum(slice []int) (int, error) {
	if len(slice) >= 1 {
		var sum int = 0
		for _, value := range slice {
			sum += value
		}
		return sum, nil
	}
	return 0, errors.New("passed slice must have more than 0 length")
}

func SumSquareDifference() float64 {
	var squareToHundred []int
	var squareOfTheSum []int

	for ind := 1; ind <= 100; ind++ {
		fmt.Println(math.Pow(float64(ind), 2))
		var iteratorPow int = int(math.Pow(float64(ind), 2))
		squareToHundred = append(squareToHundred, iteratorPow)
		squareOfTheSum = append(squareOfTheSum, ind)
	}

	sumOfSquares, _ := Sum(squareToHundred)
	sumOfSquareOfTheSum, _ := Sum(squareOfTheSum)
	powToTwo := math.Pow(float64(sumOfSquareOfTheSum), 2)

	return powToTwo - float64(sumOfSquares)

}
