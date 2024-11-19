package training

import (
	"fmt"
	"time"
)

type Post struct {
	createdAt time.Time
	author    string
	photos    struct {
		photoName    string
		photoContent []byte
	}
	content string
}

type Status struct {
	lastSee time.Time
	status  string
}

func Training() {
	var firstUserPost Post = Post{
		createdAt: time.Now(),
		author:    "Arsalan Nury",
		content:   "This situation is verification specification dagholination",
		photos: struct {
			photoName    string
			photoContent []byte
		}{
			photoName:    "our last week trip",
			photoContent: []byte("ghgyutyughjghg"),
		},
	}

	var userStatus Status = Status{
		lastSee: time.Date(2023, time.January, 2, 12, 25, 10, 1, time.UTC),
		status:  "online",
	}

	postDetail(&firstUserPost)
	statusDetail(userStatus)
}

func postDetail(model *Post) {
	//fmt.Print(model)
}

func statusDetail(model Status) {
	fmt.Print("This user was available at ", model.lastSee, " and current status is ", model.status)
}
