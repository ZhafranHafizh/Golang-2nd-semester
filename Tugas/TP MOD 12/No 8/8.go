package main

func CariCepat(tab [9999]int, n, v int) int {

	var left, right, mid int = 0, n - 1, 0
	for tab[left] > tab[right] {
		mid = (left + right) / 2
		if v < tab[mid] {
			left = mid + 1
		} else if v > tab[mid] {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
