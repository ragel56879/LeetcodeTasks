package main

import "fmt"

func main() {
	num := 123
	fmt.Println(hammingWeight(num))
}

func hammingWeight(n int) int {
	var x, total int
	var num []int
	// Переводим число в двоичное и считаем единицы
	for n > 0 {
		num = append(num, x*10+n%2)
		if n%2 == 1 {
			total++
		}
		n /= 2
	}
	return total
}
