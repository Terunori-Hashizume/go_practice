package main

import "fmt"

func main() {
	message := "Hi 👩 and 👨"
	fmt.Println("string: ", message)
	fmt.Println("[]rune: ", []rune(message))
	fmt.Println("[]rune[3] as a char: ", string([]rune(message)[3]))
}
