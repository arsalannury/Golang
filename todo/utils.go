package todo

import (
	"os"
	"strings"
	"time"
)

func writeFile(jasonString []byte) {
	_, errReadFile := os.ReadFile("./todo/todo.json")
	var osApproach int

	if errReadFile != nil {
		osApproach = os.O_CREATE
	} else {
		osApproach = os.O_APPEND
	}

	file, errWriteFile := os.OpenFile("./todo/todo.json", osApproach, 0644)
	check(errWriteFile)

	file.Write(jasonString)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type todo struct {
	Title     string
	Content   string
	State     string
	CreatedAt time.Time
}

func (t todo) New() todo {

	if len(strings.TrimSpace(t.State)) == 0 || !strings.EqualFold(t.State, "TODO") && !strings.EqualFold(t.State, "DONE") {
		panic("title is not valid: you should use TODO or DONE for valid options")
		return todo{}
	}

	if len(strings.TrimSpace(t.Content)) == 0 {
		panic("content can not be empty or white spaces")
		return todo{}
	}

	return todo{
		Title:     t.Title,
		State:     t.State,
		CreatedAt: time.Now(),
		Content:   t.Content,
	}
}
