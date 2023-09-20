package exercices

import (
	//"fmt"
	"github.com/01-edu/z01"
)

// Affiche les runes contenues dans le tableau a(string) séparées par "\n"
func PrintWordsTables(a []string) {
	//fmt.Printf("Le tableau entré est : %v\n",a)
	resultat := Join(a, "\n")
	for i := 0; i < len(resultat); i++ {
		//fmt.Printf("   = La lettre traitée est : %v l'index i est : %v \n",rune(resultat[i]),i)
		z01.PrintRune(rune(resultat[i]))
	}
	z01.PrintRune(rune(10))
}
