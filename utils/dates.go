package utils

import (
	"fmt"
	"time"

	"github.com/jalaali/go-jalaali"
)


func DateToGeorgian (year int,month jalaali.Month,day int) {
 fmt.Println(jalaali.ToGregorian(year,month,day))
}

func DateToJalali (year int,month time.Month,day int) {
 fmt.Println(jalaali.ToJalaali(year,month,day))
}

func JalaliMonth () {
	fmt.Println()
}


