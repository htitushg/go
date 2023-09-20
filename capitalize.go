package exercices

// Renvoie dans une string tous les mots contenus dans s avec la première lettre de chaque mot en majuscule
import "fmt"

func Capitalize(s string) string {
	phrase := []rune(s)
	// on recherche la première lettre d'un mot
	motencours := false
	for i := 0; i < len(phrase); i++ {
		if IsLower(string(phrase[i])) || IsUpper(string(phrase[i])) {
			// On a trouvé un caractère alphabetique dans la phrase
			fmt.Printf("la lettre numéro %v est : %v\n", i, string(phrase[i]))
			if !motencours {
				motencours = true
				if IsLower(string(phrase[i])) {
					phrase[i] -= 32
					fmt.Printf("IsLower: la lettre numéro %v est : %v\n", i, string(phrase[i]))
				}
			} else {
				if IsUpper(string(phrase[i])) {
					fmt.Printf("IsUpper: la lettre numéro %v est : %v\n", i, string(phrase[i]))
					phrase[i] += 32
				}
			}
		} else {
			motencours = false
			continue
		}
	}
	return string(phrase)
}
