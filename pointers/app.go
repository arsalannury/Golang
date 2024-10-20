package pointers

import "fmt"

/*func zero(x *int) {
	*x = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is still 5
}*/

func Pointer() {
	var initial int16 = 1000
	fmt.Println(&initial)
	logger(&initial)
}

func logger(initial *int16) {
	fmt.Println(initial)
}
