package main

import "fmt"

func main() {
	var s string

	fmt.Scanln(&s)

	for s != "selesai" {
		fmt.Println(s)
		fmt.Scanln(&s)
	}
}
