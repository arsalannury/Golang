package training

import (
	"errors"
	"strings"
	"time"
)

type province struct {
	provinceName string
	provinceCode uint8
}

type city struct {
	cityName string
	cityCode uint8
}

type user struct {
	firstName     string
	lastName      string
	fatherName    string
	birthDate     time.Time
	birthProvince province
	birthCity     city
}

func newUserCreation(firstName string, lastName string, fatherName string, birthDate time.Time, birthProvince province, birthCity city) (*user, error) {
	if len(strings.TrimSpace(firstName)) == 0 || len(strings.TrimSpace(lastName)) == 0 || len(strings.TrimSpace(fatherName)) == 0 || len(strings.TrimSpace(birthProvince.provinceName)) == 0 || len(strings.TrimSpace(birthCity.cityName)) == 0 {
		return nil, errors.New("validation error occurred")
	}

	if birthDate == time.Now() || birthDate.GoString() > time.Now().GoString() {
		return nil, errors.New("validation error occurred, birth date must be smaller than current date")
	}

	return &user{
		firstName:     firstName,
		lastName:      lastName,
		fatherName:    fatherName,
		birthDate:     birthDate,
		birthProvince: birthProvince,
		birthCity:     birthCity,
	}, nil
}
