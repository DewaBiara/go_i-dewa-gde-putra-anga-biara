package main

import "fmt"

func main() {
	var score int
	var name string

	fmt.Printf(" Masukan Nama Siswa  : ")
	fmt.Scanf("%s\n", &name)
	fmt.Printf(" Masukan Nilai Siswa : ")
	fmt.Scanf("%d\n", &score)

	fmt.Println("------------------------------")
	fmt.Println(" Nama Siswa  :", name)
	fmt.Print(" Nilai Siswa : ")
	if score > 100 || score < 0 {
		fmt.Print("Nilai yang dimasukkan salah")
		return
	}

	switch {
	case score >= 80:
		fmt.Print("A")
	case score >= 65:
		fmt.Print("B")
	case score >= 50:
		fmt.Print("C")
	case score >= 35:
		fmt.Print("D")
	case score >= 0:
		fmt.Print("E")
	}

	fmt.Println("\n------------------------------")
}
