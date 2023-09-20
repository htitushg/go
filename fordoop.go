package exercices

// Renvoie un booleen suite à l'application de la fonction f
func ForDoop(f func(a, b int) int, tab []int) int {
	resultat := 0
	// traiter le cas de l'addition : vérifier que tab[0] et tab[1] sont numériques
	// traiter le cas de la soustraction  idem
	// traiter le cas de la multiplication idem
	// traiter le cas de la division : idem + tab[1]!=0
	// traiter le cas du modulo: idem division

	for i := 0; i < len(tab); i++ {
		if i < len(tab)-1 {
			resultat = f(tab[i], tab[i+1])
			//fmt.Printf("resultat = %v\n", resultat)
			if resultat <= 0 {
				return resultat
			}
		}
	}
	return resultat
}
