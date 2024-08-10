package mail

import "fmt"

func SendEmailToSpam(emailAddress string) {
	if emailAddress == "spammer" {
		fmt.Println("Email is spam and will be delete");
	}else {
		fmt.Println("Email is not spam");
	}
}