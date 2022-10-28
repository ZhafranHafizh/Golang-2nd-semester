package main

import "fmt"

func up(n int) int {
	if n/10 == 0 {
		return 1
	}
	return 1 + up(n/10)
}

func main() {
	fmt.Print(up(231))
}
