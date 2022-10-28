package main

import "fmt"

func main() {
	var n, x int

	fmt.Scanln(&n)
	jumlah := 0
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		for x < 0 || x > 9 {
			fmt.Scanln(&x)
		}
		jumlah += x
	}
	fmt.Println(jumlah)
}
