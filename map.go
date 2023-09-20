package exercices

// Affiche pour un tableau d'entiers, applique la fonction f
func Map(f func(int) bool, a []int) []bool {
	var resultat []bool
	for i := 0; i < len(a); i++ {
		resultat = append(resultat, f(a[i]))
	}
	return resultat
}
