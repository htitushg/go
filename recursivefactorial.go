package exercices

// Renvoie la factorielle(int) de n(int) calculée de manière récursive
func RecursiveFactorial(n int) int {
	if n == 0 {
		return 1
	} else {

		return n * RecursiveFactorial(n-1)

	}
}
