package exercices

// Renvoie dan un tableau d'entiers les nombres de min(int) à max(int) non compris

func MakeRange(min, max int) []int {
	if max < min {
		return nil
	} else {
		n1 := make([]int, max-min)
		minimum := min
		for i := min; i < max; i++ {
			n1[i-minimum] = min
			min++
		}
		return n1
	}

}
