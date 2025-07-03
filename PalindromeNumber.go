package main

import "fmt"

func main() {
	x := 1212
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	n, num := x, 0

	for x > 0 {
		num = num*10 + x%10
		x /= 10
	}
	return n == num
}
