package main

import (
	"myApp/mail"
	"time"
)

func main() {
	mail.SendEmail("You should always eat meat to be healthy");
	mail.ReceiveEmail("Arsalan Nury ",time.Now().Local());
}