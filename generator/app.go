package generator

import (
	"fmt"
	"os"
)

func Generator() {
	const choose string = "Select What Type Of File Do You Want Generate ? "
	const createFile string = "1- createFile"
	const serviceFile string = "2- serviceFile"
	const displayFile string = "3- displayFile"

	var userSelectedIndex int8

	for {
		fmt.Println(choose)
		fmt.Println(createFile)
		fmt.Println(serviceFile)
		fmt.Println(displayFile)

		fmt.Scan(&userSelectedIndex)

		switch userSelectedIndex {
		case 1:
			generateCreateFile()
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		default:
			fmt.Println("invalid option selected")
		}

	}

}

func generateCreateFile() {
	var content []byte = []byte(createFileContent)
	err := os.WriteFile("./ts/create.txt", content, 0644)

	if err != nil {
		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - ")
		fmt.Println(err)
		panic("error occurred at WriteFile")
	}

	fmt.Println("file created at ./ts/create.txt")
}
