package main

import (
	"fmt"
)

type Trec struct {
	v1 int
	vx struct {
		v2, v3 int
	}
	v4 int
}

func BanyakNilai(rec *Trec) {
	n := []int{rec.v1, rec.vx.v2, rec.vx.v3, rec.v4}

	for i, min := 0, n[0]; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
		rec.v1 = min
	}

	rec.vx.v2 = 0
	for _, v := range n {
		rec.vx.v2 += v
	}

	rec.vx.v3 = rec.vx.v2 / len(n)

	for i, max := 0, n[0]; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
		rec.v4 = max
	}
}

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

func CariSekuensial(tab [9999]int, v int) int {

	for idx, val := range tab {
		if val == v {
			return idx
		}
	}
	return -1
}

func NilaiMinimum(tab [9999]int, n int) int {

	var min int
	for i := 0; i < n; i++ {
		if tab[i] < tab[min] {
			min = i
		}
	}
	return min
}

func NilaiRerata(tab [9999]int, n int) int {
	var result int
	for _, val := range tab {
		result += val
	}
	return result / n
}

func tukar(a *int, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

func TerurutA(tab *[9999]int, n int) {

	var i, j, key int
	for i = 0; i < n; i++ {
		key = i
		for j = i + 1; j < n; j++ {
			if tab[j] < tab[key] {
				key = j
			}
		}
		if tab[i] != tab[key] {
			tukar(&tab[i], &tab[key])
		}
	}
}

func TerurutB(tab [9999]int, n int) {

	var i, j, key int
	for i = 1; i < n; i++ {
		key = tab[i]
		j = i - 1
		for j >= 0 && key > tab[j] {
			tab[j+1] = tab[j]
			j = j - 1
		}
		tab[j+1] = key
	}
}

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

func Shaggy(tab [9999]Trec, n int) {
	var (
		i                     int
		found1, found2, found bool
	)
	found = false
	i = 2

	for i < n && !found {
		found1 = (tab[i-1].v1 == tab[i].vx.v2)
		found2 = (tab[i].vx.v3 == tab[i].v4)
		found = found1 && found2
		i += 1
	}

	if found {
		fmt.Println("Ada Shaggy disana. Yaitu …?")
	}
}
