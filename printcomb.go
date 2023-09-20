package exercices

import "github.com/01-edu/z01"

// Affiche les nombres de 3 chiffres séparés par ", " avec le 1er chiffre < au second < au troisième
func PrintComb() {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 9; j++ {
			for k := j + 1; k < 10; k++ {

				z01.PrintRune(rune(i + 48))
				z01.PrintRune(rune(j + 48))
				z01.PrintRune(rune(k + 48))
				if !(i == 8 && j == 7 && k == 9) {
					z01.PrintRune(rune(44))
					z01.PrintRune(rune(32))
				}
			}
		}
	}

}
