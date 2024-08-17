package datePicker

import (
	"fmt"
	"time"
)

func Now() string {
	return time.Now().GoString()
}

func CalcTimeUntilNow(add time.Duration) time.Duration {
	feature := time.Now().Add(add * time.Second)
	return time.Until(feature)
}

func Sleep(add time.Duration) time.Duration {
	feature := time.Now().Add(add * time.Second)
	untilNow := time.Until(feature)
	fmt.Println("The Task Executed At", untilNow, "From Now")
	//return time.Sleep(untilNow)
	return time.Second
}
