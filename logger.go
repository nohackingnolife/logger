package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Info: ", data)
}

func Error(data interface{}) {
	fmt.Println("Error: ", data)
}
