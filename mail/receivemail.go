package mail;

import (
	"fmt"
	"time"
)

func ReceiveEmail(sender string, sendDate time.Time) {
	fmt.Println(sender + "send to you an email at" , sendDate)
}