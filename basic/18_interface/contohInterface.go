package main

import (
	"fmt"
	"math"
)

// membuat interface
type BangunRuang interface {
	area() float64
}

// membuat bagun lingkaran
type Lingkaran struct {
	x, y, radius float64
}

// membuat persegi
type Persegi struct {
	panjang, lebar float64
}

// define method dari lingkran
// implementasi dari BangunRuang.area()
func (lingkaran Lingkaran) area() float64 {
	return math.Pi * lingkaran.radius * lingkaran.radius
}

// define method dari persegi
// implementasi dari BangunRuang.area()
func (persegi Persegi) area() float64 {
	return persegi.panjang * persegi.lebar
}

func getHasil(hasil BangunRuang) float64 {
	return hasil.area()
}

func main() {
	lingkaran := Lingkaran{x: 0, y: 0, radius: 5}
	persegi := Persegi{panjang: 10, lebar: 4}

	fmt.Printf("hasil dari lingkaran adalah %f\n", getHasil(lingkaran))

	fmt.Printf("hasil dari persegi adalah %f\n", getHasil(persegi))
}
