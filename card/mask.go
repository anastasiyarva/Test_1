package card

func MaskCard(card string) string {
	if len(card) != 16 {
		panic("неверная длина карты")
	}

	for i := 0; i < len(card); i++ {
		if card[i] < '0' || card[i] > '9' {
			panic("номер карты должен содержать только цифры")
		}
	}

	result := ""

	for i := 0; i < len(card); i++ {
		if i == 4 || i == 8 || i == 12 {
			result += " "
		}

		if i < 4 || i >= 12 {
			result += string(card[i])
		} else {
			result += "*"
		}
	}

	return result
}
