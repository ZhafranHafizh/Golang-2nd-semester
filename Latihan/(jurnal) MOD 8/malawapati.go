package main

import "fmt"

func main() {
	var n1, n2 int
	var x int = 1
	fmt.Scanln(&n1, &n2)
	if n1 > n2 {
		x = n1
	} else {
		x = n2
	}
	for {
		if x%n1 == 0 && x%n2 == 0 {
			fmt.Printf("%d", x)
			break
		}
		x++
	}
}
