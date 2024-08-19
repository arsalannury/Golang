package chars

import (
	"fmt"
	"strings"
)

var Str string = "Arsalan Nury"

func RunCode() string {
	//fmt.Println(len(Str))
	//fmt.Printf("%c", Str[1])
	fmt.Println(strings.Trim("Arsalan Nury * * * *", "  *  "))
	fmt.Print("ArsalanNury"[:3])
	fmt.Print("ArsalanNury"[3:7])
	fmt.Print("ArsalanNury"[7:])
	return ""
}
