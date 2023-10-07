package main

import "fmt"

func main() {
	var (
		b   bool   //true ou false
		s   string // "une phrase"
		i   int    // un entier
		u   uint   // un entier non signé
		u8  uint8  // <256
		i8  int8   // -128 à +127
		i16 int16
		u16 uint16
		f   float32
	)

	y := 17 //Variable déclarée et initializée en même tempsvariable assignée après la déclaration
	// ne fonctionne qu'à l'intérieur d'une fonction
	b = true
	s = "Alex"
	i = -15
	u = 3640
	u8 = 255
	i8 = -120
	i16 = -21500
	u16 = 40000
	f = 3.14

	fmt.Println(b, s, i, u, u8, i8, i16, u16, f, y)
}
