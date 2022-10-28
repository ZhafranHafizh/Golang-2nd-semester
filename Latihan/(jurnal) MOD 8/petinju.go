package main

import "fmt"

func main() {
	var n, s1, s2, s3, s4, s5, s6 int
	fmt.Scanln(&n)
	petinju1 := 0
	petinju2 := 0
	for i := 1; i <= n; i++ {
		fmt.Scanln(&s1, &s2, &s3, &s4, &s5, &s6)
		petinju1 += (s1 + s2 + s3)
		petinju2 += (s4 + s5 + s6)
	}
	fmt.Println("Petinju A:", petinju1, "dan petinju B:", petinju2)
	if petinju1 > petinju2 {
		fmt.Println("Pemenang adalah petinju A")
	} else if petinju2 > petinju1 {
		fmt.Println("Pemenang adalah petinju B")
	} else {
		fmt.Println("Draw")
	}
}
