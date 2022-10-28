package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	konversi(x, &y)
	fmt.Println(y)
}

func pangkat(a, b int) int {
	var j, hasil int
	hasil = 1
	for j = 1; j <= b; j++ {
		hasil *= a
	}
	return hasil
}

func konversi(biner int, desimal *int) {
	var x, j int
	*desimal = 0
	j = 0
	for biner > 0 {
		x = biner % 10
		*desimal += x * pangkat(2, j)
		j++
		biner = biner / 10
	}
}
