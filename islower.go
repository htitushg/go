package exercices

func IsLower(s string) bool {
	phrase := []rune(s)
	estminuscule := true
	for i := 0; i < len(s); i++ {
		if phrase[i] >= 97 && phrase[i] < 123 {
			estminuscule = true
		} else {
			return false
		}

	}
	return estminuscule
}
