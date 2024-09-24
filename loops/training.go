package loops

import "fmt"

func Counter() {
	const counterQuestion string = "Please enter the number you want count from that: "
	var counter int

	fmt.Print(counterQuestion)
	fmt.Scan(&counter)

	for i := 1; i <= counter; i++ {
		fmt.Print(i, " ")
		if i == counter {
			fmt.Println("counter end :)")
			break
		}
	}
}

func Multiple() {
	const multiplicationQuestion = "Please enter number you want multiplication to: "
	const multipleQuestion = "Please enter number you want set as multiple: "

	var multiplicationNumber int // like 1 to 30
	var multiple int             // like 5 or 2 or 6 etc...

	fmt.Print(multipleQuestion)
	fmt.Scan(&multiple)

	fmt.Print(multiplicationQuestion)
	fmt.Scan(&multiplicationNumber)

	if multiple < 0 || multiplicationNumber < 0 {
		fmt.Println(`---------------------------------------------- invalid data entered ----------------------------------------------`)
		return
	}

	for i := 1; i <= multiplicationNumber; i++ {
		fmt.Print(i*multiple, " ")

		if i == multiplicationNumber {
			fmt.Println("operation end")
			break
		}
	}
}

func PrintStar() {
	for i := 1; i <= 9; i++ {
		if i == 1 {
			fmt.Println("*")
		}
		if i == 2 {
			fmt.Println("**")
		}
		if i == 3 {
			fmt.Println("***")
		}
		if i == 4 {
			fmt.Println("****")
		}
		if i == 5 {
			fmt.Println("*****")
		}
		if i == 6 {
			fmt.Println("******")
		}
		if i == 7 {
			fmt.Println("*******")
		}
		if i == 8 {
			fmt.Println("********")
		}
		if i == 9 {
			fmt.Println("*********")
		}
	}

	for i := 8; i <= 1; i-- {
		if i == 8 {
			fmt.Println("********")
		}
		if i == 7 {
			fmt.Println("*******")
		}
		if i == 6 {
			fmt.Println("******")
		}
		if i == 5 {
			fmt.Println("*****")
		}
		if i == 4 {
			fmt.Println("****")
		}
		if i == 3 {
			fmt.Println("***")
		}
		if i == 2 {
			fmt.Println("**")
		}

		if i == 1 {
			fmt.Println("*")
		}

	}
}
