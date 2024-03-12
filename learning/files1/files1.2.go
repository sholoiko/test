package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Println("Привет! Введите число:")

	// Читаем ввод пользователя
	input, err := getUserInput()
	if err != nil {
		log.Fatal("Ошибка ввода:", err)
	}

	// Преобразуем введенную строку в число
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal("Ошибка преобразования в число:", err)
	}

	// Умножаем введенное число на 2
	result := multiplyByTwo(number)

	// Выводим результат
	fmt.Printf("Результат умножения на 2: %f\n", result)
}

// getUserInput запрашивает ввод пользователя и возвращает введенную строку
func getUserInput() (string, error) {
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}

// multiplyByTwo умножает число на 2 и возвращает результат
func multiplyByTwo(number float64) float64 {
	return number * 2
}
