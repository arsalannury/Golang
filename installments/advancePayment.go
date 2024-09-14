package installments

import "fmt"

func Installments() {
	const enterPrice = " Please enter price "
	const enterAdvancePaymentPercent = " Please enter advance payment percent "
	const resultOfPayment = " Your advance payment result is "
	const enterAnnualProfit = "Please enter your annual profit (between 20-23)"
	const enterInstallmentsMonth = "Please enter your installments month (between 1-36)"
	const resultProfit = "Your profit result is"

	var price int
	var advancePaymentPercent int
	var advancePayment int
	/*var annualProfit int
	var installmentsMonth int*/

	fmt.Print(enterPrice)
	fmt.Scan(&price)

	fmt.Print(enterAdvancePaymentPercent)
	fmt.Scan(&advancePaymentPercent)

	advancePayment = advancePaymentCalculator(advancePaymentPercent, price)

	fmt.Print(resultOfPayment)
	fmt.Print(advancePayment)
	/*
		fmt.Print(enterAnnualProfit)
		fmt.Scan(&annualProfit)

		fmt.Print(enterInstallmentsMonth)
		fmt.Scan(&installmentsMonth)*/
}

func advancePaymentCalculator(advancePaymentPercent int, price int) int {
	return (advancePaymentPercent * price) / 10
}
