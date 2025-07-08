package main

import "fmt"

func main() {
	digits := []int{9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	newDigits := []int{}
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			break
		} else {
			digits[i] = 0
		}
		if digits[0] == 0 {
			newDigits = append([]int{1}, digits...)
			return newDigits
		}
	}
	return digits
}
