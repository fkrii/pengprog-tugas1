package main

import (
	"fmt"
	"strings"
)

//Pengganda Digit
func main() {
	//Deklarasi Variabel
	var bilangan string
	var hasilGanda string

	//Baris Input
	fmt.Print("Masukkan dua digit angka yang ingin anda gandakan: ")
	fmt.Scan(&bilangan)

	//Program
	splitBilangan := strings.Split(bilangan, "")
	hasilGanda = splitBilangan[0] + splitBilangan[0] + splitBilangan[1] + splitBilangan[1]
	//for i:=0; i < len(splitBilangan); i++ {
	//	hasilGanda += splitBilangan[i] + splitBilangan[i]
	//}

	//Baris Output
	fmt.Println("Hasil penggandaan adalah", hasilGanda)
}