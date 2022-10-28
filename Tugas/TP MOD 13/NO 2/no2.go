package main

import "fmt"

type tipeArray [10]int

func printArr(T tipeArray, n int) {

	if n < 10 {
		fmt.Println(T[n])
		printArr(T, n+1)
	}

}
func main() {
	var a tipeArray = tipeArray{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var b int = 0
	printArr(a, b)

}
