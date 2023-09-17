package exercices

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
