package exercices

// Renvoie la factorielle(int) de n(int) calculé de manière itérative
func IterativeFactorial(n int) int {
	resultat := 1
	for i := n; i > 0; i-- {
		resultat = resultat * n
		n--
	}
	return resultat
}
