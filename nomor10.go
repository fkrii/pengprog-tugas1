package main
import (
	"fmt"
)

func main() {
	//Deklarasi Variabel
	var panjang int
	var lebar int
	var tinggi int
	var berat int

	//Baris Input
	fmt.Print("Masukkan panjang, lebar, tinggi, dan berat paket: ")
	fmt.Scan(&panjang, &lebar, &tinggi, &berat)

	//Program
	var volumePaket int = (panjang * lebar * tinggi) / 100
	var beratKg int = berat / 1000
	var bolehKirim bool = volumePaket < 1 && beratKg < 30

	//Baris Output
	fmt.Println("Paket boleh dikirim?", bolehKirim)
}	