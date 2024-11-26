package training

import (
	"errors"
	"fmt"
	"time"
)

type province struct {
	name     string
	location [2]uint32
}

type users struct {
	fullName      string
	fatherName    string
	birthProvince province
	birthDate     time.Time
	nationalCode  string
}

func (u users) getUserByProperty(property string) (error, string, users) {
	var propertyToLog string
	var all users

	switch property {
	case "fullName":
		propertyToLog = "fullName"
	case "fatherName":
		propertyToLog = "fatherName"
	case "birthProvince":
		propertyToLog = "birthProvince"
	case "birthDate":
		propertyToLog = "birthDate"
	case "nationalCode":
		propertyToLog = "nationalCode"
	case "all":
		all = u
	default:
		propertyToLog = ""
		all = users{}
		return errors.New("property passed is not valid"), propertyToLog, all
	}
	return nil, propertyToLog, all
}

func userActions() {
	var admin users = users{
		fullName:   "Arsalan Nury",
		fatherName: "Mohammad",
		birthProvince: province{
			name:     "Tehran",
			location: [2]uint32{435543545, 3232333232},
		},
		birthDate:    time.Date(1999, 9, 6, 12, 30, 0, 0, time.Local),
		nationalCode: "0022143503",
	}

	var err, property, all = users.getUserByProperty(admin, "f")

	if err != nil {
		panic(err)
	}

	fmt.Println(property, "property")
	fmt.Println(all, "all")
}

func Training() {
	userActions()
}
