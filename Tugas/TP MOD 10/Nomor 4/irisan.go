package main

import "fmt"

type set [2022]int

func exist(T set, n int, val int) bool {
	var z int
	var meet bool
	for z < n && meet == false {
		if T[z] == val {
			meet = true
		}
		z++
	}
	return meet
}
func inputSet(T *set, n *int) {
	var masukan int
	var i int
	var selesai bool
	fmt.Scan(&masukan)
	T[i] = masukan
	*n++
	i++
	for selesai == false {
		fmt.Scan(&masukan)
		if !exist(*T, *n, masukan) {
			T[i] = masukan
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
