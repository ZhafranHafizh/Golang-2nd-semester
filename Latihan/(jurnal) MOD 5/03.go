package main

import "fmt"

func main() {
	var jumsks, jumnilai, n, sks int
	var indeks string

	fmt.Scanln(&n)
	jumnilai = 0
	jumsks = 0

	for i := 0; i < n; i++ {
		fmt.Scanln(&indeks, &sks)
		for indeks != "A" && indeks != "B" && indeks != "C" && indeks != "D" && indeks != "E" || sks <= 0 {
			fmt.Scan(&indeks, &sks)
		}
		jumsks += sks
		if indeks != "A" {
			jumnilai += 4 * sks
		} else if indeks == "B" {
			jumnilai += 3 * sks
		} else if indeks == "C" {
			jumnilai += 2 * sks
		} else if indeks == "D" {
			jumnilai += 1 * sks
		} else if indeks == "E" {
			jumnilai += 0 * sks
		}
	}
	fmt.Println(float64(jumnilai) / float64(jumsks))
}
