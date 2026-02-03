package main

import "fmt"

func main() {
	// i
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// ii
	subslice1 := greetings[:2]
	subslice2 := greetings[1:4]
	subslice3 := greetings[3:]

	// iii
	fmt.Println(subslice1)
	fmt.Println(subslice2)
	fmt.Println(subslice3)
}
