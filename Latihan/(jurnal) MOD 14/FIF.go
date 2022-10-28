package main

import "fmt"

var pita string
var CC byte
var EOP bool
var indeks int

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}

func ADV() {
	indeks++
	CC = pita[indeks]
	EOP = CC == '.'
}

func main() {
	fmt.Scan(&pita)
	START()

	var nKata int = 0
	var status int = 1

	for !EOP {
		if status == 1 {
			if CC == 'F' {
				status = 2
			}
		} else if status == 2 {
			if CC == 'I' {
				status = 3
			} else {
				status = 1
			}
		} else if status == 3 {
			if CC == 'F' {
				nKata++
			}
			status = 1
		}
		ADV()
	}
	fmt.Println(nKata)
}
