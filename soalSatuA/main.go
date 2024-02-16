package main

import (
	"fmt"
)

func main() {
	// Masukan Saklar Yang Ingin Di Nyalakan A - E
	var saklar []string = []string{"A", "B", "C", "D", "E"}

	// Mengembalikan value maximalLampuMenyala
	value := saklarLampu(saklar)

	// Mencetak value maximalLampuMenyala
	fmt.Println(value)
}

func maximalLampuMenyala(lampu []bool) int {
	var result int = 0

	// pengecekan lampu yang true
	for i := 0; i < len(lampu); i++ {
		if lampu[i] == true {
			result += 1
		}
	}
	return result
}

func saklarLampu(saklar []string) int {
	// Lampu default false
	var lampu []bool = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false}
	var totalLampu = len(lampu) // total lampu dari panjang lampu

	// pengecekan saklar yang dinyalakan
	for i := 0; i < len(saklar); i++ {
		if saklar[i] == "A" {
			for j := 1; j <= totalLampu; j++ {
				if j%2 == 0 {
					lampu[j-1] = cekIfMenyala(lampu[j-1])
				}
			}
		} else if saklar[i] == "B" {
			for j := 1; j <= totalLampu; j++ {
				if j%3 == 0 {
					lampu[j-1] = cekIfMenyala(lampu[j-1])
				}
			}
		} else if saklar[i] == "C" {
			for j := 1; j <= totalLampu; j++ {
				if j%5 == 0 {
					lampu[j-1] = cekIfMenyala(lampu[j-1])
				}
			}
		} else if saklar[i] == "D" {
			for j := 1; j <= totalLampu; j++ {
				if totalLampu%7 == 0 {
					lampu[j-1] = cekIfMenyala(lampu[j-1])
				}
			}
		} else if saklar[i] == "E" {
			for j := 1; j <= totalLampu; j++ {
				if j%1 == 0 {
					lampu[j-1] = cekIfMenyala(lampu[j-1])
				}
			}
		}
	}

	// mengembalikan nilai maximal lampu yang menyala
	return maximalLampuMenyala(lampu)
}

func cekIfMenyala(lampu bool) bool {
	if lampu == true {
		return false
	}
	return true
}
