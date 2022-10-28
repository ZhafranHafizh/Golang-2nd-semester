package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("permainan tebak angka antara 0 hingga 4")

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	rahasia := random.Intn(4)

	var guess int
	for try := 1; try <= 3; try++ {
		fmt.Printf("TRIAL %d\n", try)
		fmt.Println("tolong masukan nomor anda")
		fmt.Scan(&guess)

		if guess < rahasia {
			fmt.Printf("Maaf, tebakan salah ; nomor terlalu kecil\n")
		} else if guess > rahasia {
			fmt.Printf("Maaf, tebakan salah ; nomor terlalu besar\n")
		} else {
			fmt.Printf("Anda menang!\n")
			break
		}
		if try == 3 {
			fmt.Printf("Game Over!!\n")
			fmt.Printf("Nomor yang benar adalah %d\n", rahasia)
			break

		}
	}

}
