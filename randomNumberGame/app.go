package randomNumberGame

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func Run() {
	const userInputMessage = "Hello and welcome to the RNG :)\nYou should guess what number will be print in the terminal between 0 and 50 \n"
	fmt.Print(userInputMessage)
	fmt.Println("--------------------------------")
	fmt.Println("- Enter E for exit app -")
	for {
		fmt.Print("\n")
		fmt.Print("Enter the number: ")

		reader := bufio.NewReader(os.Stdin)
		terminalData, err := reader.ReadString('\n')

		var removeEnters = strings.Trim(terminalData, "\n")
		var removeSpecialChars = strings.Trim(removeEnters, "\r")

		if strings.EqualFold(removeSpecialChars, "e") {
			break
		}

		if err != nil {
			panic(err)
		}

		var randomNumber int = rand.IntN(50)
		var result, _ = strconv.Atoi(removeSpecialChars)

		if randomNumber == result {
			fmt.Printf("Congratulations, You guessed the correct number %v", result)
		} else {
			fmt.Printf("Oh, You guessed the wrong number %v", result)
		}

	}
}
