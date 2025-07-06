package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	num := 1048577
	fmt.Println(isPowerOfTwo(num))
	duration := time.Since(start)
	fmt.Println(duration)
}

func isPowerOfTwo(n int) bool {
	for i := 0; ; i++ {
		pow := int(math.Pow(float64(2), float64(i)))
		if pow == n {
			return true
		}
		if pow > n {
			return false
		}
	}
	return false
}
