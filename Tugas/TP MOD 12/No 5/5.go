package main

func NilaiRerata(tab [9999]int, n int) int {
	var result int
	for _, val := range tab {
		result += val
	}
	return result / n
}
