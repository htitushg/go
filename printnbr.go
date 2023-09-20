package exercices

// Affiche 0(rune) si n(int)=0, sinon -n (runes)
import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var resultat []rune
	if n < 0 {
		n = -n
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune(48)
	} else {
		for n != 0 {
			u := n % 10
			resultat = append(resultat, rune(u)+48)
			n = n / 10
		}
		for i := len(resultat) - 1; i > -1; i-- {
			z01.PrintRune(resultat[i])
		}
	}
}
