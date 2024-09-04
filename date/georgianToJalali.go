package date

import (
	"fmt"
	"github.com/jalaali/go-jalaali"
	"time"
)

func GeorgianToJalali() {
	const askGeorgianYear string = "enter the georgian year"
	const askGeorgianMonth string = "enter the georgian month"
	const askGeorgianDay string = "enter the georgian day"
	var enteredYear, enteredDay int
	var enteredMonth time.Month

	fmt.Println(askGeorgianYear)
	fmt.Scan(&enteredYear)

	fmt.Println(askGeorgianMonth)
	fmt.Scan(&enteredMonth)

	fmt.Println(askGeorgianDay)
	fmt.Scan(&enteredDay)

	fmt.Println(jalaali.ToJalaali(enteredYear, enteredMonth, enteredDay))
	fmt.Scan()
}

func JalaliToGeorgian() {
	const askJalaliYear string = "enter the jalali year"
	const askJalaliMonth string = "enter the jalali month"
	const askJalaliDay string = "enter the jalali day"
	var enteredYear, enteredDay int
	var enteredMonth jalaali.Month

	fmt.Println(askJalaliYear)
	fmt.Scan(&enteredYear)

	fmt.Println(askJalaliMonth)
	fmt.Scan(&enteredMonth)

	fmt.Println(askJalaliDay)
	fmt.Scan(&enteredDay)

	fmt.Println(jalaali.ToGregorian(enteredYear, enteredMonth, enteredDay))
	fmt.Scan()
}
