package todo

import (
	"encoding/json"
	"fmt"
	"strings"
)

func TodoApplication() {
	var title, content, state string

	fmt.Println("Welcome to Todo Application")
	fmt.Println("Please fill out related fields and we will save your todo in Todo.txt file.")

	fmt.Print("Enter title:")
	fmt.Scanln(&title)

	fmt.Print("Enter content:")
	fmt.Scanln(&content)

	fmt.Print("Enter state:")
	fmt.Scanln(&state)

	var newTodo = todo.New(todo{
		Title:   strings.ToUpper(title),
		Content: strings.ToUpper(content),
		State:   state,
	})

	todoJasonByte, err := json.Marshal(newTodo)

	check(err)

	writeFile(todoJasonByte)

	fmt.Println("Todo added successfully and saved in Todo.txt file")

}
