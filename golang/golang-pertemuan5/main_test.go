package main

import "testing"

var (
	persegi      Persegi = Persegi{S: 7}
	luasExpected float64 = 49
	kelExpected  float64 = 28
)

func TestLuasPersegi(t *testing.T) {
	t.Logf("Luas persegi : %v", persegi.Luas())

	if persegi.Luas() != luasExpected {
		t.Errorf("Salah! Seharusnya : %v", luasExpected)
	}
}

func TestKelilingPersegi(t *testing.T) {
	t.Logf("Keliling persegi : %v", persegi.Keliling())

	if persegi.Keliling() != kelExpected {
		t.Errorf("Salah! Seharusnya kelilingnya : %v", kelExpected)
	}
}
