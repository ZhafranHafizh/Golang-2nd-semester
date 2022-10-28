package main

import "fmt"

func TambahData(tab *[9999]int, n *int) {

	var data int
	for {
		fmt.Scan(&data)
		if data != 9999 {
			tab[*n] = data
			*n++
		} else {
			break
		}
	}

}
