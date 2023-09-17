package exercices

import "errors"

func IterativeFactorial2(n int) (int, error) {
	if n > 50 {
		return 0, errors.New("Erreur: Le chiffre entré est trop grand!")
	}
	resultat := 1
	for i := n; i > 0; i-- {
		resultat = resultat * n
		n--
	}
	return resultat, nil
}
