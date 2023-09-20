package exercices

// Renvoie un entier suite à l'application de la fonction f
func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for i := 0; i < len(tab); i++ {
		resultat := f(tab[i])
		if resultat {
			counter++
		}
	}
	return counter
}
