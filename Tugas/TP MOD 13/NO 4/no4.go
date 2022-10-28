package main

import "fmt"

func desimalkebiner(n int) int {
	if n == 0 {
		return 0
	} else {
		return n%2 + (desimalkebiner(n/2) * 10)
	}
}

func main() {
	fmt.Println(desimalkebiner(100))
}
