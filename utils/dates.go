package utils

import (
	"fmt"
	"time"

	"github.com/jalaali/go-jalaali"
)

func DateToGeorgian(year int, month jalaali.Month, day int) {
	fmt.Println(jalaali.ToGregorian(year, month, day))
}

func DateToJalali(year int, month time.Month, day int) {
	fmt.Println(jalaali.ToJalaali(year, month, day))
}

func JalaliMonth() {
	var months = [12]jalaali.Month{
		jalaali.Farvardin,
		jalaali.Ordibehesht,
		jalaali.Khordad,
		jalaali.Tir,
		jalaali.Mordad,
		jalaali.Shahrivar,
		jalaali.Mehr,
		jalaali.Aban,
		jalaali.Azar,
		jalaali.Dey,
		jalaali.Bahman,
		jalaali.Esfand}

	fmt.Println(months)
}
