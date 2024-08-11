package logger

import "fmt"

type Types struct {
    ERROR string
    WARN string
	LOG string
}

func Logger(types Types,content string) {
	if types.ERROR == "ERROR" {
		fmt.Println("Error Log Enabled - " + " " + content);
	}else if types.WARN == "WARN" {
		fmt.Println("Warn Log Enabled - " + " " + content);
	}else if types.LOG == "LOG" {
		fmt.Println("Log Enabled - " + " " + content);
	}
}