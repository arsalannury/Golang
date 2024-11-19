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

	postDetail(&firstUserPost)
}

func postDetail(model *Post) {
	fmt.Println((*model).photos)
	fmt.Println((*model).content)
	fmt.Println((*model).author)
	fmt.Println((*model).createdAt)
}
