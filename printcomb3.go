package exercices

import "github.com/01-edu/z01"

// Affiche toutes les combinaisons de 2 nombres de 2 chiffres séparés par "espace" et suivi de ",espace" avec le 1er nombre < au second
func PrintComb3() {
	x := '0'
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				if k == i {
					x = j + 1
				} else {
					x = '0'
				}
				for l := x; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					if !(i == '9' && j == '8' && k == '9' && l == '9') {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
