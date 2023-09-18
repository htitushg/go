package exercices

func ToLower(s string) string {
	phrase := []rune(s)
	// on recherche une lettre minuscule dans phrase
	for i := 0; i < len(phrase); i++ {
		if phrase[i] >= 65 && phrase[i] < 91 {
			// On a trouvé un caractère alphabetique dans la phrase
			phrase[i] += 32
		}
	}
	return string(phrase)
}
