package training

import (
	"fmt"
	"time"
)

func Training() {
	var adminUser, err = newUserCreation(
		"Amir",
		"Khiabani",
		"Ahmad",
		time.Date(1890, 1, 1, 1, 1, 1, 1, time.Local),
		province{provinceName: "Ghazvin", provinceCode: 12},
		city{cityName: "Alamut", cityCode: 1},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(adminUser)

}
