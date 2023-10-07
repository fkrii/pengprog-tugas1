package main
import (
	"fmt"
)

//Rumus Keliling Persegi Panjang
func main() {
	//Deklarasi Variabel
	var panjangPersegiPanjang int
	var lebarPersegiPanjang int

	//Baris Input
	fmt.Print("Masukkan panjang dan lebar persegi panjang: ")
	fmt.Scanln(&panjangPersegiPanjang, &lebarPersegiPanjang)

	//Program
	var kelilingPersegiPanjang int = (panjangPersegiPanjang * 2) + (lebarPersegiPanjang * 2)

	//Baris Output
	fmt.Println("Keliling dari persegi panjang adalah", kelilingPersegiPanjang)	
}