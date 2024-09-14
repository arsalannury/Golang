package installments

import "fmt"

func Installments() {
	const enterPrice = " Please enter price "
	const enterAdvancePaymentPercent = " Please enter advance payment percent "
	const resultOfPayment = " Your advance payment result is "
	var price int
	var advancePaymentPercent int
	var advancePayment int

	fmt.Print(enterPrice)
	fmt.Scan(&price)

	fmt.Print(enterAdvancePaymentPercent)
	fmt.Scan(&advancePaymentPercent)

	advancePayment = advancePaymentCalculator(advancePaymentPercent, price)

	fmt.Print(resultOfPayment)
	fmt.Print(advancePayment)
}

func advancePaymentCalculator(advancePaymentPercent int, price int) int {
	return (advancePaymentPercent * price) / 10
}
