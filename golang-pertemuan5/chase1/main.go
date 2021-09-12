package main

import "fmt"

func tambah(numbOne, numbTwo int) int {
	return numbOne + numbTwo
}
func Kurang(numbOne, numbTwo int) int {
	return numbOne - numbTwo
}
func Kali(numbOne, numbTwo int) int {
	return numbOne * numbTwo
}
func Bagi(numbOne, numbTwo int) float64 {
	return float64(numbOne) / float64(numbTwo)
}
func Modulus(numbOne, numbTwo int) int {
	return numbOne % numbTwo
}

func main() {
	var iniTambah = fmt.Sprintf(tambah(5, 5))
	fmt.Println(iniTambah)
}
