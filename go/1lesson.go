package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var userInput string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&userInput)

	reversedString := reverseString(userInput)
	fmt.Printf("Обратная строка: %s\n", reversedString)
}
