package main

import "fmt"

type Trec struct {
	v1 int
	vx struct {
		v2, v3 int
	}
	v4 int
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
