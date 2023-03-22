package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukan angka 12 atau 13 : ")
	fmt.Scan(&angka)
	if angka == 12 {
		fmt.Println("Angka adalah 12")
	} else if angka == 13 {
		fmt.Println("Angka adalah 13")
	} else {
		fmt.Println("Angka bukan 12 dan 13")
	}
}
