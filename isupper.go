package exercices

// Renvoie true si la chaine s(string) ne contient que des caractères en majuscule, sinon false
func IsUpper(s string) bool {
	phrase := []rune(s)
	estmajuscule := true
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 65 && phrase[i] < 91 {
			estmajuscule = true
		} else {
			return false
		}
	}
	return estmajuscule
}
