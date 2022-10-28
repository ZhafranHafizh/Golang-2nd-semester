package main

import "fmt"

func main() {
	var member string
	var a, b, c, d, e int
	var cashback, diskon float64

	diskon = 0
	fmt.Scanln(&member, &a, &b, &c, &d, &e)

	if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 3.1 * float64(a+b+c)

	} else if a%2 == 0 && b%2 == 0 && c%2 == 0 && d%2 == 0 && e%2 == 0 {
		cashback = 3.1 * float64(c+d+e)
	} else {
		diskon = 1.7 * float64(c+d+e)
		cashback = 3.1 * float64(a+b+e)
	}
	if member == "yes" {
		diskon *= 1.15
		cashback *= 1.15
	}
	if cashback > 35 {
		cashback = 35
	}
	if diskon > 50 {
		diskon = 50
	}
	fmt.Println(cashback, diskon)
}
