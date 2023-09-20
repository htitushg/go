package exercices

// Renvoie dans une string les mots contenus dans le tableau de string elems
func BasicJoin(elems []string) string {
	phrase := ""
	for i := 0; i < len(elems); i++ {
		phrase += elems[i]
	}
	return phrase
}
