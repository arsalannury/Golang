package training

import "fmt"

type Person struct {
	name      string
	age       uint8
	weight    uint16
	tall      uint8
	eyeColor  string
	bloodType struct {
		title string
	}
}

type Employee struct {
	title string
}

func Training() {
	/*var mapping map[string]uint8 = map[string]uint8{
		"APP_ENV_KEY_COUNT": 100,
	}
	fmt.Println(mapping)*/

	/*	var arsalan Person = Person{
			name:      "Arsalan",
			age:       25,
			weight:    100,
			tall:      172,
			eyeColor:  "Brown",
			bloodType: struct{ title string }{title: "A+"},
		}
	*/

	/*var anon = Anonymous()
	for index := 0; index <= 10; index++ {
		fmt.Print(anon(uint8(index)))

		if index < 10 {
			fmt.Print(",")
		}
	}*/

	/*var addOperation = add()

	for index := 0; index <= 30; index++ {
		fmt.Println(addOperation(uint(index), uint(index*2)))
	}*/

	var emp Employee
	emp.title = "Arsalan"
	emp.UpdateName("Ali")
	emp.PrintName()
}

func Anonymous() func(num uint8) uint8 {
	return func(num uint8) uint8 {
		return num * 22

	}
}

func add() func(input, output uint) (string, uint) {
	const resultHelper = "result of the operation is ->"

	return func(input, output uint) (string, uint) {
		return resultHelper, input + output
	}
}

func (e *Employee) UpdateName(newName string) {
	e.title = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.title)
}
