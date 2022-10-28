package main

import "fmt"

func main() {
	var n, a, r, x int

	fmt.Scanln(&n, &a, &r)

	fmt.Print("0")
	for i := 1; i <= n; i++ {
		x = a * i * r
		fmt.Print(" + ", x)
	}
}
