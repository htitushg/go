package exercices

// Renvoie la chaine convertie en nombre
func StringToNumeric(s string) (int, string) {
	resultat := 0
	if IsNumeric(s) {
		resultat = TrimAtoi(s)
		return resultat, ""
	} else {
		return resultat, "Erreur: Vous devez entrer deux nombres entiers !" //création de l'erreur
	}
}
