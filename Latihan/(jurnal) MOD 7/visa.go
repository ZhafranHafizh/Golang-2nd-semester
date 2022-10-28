package main

import (
	"fmt"
)

// Zhafran(1302210036)
func main() {
	// deklarasi variabel
	var (
		hari1, hari2, bulan1                       string
		tanggal1, tanggal2, bulan2, tahun1, tahun2 int
	)
	// pembacaan masukan
	fmt.Scan(&hari1, &tanggal1, &bulan1, &tahun1)
	stringKalender := map[string]int{
		"januari":   1,
		"februari":  2,
		"maret":     3,
		"april":     4,
		"mei":       5,
		"juni":      6,
		"juli":      7,
		"agustus":   8,
		"september": 9,
		"oktober":   10,
		"november":  11,
		"desember":  12,
	}
	intKalender := map[int]string{
		1:  "januari",
		2:  "februari",
		3:  "maret",
		4:  "april",
		5:  "mei",
		6:  "juni",
		7:  "juli",
		8:  "agustus",
		9:  "september",
		10: "oktober",
		11: "november",
		12: "desember",
	}
	// panggil subprogram untuk penentuan tanggal pengambilan
	pengambilan(tanggal1, stringKalender[bulan1], tahun1, hari1, &tanggal2, &bulan2, &tahun2, &hari2)
	if bulan2 > 12 {
		bulan2 = 1
	}
	// tampilkan tanggal pengambilan visa
	fmt.Println(hari2, tanggal2, intKalender[bulan2], tahun2)
}

// Zhafran(1302210036)
func kabisat(tahun int) bool {
	// Mengembalikan true apabila tahun adalah kabisat, false apabila sebaliknya.
	if tahun%400 == 0 {
		return true
	} else if tahun%100 == 0 {
		return true
	} else if tahun%4 == 0 {
		return true
	} else {
		return false
	}
}

// Zhafran(1302210036)
func angkaBulan(bulan string) int {
	/*
	   Mengembalikan angka berdasarkan urutan nama bulan pada kalender masehi (1 hingga 12).
	   0 untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua.
	   Contoh: "januari" menjadi 1
	*/
	kalender := map[string]int{
		"januari":   1,
		"februari":  2,
		"maret":     3,
		"april":     4,
		"mei":       5,
		"juni":      6,
		"juli":      7,
		"agustus":   8,
		"september": 9,
		"oktober":   10,
		"november":  11,
		"desember":  12,
	}
	return kalender[bulan]
}

// Zhafran(1302210036)
func bulanAngka(angka int) string {
	/*
	   Mengembalikan nama bulan berdasarkan urutan angka bulan pada kalender masehi (1 hingga 12).
	   "invalid" untuk bulan yang tidak valid. Asumsi nama bulan ditulis dengan huruf kecil semua.
	   Contoh: 1 menjadi "januari"
	*/
	angkaKalender := map[int]string{ // map untuk membuat array
		1:  "januari",
		2:  "februari",
		3:  "maret",
		4:  "april",
		5:  "mei",
		6:  "juni",
		7:  "juli",
		8:  "agustus",
		9:  "september",
		10: "oktober",
		11: "november",
		12: "desember",
	}

	if angkaKalender[angka] == "" {
		return "invalid"
	}
	return angkaKalender[angka]
}

// Zhafran(1302210036)
func jumlahHari(bulan, tahun int) int {
	/*
	   Mengembalikan jumlah hari berdasarkan bulan dan tahun yang terdefinisi,
	   hati-hati pada bulan februari pada saat kabisat.
	   -1 apabila bulan tidak valid
	*/
	hari := map[int]int{
		1:  31,
		2:  map[bool]int{true: 29, false: 28}[kabisat(tahun)], // apabila tahun kabisat pada bulan februari memiliki jumlah hari 29 maka true, jika bukan tahun kabisat maka false
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
		13: 30,
	}
	if hari[bulan] == 0 {
		return -1
	}
	return hari[bulan]
}

// Zhafran(1302210036)
func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	/*
	   I.S. terdefinisi hari1 yang menyatakan hari pengajuan string, asumsi huruf kecil semua
	   F.S. hari2 berisi hari pengambilan dan durasi berisi lama pembuatan visa,
	   karena sabtu dan minggu tidak dihitung
	*/
	const prosesPembuatan int = 2
	intHari := map[int]string{
		1: "senin",
		2: "selasa",
		3: "rabu",
		4: "kamis",
		5: "jumat",
		6: "sabtu",
		7: "minggu",
	}
	strHari := map[string]int{
		"senin":  1,
		"selasa": 2,
		"rabu":   3,
		"kamis":  4,
		"jumat":  5,
		"sabtu":  6,
		"minggu": 7,
	}
	if strHari[hari1] < 4 {
		*hari2 = intHari[(strHari[hari1] + prosesPembuatan)]
		*durasi = prosesPembuatan
	} else {
		*hari2 = intHari[((strHari[hari1] + prosesPembuatan) % 5)]
		*durasi = ((strHari[hari1] + prosesPembuatan) % 5) + ((5 - (strHari[hari1] % 5)) % 5) + prosesPembuatan
	}
}

// Zhafran(1302210036)
func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	/*
				I.S. terdefinisi waktu pengajuan visa, yaitu tanggal1, bulan1, tahun1 dan hari1
		   		F.S. tanggal2, bulan2, tahun2 dan hari2 berisi waktu pengambilan visa
	*/
	// tentukan durasi, hari pengambilan, dan jumlah hari pada bulan pengajuan
	var durasi int
	mencariDurasi(hari1, hari2, &durasi)
	var jumlahH = jumlahHari(bulan1, tahun1)

	// dapatkan tanggal pengambilan, asumsi bulan pengambilan dan tahun pengambilan sama dengan waktu pengajuan
	x := map[string]int{}
	if tanggal1+durasi > jumlahH {
		x["tanggal"] = (tanggal1 + durasi) % jumlahH
		x["bulan"] = bulan1 + 1
		if x["bulan"] > 12 {
			x["tahun"] = tahun1 + 1
		} else {
			x["tahun"] = tahun1
		}
	} else {
		x["tanggal"] = tanggal1 + durasi
		x["bulan"] = bulan1
		x["tahun"] = tahun1
	}
	// cek apabila tanggal pengambilan melebihi lama hari, update tanggal dan bulan pengambilan dengan yang seharusnya
	// cek apabila bulan pengambilan melebihi 12, update dengan bulan dan tahun pengambilan yang seharusnya
	*tanggal2 = x["tanggal"]
	*bulan2 = x["bulan"]
	*tahun2 = x["tahun"]
}
