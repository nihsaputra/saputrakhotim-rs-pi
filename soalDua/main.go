package main

import "fmt"

func main() {
	// total permen milik pak zaki
	var totalPermen int = 12

	// cucu pak zaki default 0
	var cucu = make(map[string]int)
	cucu["Abi"] = 0
	cucu["Bibi"] = 0
	cucu["Cibi"] = 0

	// masukan permen untuk cibi > 6
	var permenUntukCibi = 7

	// pengkondisian permencibi & total permen
	if permenUntukCibi > 6 && totalPermen >= permenUntukCibi+2 {
		for i := 0; i < 2; i++ {
			if i == 0 {
				// memberi permen kecibi
				cucu["Cibi"] = permenUntukCibi
				totalPermen -= permenUntukCibi
			} else {
				// apabila total permen ganjil maka akan di berikan 1 ke cibi
				if totalPermen%2 != 0 {
					cucu["Cibi"] += 1
					totalPermen -= 1
				}

				// sisa total permen dibagi 2 dan diberikan kepada cucu yg tersisa
				totalPermen /= (len(cucu) - 1)
				cucu["Abi"] = totalPermen
				cucu["Bibi"] = totalPermen
			}
		}
		fmt.Println(cucu) // cetak semua cucu dan permen yang didapat
	} else {
		fmt.Println("PERMEN KURANG")
	}

}
