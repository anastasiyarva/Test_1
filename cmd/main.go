package main

import (
	"fmt"
	"test/card"
)

func main() {
	var input string

	fmt.Print("Введите ваш номер карты: ")
	fmt.Scan(&input)

	fmt.Println("Карта:", card.MaskCard(input))
}
