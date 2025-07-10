package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))
}

func addBinary(a string, b string) string {
	// создаем новые big.Int для чисел
	a1 := new(big.Int)
	b1 := new(big.Int)

	// преобразуем двоичную строку в десятичное число
	a1.SetString(a, 2)
	b1.SetString(b, 2)

	// складываем числа
	sum := new(big.Int).Add(a1, b1)

	// преобразуем в двоичную строку и выводим
	return sum.Text(2)
}
