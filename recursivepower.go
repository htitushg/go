package exercices

// Renvoie un entier résultat de nb puissance power, si nb<0 renvoie 0
func RecursivePower(nb int, power int) int {
	if nb == 0 || power == 0 {
		return 1
	}
	if power < 0 {
		return 0
	}
	if power > 1 {
		return nb * RecursivePower(nb, power-1)
	}
	return nb
}
