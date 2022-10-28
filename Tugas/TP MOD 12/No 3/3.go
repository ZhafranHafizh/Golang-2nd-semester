package main

func CariSekuensial(tab [9999]int, v int) int {

	for idx, val := range tab {
		if val == v {
			return idx
		}
	}
	return -1
}
