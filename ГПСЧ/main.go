package main

import (
	"fmt"
)

func CodedByFibonachi(a, b int, message string, k []float32) string {
	runes := []rune(message)
	size := len(runes)
	for i := 0; i < size; i++ {
		if i >= len(k) {
			k = append(k, helper(a, b, i, k))
		}
		key := int32(k[i] * 100)
		runes[i] = runes[i] ^ key
	}
	return string(runes)
}

func helper(a, b, i int, k []float32) float32 {
	temp := k[i-a] - k[i-b]
	if temp < 0 {
		temp += 1.
	}
	return temp
}

func main() {
	a, b := 17, 5
	k := []float32{0.324, 0.125, 0.153, 0.545, 0.541, 0.879, 0.147, 0.658, 0.354, 0.912, 0.456, 0.694, 0.954, 0.357, 0.014, 0.751, 0.469}
	message := "Hello, World!shfgdcfjahdgbhlskvjhfioshgliwysgbedhjblwjed"
	message = "Задача организации, в особенности же внедрение современных методик представляет собой интересный эксперимент проверки первоочередных требований. В рамках спецификации современных стандартов, действия представителей оппозиции представляют собой не что иное, как квинтэссенцию победы маркетинга над разумом и должны быть превращены в посмешище, хотя само их существование приносит несомненную пользу обществу."


	coded := CodedByFibonachi(a, b, message, k)
	fmt.Println(coded)

	uncoded := CodedByFibonachi(a, b, coded, k)
	fmt.Println(uncoded)
}
