package main

import "fmt"

func main() {
	var x int
	var y bool = false

	for !y {
		fmt.Scan(&x)
		if x%5 == 0 {
			y = true
		} else {
			y = false
		}
		fmt.Println(y)
	}
}
