package exercices

// Renvoie -1 si a<b, 0 si a=b, 1 si a>b
func Compare(a, b string) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}
