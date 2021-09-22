package main

import (
	"fmt"
	"testing"
)

var (
	hasilTambahExp int     = 10
	hasilKurangExp int     = 0
	hasilKaliExp   int     = 25
	hasilBagiExp   float64 = 1
	hasilModExp    int     = 0
)

func TestTambah(t *testing.T) {
	t.Logf("Hasil tambah 5: %v", fmt.Sprintf(Tambah(5, 5)))

	if fmt.Sprintf(Tambah(5, 5) != hasilTambahExp {
		t.Errorf("Salah! Seharusnya: %v", hasilTambahExp)
	}
}
