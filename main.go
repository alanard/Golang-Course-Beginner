package main

import "fmt"

// Cetak gambar

func printPattern(numb int) {
	if numb%2 != 1 {
		fmt.Println("angka yang diinputkan harus berupa bilangan ganjil")
	} else {
		// Ini untuk baris kebawah
		for i := 0; i < numb; i++ {
			// Dan ini untuk kesamping
			for j := 0; j < numb; j++ {
				if j == 0 || i == numb/2 || j == numb-1 {
					fmt.Print("* ")
				} else {
					fmt.Print("= ")
				}
			}
			fmt.Println("\n")
		}
	}

}

func main() {
	printPattern(5)
}
