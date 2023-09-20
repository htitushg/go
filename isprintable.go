package exercices

// Renvoie true si la chaine s(string) ne contient que des caractères imprimables (>31)
func IsPrintable(s string) bool {
	phrase := []rune(s)
	for i := 0; i < len(phrase); i++ {
		if !(phrase[i] > 31) {
			return false
		}
	}
	return true
}
