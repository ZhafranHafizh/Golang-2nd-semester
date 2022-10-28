package main

import (
	"fmt"
)

type pemain struct {
	nama   string
	nama_2 string
	gol    int
	assist int
}

func main() {
	var player [1001]pemain
	var n int
	fmt.Scan(&n)
	inputdata(&player, n)
	sorting(&player)
	for i := 0; i < n; i++ {
		fmt.Print(player[i].nama, player[i].nama_2, player[i].gol, player[i].assist)
	}
}

func inputdata(player *[1001]pemain, n int) {
	var nama, nama_2 string
	var gol, assist int
	for i := 0; i <= n; i++ {
		fmt.Scanln(&nama, &nama_2, &gol, &assist)
		player[i] = pemain{nama, nama_2, gol, assist}
	}
}

func sorting(player *[1001]pemain) {
	for i := 0; i < len(player); i++ {
		for j := i; j < len(player); j++ {
			if player[j].gol == player[i].gol {
				if player[j].assist < player[i].assist {
					temp := player[i].gol
					temp2 := player[i]
					player[i].gol = player[j].gol
					player[j].gol = temp
					player[i] = player[j]
					player[j] = temp2
				}
			} else if player[j].gol < player[i].gol {
				temp := player[i].gol
				temp2 := player[i]
				player[i].gol = player[j].gol
				player[j].gol = temp
				player[i] = player[j]
				player[j] = temp2
			}
		}
	}
}
