package main

import "fmt"

func calcLuasTabung() {
	var T, r float32
	const phi = 3.14
	fmt.Print("Masukkan Tinggi Tabung: ")
	fmt.Scan(&T)
	fmt.Print("Masukkan Jari-Jari Tabung: ")
	fmt.Scan(&r)
	var Lp = 2*(phi*(r*r)) + 2*(phi*(r*T))
	fmt.Println("Hasil Luas Permukaan Tabung : ", Lp)
}

func calcGradeNilai() {
	var nama string
	var nilai float32
	fmt.Print("Masukkan Nama Siswa: ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Nilai Siswa: ")
	fmt.Scan(&nilai)
	switch {
	case nilai >= 80 && nilai <= 100:
		fmt.Print("Siswa Dengan Nama : ", nama, " Mempunyai ")
		fmt.Print("Nilai A")
	case nilai >= 65 && nilai <= 79:
		fmt.Print("Siswa Dengan Nama : ", nama, " Mempunyai ")
		fmt.Print("Nilai B")
	case nilai >= 50 && nilai <= 65:
		fmt.Print("Siswa Dengan Nama : ", nama, " Mempunyai ")
		fmt.Print("Nilai C")
	case nilai >= 35 && nilai <= 49:
		fmt.Print("Siswa Dengan Nama : ", nama, " Mempunyai ")
		fmt.Print("Nilai D")
	case nilai >= 0 && nilai <= 34:
		fmt.Print("Siswa Dengan Nama : ", nama, " Mempunyai ")
		fmt.Print("Nilai E")
	default:
		fmt.Print("Invalid Nilai")
	}
}

func calcFaktorBilangan() {
	var number int
	fmt.Print("Masukkan Angka Faktor Bilangan: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}

func calcPrimeNumber() {
	var number int
	result := 0

	fmt.Print("Masukkan Bilangan Prima: ")
	fmt.Scan(&number)
	for x := 2; x < number; x++ {
		if number%x == 0 {
			result += 1
		}
	}

	if result == 2 {
		fmt.Print("Bukan Bilangan Prima")

	} else {
		fmt.Print("Bilangan Prima")

	}

}

func palindrome() {
	// var word string
	// fmt.Print("Masukkan Kata: ")
	// fmt.Scan(&word)
	// var count int = 0

	// for i := 1; i < len(word)/2; i++ {
	// 	if word[i] != word[len(word)-i-1] {
	// 		count++
	// 	}
	// }

	// fmt.Print(count)

	// for i := len(word) - 1; i >= 0; i-- {
	// 	reverseWord += string(word[i])
	// }
	// fmt.Println("Hasil: ", reverseWord)
	// for i := range word {
	// 	if word[i] != reverseWord[i] {

	// 	}
	// }
	// if word == reverseWord {
	// 	fmt.Println("Palindrome")
	// } else {
	// 	fmt.Println("Bukan Palindrome")
	// }
}

func exponentiation() {
	var base, pangkat int
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&base)
	fmt.Print("Masukkan Pangkat: ")
	fmt.Scan(&pangkat)
	var result int = 1
	for i := 1; i <= pangkat; i++ {
		result = result * base
	}
	fmt.Println("Hasil dari perhitungan : ", result)
}

func playWithAsterisk(n int) {
	for i := n; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		for j := i; j < n; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func cetakTablePerkalian(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print(j * i)
			if j*i > 9 {
				fmt.Print("     ")
			} else {
				fmt.Print("      ")
			}
		}
		fmt.Println()
	}
}

func main() {
	// calcLuasTabung()
	// calcGradeNilai()
	// calcFaktorBilangan()
	// calcPrimeNumber()
	// palindrome()
	// exponentiation()
	// playWithAsterisk(5)
	// cetakTablePerkalian(5)
}
