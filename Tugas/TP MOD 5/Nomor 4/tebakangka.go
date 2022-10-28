package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var masukan, pras, juri int
	var kepo int64

	fmt.Printf("Silahkan masukan nomor rahasia Anda")
	fmt.Scanln(&kepo)
	kepo = time.Now().UnixMilli()
	rand.Seed(kepo)
	fmt.Scanln(&masukan)

	pras = rand.Intn(5)
	fmt.Println(pras)
	juri = rand.Intn(5)
	fmt.Println(juri)

	switch {

	case masukan == juri && pras != juri:
		fmt.Print("Bilangan Anda sama dan Juri adalah", juri)

	case masukan != juri && pras == juri:
		fmt.Print("Bilangan Pak Pras dan Juri adalah", juri)

	case masukan == juri && pras == juri:
		fmt.Print("Bilangan Anda dan Pak Pras sama dengan Juri")

	default:
		fmt.Print("Bilangan Anda dan Pak Pras tidak sama dengan Juri")
	}
}
