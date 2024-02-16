package main

import "fmt"

func main() {
	// Mengembalikan value lampu tidak bisa dinyalan/di Matikan
	value := lampuTidakBisaNyalaAtauMati()

	// Mencetak lampu tidak bisa dinyalan/di Matikan
	fmt.Println(value)
}

func lampuTidakBisaNyalaAtauMati() int {
	// Lampu default false
	var lampu []bool = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	var totalLampu = len(lampu) // total lampu dari panjang lampu

	var result int = 0 // result yang akan dikembalikan

	for i := 1; i <= totalLampu; i++ {

		if (i-1)%2 != 0 && (i-1)%3 != 0 && (i-1)%5 != 0 && (i-1)%7 != 0 && (i-1)%11 != 0 {
			result += 1 // pengecekan apabila semua kondisi != 0 (tidak bisa dinyalakan)
		} else if (i-1)%2 == 0 && (i-1)%3 == 0 && (i-1)%5 == 0 && (i-1)%7 == 0 && (i-1)%11 == 0 {
			result += 1 // pengecekan apabila semua kondisi == 0 (lampu selalu menyala)
		}
	}

	return result
}
