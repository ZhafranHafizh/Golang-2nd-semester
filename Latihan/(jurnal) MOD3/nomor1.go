package main

import "fmt"

func main() {
	var xA1, yA1, xB1, yB1 int
	var xA2, yA2, xB2, yB2 int
	var xA3, yA3, xB3, yB3 int
	var g1, g2, g3 float64

	fmt.Scanln(&xA1, &yA1, &xB1, &yB1)
	fmt.Scanln(&xA2, &yA2, &xB2, &yB2)
	fmt.Scanln(&xA3, &yA3, &xB3, &yB3)

	g1 = float64(yA1-yB1) / float64(xA1-xB1)
	g2 = float64(yA2-yB2) / float64(xA2-xB2)
	g3 = float64(yA3-yB3) / float64(xA3-xB3)

	fmt.Println(g1)
	fmt.Println(g2)
	fmt.Println(g3)
}
