package randomNumberGame

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func Run() (string, error) {
	const userInputMessage = "Hello and welcome to the RNG :)\nYou should guess what number will be print in the terminal between 1 and 10 \n"

	fmt.Print(userInputMessage)
	fmt.Println("--------------------------------")
	fmt.Print("Enter guess number: ")

	reader := bufio.NewReader(os.Stdin)
	terminalData, err := reader.ReadString('\n')
	//terminalData = strings.Replace(terminalData, "\n", "", -1)

	if err != nil {
		return "", errors.New("")
	}

	return terminalData, nil
}
