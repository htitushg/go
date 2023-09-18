package exercices

func IsAlpha(s string) bool {
	phrase := []rune(s)
	estalpha := true
	for i := 0; i < len(s); i++ {
		if !(phrase[i] >= 48 && phrase[i] < 58 || phrase[i] >= 65 && phrase[i] < 91 || phrase[i] >= 97 && phrase[i] < 123) {
			estalpha = false
		}

	}
	return estalpha
}
