package main

import "fmt"

func main() {
	var n, nFaktor int

	nFaktor = 0
	fmt.Scanln(&n)
	fmt.Print("Fakotr :")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, " ")
			nFaktor++
		}
	}
	fmt.Println("\nPrima : ", nFaktor == 2)
}
