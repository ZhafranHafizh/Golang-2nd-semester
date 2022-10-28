package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	var i int
	var ketemu bool
	for i < n && ketemu == false {
		if T[i] == val {
			ketemu = true
		}
		i++
	}
	return ketemu
}
func inputSet(T *set, n *int) {
	var inputan int
	var i int
	var selesai bool
	fmt.Scan(&inputan)
	T[i] = inputan
	*n++
	i++
	for selesai == false {
		fmt.Scan(&inputan)
		if !exist(*T, *n, inputan) {
			T[i] = inputan
			*n++

		} else {
			selesai = true
		}
		i++
	}
}
func findIntersection(T1, T2 set, n, m int, T3 *set, h *int) {
	var i int

	for i < n {
		if exist(T2, m, T1[i]) {
			T3[*h] = T1[i]
			*h++
		}
		i++
	}

}
func printSet(T set, n int) {

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", T[i])
	}
}
func main() {
	var s1, s2, s3 set
	var n1, n2, n3 int
	inputSet(&s1, &n1)
	inputSet(&s2, &n2)
	findIntersection(s1, s2, n1, n2, &s3, &n3)
	printSet(s3, n3)
}
