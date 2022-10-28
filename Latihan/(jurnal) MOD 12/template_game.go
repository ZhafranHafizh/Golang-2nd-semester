package main
import "fmt"
import "math/rand" // modul random
import "time" // meminjam modul time untuk membangkitkan seed

type domino struct {
	left,right int
}
const NMAX = 28
type arrDomino [NMAX]domino
type arrInt [NMAX]int

func main(){
	var playerScore int = 0; var dealerScore int = 0
	var boneyard, playerTiles, dealerTiles arrDomino
	var decision, nBoneyard, pSkor, dSkor, total int
	fmt.Println("Welcome to the Algorithm and Programming Practicum")
	rand.Seed(time.Now().UnixNano())

	buatTile(&boneyard)

	nBoneyard = NMAX

	total = 0

	fmt.Println("Boneyard tiles:",nBoneyard)

	for decision == 0 && nBoneyard >= 4 {
		fmt.Printf("Your score is %d/%d \n\n",playerScore,total)

		acakTile(&boneyard, &nBoneyard)

		bagiTile(&boneyard, &playerTiles, &nBoneyard)

		bagiTile(&boneyard, &dealerTiles, &nBoneyard)

		printTile(playerTiles, "player")

		fmt.Println("Boneyard tiles:", nBoneyard)

		getDecision(&decision)

		for decision == 1 || decision == 2 {
			
			replaceTile(&boneyard, &playerTiles, &nBoneyard, decision)
			
			printTile(playerTiles, "player")

			fmt.Println("Boneyard tiles:", nBoneyard)

			getDecision(&decision)
		}
		if decision == 0{
			printTile (dealerTiles, "dealer")
			menang(playerTiles, dealerTiles, &pSkor, &dSkor)
			if pSkor == 1 {
				fmt.Println("you won")
				playerScore++
			}else if dSkor == 1 {
				fmt.Println("you lose")
				dealerScore++
			}
			total++
		}
	}
	fmt.Printf("Your last score is %d/%d \n",playerScore,total)
	fmt.Println("Thank you for playing with us.")
	fmt.Println("Your winning rate is",playerScore*100/total,"%")
}

func getDecision(decision *int){


	.... 
}

func buatTile(T *arrDomino){
	var kiri,kanan,i int
	i = 0
	for kiri = 0; kiri <= 6; kiri++ {
		for kanan = 0; kanan <= kiri; kanan++ {
			T[i].left = kiri
			T[i].right = kanan
			i++
		}
	}
}

func searchInt(T arrInt, n,x int) int{


	.... 
}

func acakTile(T *arrDomino, n int){
/* I.S. terdefinisi array T berisi sejumlah n Tile domino
	F.S. array T tersusun acak */
// BOLEH DIBANTU ASISTEN PRAKTIKUM
	fmt.Println("Dealing ...")
	var temp arrInt
	var temp2 arrDomino = *T
	// 1. buat array of integer (temp) yang berisi bilangan acak dari 1 hingga n
// 2. bilangan acak yang dibuat hanya akan ditambahkan ke array (temp) apabila belum ada di array
// 3. temp akan berisi bilangan 1 hingga n dengan susunan acak dan tidak ada duplikat
	...
// 4. gunakan temp sebagai indeks dari temp2 sehingga tile pada T akan tersusun acak
	...
}

func bagiTile(T,p *arrDomino, n *int){
/* I.S. terdefinisi array T yang berisi sejumlah n Tile domino
	F.S. pemain p memperoleh 2 Tile terakhir dari T, nilai n berkurang
	catatan: ambil Tile sejumlah m yang terakhir dari T, nilai n berkurang*/
// KERJAKAN MANDIRI
	for i := 0; i<*n;i++{
		p[i] = T[i+1]
	}
}

func replaceTile(T,p *arrDomino, n *int, posisi int){
/* I.S. terdefinisi array T yang berisi sejumlah n Tile domino, dan array p yang berisi sejumlah Tile pemain
	F.S. mengambil satu Tile dari T dan mengganti sesuai nomor posisi Tile pada array p*/
// KERJAKAN MANDIRI
	for i:=0;i<*n;i++{

	}
}

func haveBalak(p domino) bool{
/* Mengembalikan true apabila p adalah balak, false untuk kondisi sebaliknya */
// KERJAKAN MANDIRI
	return p.left == p.right
}

func tilePoin(p domino) int{
/* Mengembalikan total poin sisi left dan right dari p*/
// KERJAKAN MANDIRI
	...
}

func haveTwoBalak(p arrDomino) bool{
/* Mengembalikan true apabila pemain p memiliki 2 Tile balak. */
// KERJAKAN MANDIRI
	...
}

func menang(p1, p2 arrDomino, p1Skor,p2Skor *int){
/* I.S. terdefinisi 2 Tile yang masing-masing dimiliki oleh pemain p1 dan p2
F.S. p1Skor bernilai 1 dan p2Skor bernilai 0 apabila p1 menang, atau p1Skor bernilai 0 dan p2Skor bernilai 1 apabila p2 menang */
// KERJAKAN MANDIRI dan BOLEH BERTANYA KE ASISTEN PRAKTIKUM
	var poin1,poin2 int
	*p1Skor = 0; *p2Skor = 0
	// hitung total poin setiap pemain
	poin1 = ....
	poin2 = .... 
	// hanya p1 yang memiliki 2 balak
	if ... {
		...
	// hanya p2 yang memilikiki 2 balak
	}else if ... {
		...
	// untuk kondisi yang lain
	}else{
		if poin1 > poin2 {
			... 
		}else{
			... 
		}
	}
}

func printTile(p arrDomino, s string){
/* I.S. terdefinisi 2 Tile pada array p, dan sting s yang berisi nama pemain
F.S. menampilkan isi dari p dan s sesuai contoh masukan dan keluaran */
	var teks string
	if s == "player" {
		teks = "Your tiles"
	}else{ // s == "dealer"
		teks = "Dealer tiles"
	}
	fmt.Printf("%s: (%d,%d) (%d,%d)\n", teks,p[0].left, p[0].right, p[1].left, p[1].right)
}