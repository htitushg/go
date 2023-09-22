package exercices

// Renvoie true si la chaine s(string) ne contient que des caractères numériques précédée ou non du signe + ou -, sinon false
//func IsNumeric(s string) bool {
//	phrase := []rune(s)
//	estnumerique := true
//	for i := 0; i < len(s); i++ {
//		if !(phrase[i] >= 48 && phrase[i] < 58) {
//			estnumerique = false
//		}
//
//	}
//	return estnumerique
//}

func IsNumeric(s string) bool {
	for index, char := range s {
		if index == 0 {
			if char == 43 || char == 45 {
				continue
			}
		}
		if char < 48 || char > 57 {
			return false
		}
	}
	return true
}
