package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var judul string
var penyanyi string
var vote int
var arr1 []string
var arrVote []int
var arr [][]string
var id int

func addData(judul string, penyanyi string) {
	if len(arr) == 0 {
		vote = 1

	} else {
		for i := 0; i < len(arr); i++ {
			if judul == arr[i][0] && penyanyi == arr[i][1] {
				vote = arrVote[i] + 1
				deleteData(i + 1)
			} else {
				vote = 1
			}
		}
	}

	arrVote = append(arrVote, vote)
	arr1 = []string{judul, penyanyi}
	arr = append(arr, arr1)

	sort.Slice(arr[:], func(i, j int) bool {
		return arrVote[i] > arrVote[j]
	})

	sort.Slice(arrVote[:], func(i, j int) bool {
		return arrVote[i] > arrVote[j]
	})

}

// func tambah(judul string, penyanyi string, vote int) bool {
// 	for i := 0; i <= len(arr); i++ {
// 		if judul == arr[0][i] || penyanyi == arr[i][0] || vote == arrVote[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

func deleteData(ID int) {
	var before, after [][]string
	before = arr[:ID-1]
	after = arr[ID:]
	arr = append(after, before...)
	arrVote = append(arrVote[:ID-1], arrVote[ID:]...)
}

func main() {
	var pilih int
	var x = true

	for x {
		fmt.Println("\n\t====Daftar Pilihan Voting Musik==== ")
		fmt.Println("\n1. Input Data Musik \n2. Hapus data Musik berdasarkan ID \n3. Tampilkan seluruh data musik beserta jumlah data yang tersimpan dalam list")
		fmt.Println("4. Menampilkan top 3 musik terfavorit \n5. Jumlah Vote \n6. Menampilkan seluruh data dengan penyanyi berinisial A \n0. Keluar")
		fmt.Print("Masukkan Pilihan\t\t: ")
		fmt.Scanln(&pilih)

		switch pilih {
		case 1:
			fmt.Print("\nMasukkan Judul Lagu\t\t: ")
			scan := bufio.NewReader(os.Stdin)
			judul, _ = scan.ReadString('\n')
			judul = strings.TrimSpace(judul)
			fmt.Print("Masukkan Penyanyi \t\t: ")
			penyanyi, _ = scan.ReadString('\n')
			penyanyi = strings.TrimSpace(penyanyi)

			addData(judul, penyanyi)
			fmt.Printf("%s, %s", strings.TrimSpace(judul), penyanyi)
		case 2:
			fmt.Println("ID \tJudul Lagu \t\tPenyanyi \t\tVote")
			for i := 0; i < len(arr); i++ {
				fmt.Printf("%d \t%s \t%s \t%d\n", i+1, arr[i][0], arr[i][1], arrVote[i])
			}

			fmt.Print("\nMasukkan ID musik yang akan dihapus: ")
			fmt.Scanln(&id)

			fmt.Printf("Judul: %s \t penyanyi: %s \nBERHASIL TERHAPUS", arr[id-1][0], arr[id-1][1])

			deleteData(id)

		case 3:
			fmt.Println("ID \tJudul Lagu \t\tPenyanyi \t\tVote")
			for i := 0; i < len(arr); i++ {
				fmt.Printf("%d \t%s \t\t%s \t\t%d\n", i+1, arr[i][0], arr[i][1], arrVote[i])
			}
		case 4:
			fmt.Println("TOP 3 MUSIC")
			fmt.Println("Rank \tNama Game \t\tpenyanyi \tJumlah Vote")
			n := 0
			if len(arr) < 3 {
				n = len(arr)
			} else {
				n = 3
			}
			for i := 0; i < n; i++ {
				fmt.Printf("%d \t%s \t\t%s \t\t%d\n", i+1, arr[i][0], arr[i][1], arrVote[i])
			}
		case 5:
			jumlah := 0
			for i := 0; i < len(arr); i++ {
				jumlah = jumlah + arrVote[i]
			}
			fmt.Printf("Jumlah seluruh Vote yang terkumpul = %d", jumlah)
		case 6:
			cek := 0
			for i := 0; i < len(arr); i++ {
				if strings.HasPrefix(arr[0][i], "A") || strings.HasPrefix(arr[0][i], "a") {
					fmt.Printf("%d \t%s \t\t%s \t\t%d\n", i+1, arr[i][0], arr[0][i], arrVote[i])
				} else {
					cek = cek + 1
				}
			}
			if cek == len(arr) {
				fmt.Println("Data tidak tersedia!")
			}

		case 0:
			os.Exit(0)
		}
	}
}
