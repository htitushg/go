package exercices

import "github.com/01-edu/z01"

// Affiche toutes les combinaisons de 2 nombres de 2 chiffres séparés par "espace" et suivi de ",espace" avec le 1er nombre < au second
func PrintComb2() {
	x := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			for k := i; k < 10; k++ {
				if k == i {
					x = j + 1
				} else {
					x = 0
				}
				for l := x; l < 10; l++ {
					z01.PrintRune(rune(i + 48))
					z01.PrintRune(rune(j + 48))
					z01.PrintRune(rune(32))
					z01.PrintRune(rune(k + 48))
					z01.PrintRune(rune(l + 48))
					if !(i == 9 && j == 8 && k == 9 && l == 9) {
						z01.PrintRune(rune(44))
						z01.PrintRune(rune(32))
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
