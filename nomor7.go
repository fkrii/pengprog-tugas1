package main
import (
	"fmt"
)

//Permainan Tebak-Tebakan
func main() {
	//Deklarasi Variabel
	var bilanganAdik int
	var bilanganKakak int

	//Baris Input
	fmt.Print("Masukkan dua bilangan antara 0-9: ")
	fmt.Scan(&bilanganAdik, &bilanganKakak)

	//Program
	var adikMenang bool = bilanganAdik % 2 == 1

	//Baris Output
	fmt.Println("Apakah adik menang?", adikMenang)
}