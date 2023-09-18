package exercices

func ToUpper(s string) string {
	phrase := []rune(s)
	// on recherche une lettre minuscule dans phrase
	for i := 0; i < len(phrase); i++ {
		if phrase[i] >= 97 && phrase[i] < 123 {
			// On a trouvé un caractère alphabetique dans la phrase
			phrase[i] -= 32
		}
	}
	return string(phrase)
}
