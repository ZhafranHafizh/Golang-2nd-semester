package main

import "fmt"

func main() {
	var tahun int
	var k1, k2, k3, k4 bool

	fmt.Scan(&tahun)
	k1 = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
	fmt.Scan(&tahun)
	k2 = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
	fmt.Scan(&tahun)
	k3 = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0
	fmt.Scan(&tahun)
	k4 = tahun%400 == 0 || tahun%4 == 0 && tahun%100 != 0

	fmt.Println(k1 && k2 && k3 && k4)

}
