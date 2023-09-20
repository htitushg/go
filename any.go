package exercices

// Affiche pour un tableau d'entiers, applique la fonction f
func Any(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		resultat := f(a[i])
		if !resultat {
			return false
		}
	}
	return true
}
