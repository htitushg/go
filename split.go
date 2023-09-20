package exercices

// Renvoie un tableau de chaines qui contient les chaines contenues dans s et qui sont séparées par la chaine sep
func Split(s, sep string) []string {
	var tableau []string
	var indexcaractere []int
	i := 0
	y := 0
	//fmt.Printf("s : %v, sep : %v", s, sep)
	for i < len(s) {
		indexcaractere = append(indexcaractere, Index(s[i:len(s)-1], string(sep)))
		//fmt.Printf("Ligne 14 i: %v, s : %v, sep : %v, indexcaractere: %v\n", i, s, sep, indexcaractere)
		if indexcaractere[y] >= 0 {
			if y > 0 {
				indexcaractere[y] += indexcaractere[y-1] + len(sep)
			}
			//fmt.Printf("ligne 19 i: %v, s : %v, sep : %v, indexcaractere[y]: %v, y: %v\n", i, s, sep, indexcaractere[y], y)
			tableau = append(tableau, s[i:indexcaractere[y]])
			//fmt.Printf("ligne 21 i: %v, s : %v, sep : %v, indexcaractere: %v, tableau[y]: %v, y: %v\n", i, s, sep, indexcaractere, tableau[y], y)
			i = i + len(tableau[y]) + len(sep) - 1
			//fmt.Printf("ligne 23 i: %v, len(s): %v s : %v, sep : %v, indexcaractere: %v, tableau[y]: %v, y: %v\n", i, len(s), s, sep, indexcaractere, tableau[y], y)
			y++
		} else {
			if i < len(s) {
				tableau = append(tableau, s[i:len(s)])
				return tableau
			}
		}
		i++
		//fmt.Printf("ligne 27 i: %v, len(s): %v s : %v, sep : %v, indexcaractere: %v, tableau: %v, y: %v\n", i, len(s), s, sep, indexcaractere, tableau, y)
	}
	return tableau
}
