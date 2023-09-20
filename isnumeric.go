package exercices

// Renvoie true si la chaine s(string) ne contient que des caractères numériques, sinon false
func IsNumeric(s string) bool {
	phrase := []rune(s)
	estnumerique := true
	for i := 0; i < len(s); i++ {
		if !(phrase[i] >= 48 && phrase[i] < 58) {
			estnumerique = false
		}

	}
	return estnumerique
}
