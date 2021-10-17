package main

import (
	"fmt"
)

var jumlah = 0

func pangkat(i int, j int) int {
	if i <= 1 {
		jumlah = jumlah + 1
		fmt.Print(1)
		return 1
	}
	y := pangkat(i-1, j) * j
	jumlah = jumlah + y
	fmt.Printf(" + %d", y)
	return y
}

func puluhan(i int, j int) int {
	if i <= 1 {
		jumlah = jumlah + j
		fmt.Print(j)
		return j
	}
	y := puluhan(i-1, j) * 10
	jumlah = jumlah + y
	fmt.Printf(" + %d", y)
	return y
}

func deret(i int) int {
	if i <= 1 {
		jumlah = jumlah + 1
		fmt.Print(1)
		return 1
	}
	y := deret(i-1) * 10
	jumlah = jumlah + y*i
	fmt.Printf(" + %d", y*i)
	return y
}
func main() {
	var batas int
	var kelipatan int
	var pilih int

	fmt.Println("- Soal no 1 dan 3")
	fmt.Println("- Soal no 2 dan 5")
	fmt.Println("- Soal no 4")
	fmt.Print("Masukkan pilihan soal \t:")
	fmt.Scan(&pilih)

	switch pilih {
	case 1, 3:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)
		fmt.Print("Masukkan kelipatan: ")
		fmt.Scan(&kelipatan)

		pangkat(batas, kelipatan)
		fmt.Printf(" = %d", jumlah)

	case 2, 5:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)
		fmt.Print("Masukkan angka awal: ")
		fmt.Scan(&kelipatan)

		puluhan(batas, kelipatan)
		fmt.Printf(" = %d", jumlah)

	case 4:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)

		deret(batas)
		fmt.Printf(" = %d", jumlah)

	}

}
