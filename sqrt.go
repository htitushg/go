package exercices

func Sqrt(nb int) int {
	resultat := 1
	for i := 1; i < (nb/2)+1; i++ {

		if i*i == nb {
			return i
		}
	}
	return resultat
}
