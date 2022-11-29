package greetings

import "fmt"

// Hello provided the hello function
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
