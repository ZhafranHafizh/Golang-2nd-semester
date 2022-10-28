package main

import "fmt"

func main() {
	var (
		n, i           int
		p1, p2, p3, p4 bool
		status         bool = true
	)

	fmt.Scanln(&n)
	i = 0
	for i < n && status {
		fmt.Scanln(&p1, &p2, &p3, &p4)
		status = p1 && p2 && p3 && p4
		i++

	}
	fmt.Println(status)
}
