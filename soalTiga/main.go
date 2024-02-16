package main

import "fmt"

func main() {
	var n int = 5  // input kartu n
	var q int = 45 // input q untuk loop

	// mengembalikan value kartu paling atas setelah loop
	value := kartu(n, q)

	// print kartu teratas saat ini
	fmt.Println(value)
}

func kartu(n int, q int) int {
	// kartu 1 - 31
	var kartu []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31}
	// kartu ke N dari terbawah
	var kartuN = len(kartu) - n

	for i := 0; i < q; i++ {
		// buat array baru lalu ambil kartu N dan taruh paling atas
		var newKartu []int = []int{kartu[kartuN]}

		// add kartu lama
		newKartu = append(newKartu, kartu[:kartuN]...)   // saat sebelum n
		newKartu = append(newKartu, kartu[kartuN+1:]...) // saat sesudah n

		// rubah kartu lama jadi yang baru
		kartu = newKartu
	}

	// kembalikan kartu dengan index teratas
	return kartu[0]
}
