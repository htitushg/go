package exercices

// Recherche la 1ère occurence d'une lettre dans s et renvoie l'index de cette lettre
func Index(s string, tofind string) int {
	phrase := []rune(s)
	mot := []rune(tofind)
	trouve := 0
	motencours := false
	index := -1
	// on recherche une lettre de mot dans phrase
	y := 0
	for i := 0; i < len(phrase); i++ {
		if mot[y] == phrase[i] {
			// On a trouvé la lettre du mot dans la phrase
			if !motencours {
				// si on a pas déjà trouvé une lettre on mémorise l'index de la lettre trouvée
				index = i
			}
			motencours = true
			trouve++
			if trouve == len(mot) && motencours {
				return index
			} else {
				y++
			}
		} else {
			trouve = 0
			motencours = false
			index = -1
			y = 0
		}

	}
	return index
}
