package main

import "fmt"

func main() {
	var s1, s2, s3 string
	var hasil bool

	fmt.Scanln(&s1, &s2, &s3)
	hasil = s1 == s2 || s1 == s3 || s2 == s3
	fmt.Println(hasil)
}
