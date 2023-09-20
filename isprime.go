package exercices

// Renvoie true si le nombre nb(int) est un nombre premier, sinon false
func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
