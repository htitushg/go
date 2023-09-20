package exercices

// Renvoie le nombre premier(int) qui est égal ou qui suit n(int)
func FindNextPrime(n int) int {
	if IsPrime(n) {
		return n
	} else {
		for {
			n++
			if IsPrime(n) {
				return n
			}
		}
	}
}
