package exercices

// Renvoie une chaine(string) qui contient chaque string de elems séparées par le caractère contenu dans sep(string)
func Join(elems []string, sep string) string {
	phrase := elems[0]
	for i := 1; i < len(elems); i++ {
		phrase += sep + elems[i]
	}
	return phrase
}
