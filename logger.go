package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Info: ", data)
}
