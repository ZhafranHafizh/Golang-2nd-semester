package main

import "fmt"

func main() {
	var d1, d2 int
	var stop bool

	fmt.Scanln(&d1, &d2)
	jumlah := 0
	stop = (d1%2 == 0) && (d2%2 == 0)
	for !stop {
		if (d1%2 != 0) && (d2%2 != 0) {
			jumlah++
		}
		fmt.Scanln(&d1, &d2)
		stop = (d1%2 == 0) && (d2%2 == 0)
	}
	fmt.Println(jumlah)
}
