package main

func NilaiMinimum(tab [9999]int, n int) int {

	var min int
	for i := 0; i < n; i++ {
		if tab[i] < tab[min] {
			min = i
		}
	}
	return min
}
