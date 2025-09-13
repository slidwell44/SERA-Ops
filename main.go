package main

import "fmt"

func main() {
	var message string = Hello("Simon")
	fmt.Println(message)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
