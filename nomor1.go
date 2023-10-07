package main

import (
	"fmt"
)

//Rumus Luas Persegi Panjang
func main() {
	//Deklarasi Variabel
	var panjangPersegiPanjang int
	var lebarPersegiPanjang int

	//Baris Input
	fmt.Print("Masukkan panjang dan lebar persegi panjang: ")
	fmt.Scanln(&panjangPersegiPanjang, &lebarPersegiPanjang)

	//Program
	var luasPersegiPanjang int = panjangPersegiPanjang * lebarPersegiPanjang

	//Baris Output
	fmt.Println("Luas dari persegi panjang adalah", luasPersegiPanjang)
}
