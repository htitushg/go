package exercices

import "fmt"

// Renvoie un booleen suite à l'application de la fonction f
func IsSorted(f func(a, b int) int, tab []int) bool {
	resultat := 0
	for i := 0; i < len(tab); i++ {
		if i < len(tab)-1 {
			resultat = f(tab[i], tab[i+1])
			fmt.Printf("resultat = %v\n", resultat)
			if resultat <= 0 {
				return false
			}
		}
	}
	return true
}
