package tests

import (
	"fmt"
	"test/card"
)

func ExampleMaskCard() {
	fmt.Println(card.MaskCard("4324563451234567"))
	//Output:
	//4324 **** **** 4567
}

func ExampleMaskCard_panicLength() {
	defer func() {
		fmt.Println(recover())
	}()

	fmt.Println(card.MaskCard("123"))
	//Output:
	//неверная длина карты
}

func ExampleMaskCard_panicSymbols() {
	defer func() {
		fmt.Println(recover())
	}()
	
	fmt.Println(card.MaskCard("8605abcd51234567"))
	//Output:
	//номер карты должен содержать только цифры
}
