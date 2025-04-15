package main

import "fmt"

func IsPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func main() {
	fmt.Println("Hello, World!")
}
