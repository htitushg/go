package exercices

import "github.com/01-edu/z01"

// Affiche 'T' si nb(int) <0, sinon 'F'
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune(rune(84))
	} else {
		z01.PrintRune(rune(70))
	}
}
