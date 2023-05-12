package logger

import "fmt"

func Info(data interface{}) {
	fmt.Println("Info: ", data)
}

func Warning(data interface{}) {
	fmt.Println("Warning: ", data)
}

func Error(data interface{}) {
	fmt.Println("Error: ", data)
}

func Fatal(data interface{}) {
	fmt.Println("Fatal: ", data)
}
