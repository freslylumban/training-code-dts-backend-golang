package main

import (
	"fmt"
	"math"
)

//FUNCTION
// func ShowName(name string) string {
// 	return name
// }

// func addition(numberOne int, numberTwo int) {
// 	fmt.Println(numberOne + numberTwo)
// }

// func tambah(numbOne int, numbTwo int) int {
// 	return numbOne + numbTwo
// }
// func kurang(numbOne int, numbTwo int) int {
// 	return numbOne - numbTwo
// }
// func kali(numbOne int, numbTwo int) int {
// 	return numbOne * numbTwo
// }
// func bagi(numbOne int, numbTwo int) float64 {
// 	return float64(numbOne) / float64(numbTwo)
// }
// func modulus(numbOne int, numbTwo int) int {
// 	return numbOne % numbTwo
// }
// func akar(number int) float64 {
// 	return math.Sqrt(float64(number))
// }

// METHOD
// type Persegi struct {
// 	Sisi float64
// }

// // traditional function
// func Luas(sisi Persegi) float64 {
// 	return sisi.Sisi * sisi.Sisi
// }

// // Method with receiver param
// func (s Persegi) Luas() float64 {
// 	return s.Sisi * s.Sisi
// }

// Chase
// type PersegiPanjang struct {
// 	Panjang float64
// 	Lebar   float64
// }

// func (p PersegiPanjang) Luas() float64 {
// 	return p.Panjang * p.Lebar
// }

// func (p PersegiPanjang) Keliling() float64 {
// 	return 2*p.Panjang + 2*p.Lebar
// }

// type Lingkaran struct {
// 	Jarijari float64
// }

// func (l Lingkaran) Luas() float64 {
// 	return math.Pi * l.Jarijari * l.Jarijari
// }

// func (l Lingkaran) Keliling() float64 {
// 	return 2 * math.Pi * l.Jarijari
// }

//INTERFACE
type BangunDatar interface {
	Luas() float64
	Keliling() float64
}

type Persegi struct {
	S float64
}

type PersegiPanjang struct {
	P float64
	L float64
}

type Lingkaran struct {
	R float64
}

func (p Persegi) Luas() float64 {
	return p.S * p.S
}
func (p Persegi) Keliling() float64 {
	return 4 * p.S
}
func hitungPersegi(p Persegi) {
	fmt.Println(p.Luas())
	fmt.Println(p.Keliling())
}

func (pp PersegiPanjang) Luas() float64 {
	return pp.P * pp.L
}
func (pp PersegiPanjang) Keliling() float64 {
	return 2*pp.P + 2*pp.L
}
func hitungPersegiPanjang(pp PersegiPanjang) {
	fmt.Println(pp.Luas())
	fmt.Println(pp.Keliling())
}

func (ling Lingkaran) Luas() float64 {
	return math.Pi * ling.R * ling.R
}
func (ling Lingkaran) Keliling() float64 {
	return 2 * math.Pi * ling.R
}
func hitungLingkaran(ling Lingkaran) {
	fmt.Println(ling.Luas())
	fmt.Println(ling.Keliling())
}

func main() {
	//FUNCTION
	// fmt.Println(ShowName("Alpha"))
	// addition(5, 6)

	// fmt.Println(tambah(5, 5))
	// fmt.Println(kurang(5, 5))
	// fmt.Println(kali(5, 5))
	// fmt.Println(bagi(5, 3))
	// fmt.Println(modulus(5, 4))
	// fmt.Println(akar(7))

	//METHOD
	// persegi1 := Persegi{5}
	// fmt.Println(Luas(persegi1))
	// fmt.Println(persegi1.Luas())

	// persegiPanjang1 := PersegiPanjang{5, 6}
	// fmt.Println("Luas:", persegiPanjang1.Luas())
	// fmt.Println("Keliling:", persegiPanjang1.Keliling())

	// lingkaran1 := Lingkaran{10}
	// fmt.Println(lingkaran1.Luas())
	// fmt.Println(lingkaran1.Keliling())

	//INTERFACE
	p := Persegi{S: 5}
	hitungPersegi(p)

	pp := PersegiPanjang{P: 3, L: 4}
	hitungPersegiPanjang(pp)

	l1 := Lingkaran{R: 7}
	hitungLingkaran(l1)
}
