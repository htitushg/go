package exercices

import "fmt"

// Affiche pour un tableau d'entiers, applique la fonction f
func ForEach(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
	fmt.Printf("\n")
}
