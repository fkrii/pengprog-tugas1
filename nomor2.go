package main

import (
	"fmt"
	"math"
)

//Rumus Volume Bola
func main() {
	//Deklarasi Variabel
	var varPi float64 = 3.14
	var jariJariBola float64

	//Baris Input
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scanln(&jariJariBola)

	//Program
	var volumeBola float64 = varPi * math.Pow(jariJariBola, 3) * 4 / 3

	//Baris Output
	fmt.Println(math.Pow(jariJariBola, 2))
	fmt.Println("Volume bola adalah", volumeBola)
}
