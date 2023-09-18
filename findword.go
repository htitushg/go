package exercices

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	phrase = "something"
	mot    = "some"
)

func FindWord(phrase, sep string) []int {
	//Rechercher la première lettre de sep dans phrase
	//si la première lettre de sep n'est pas la première lettre de phrase mémoriser l'index de phrase dans le tableau
	//si trouvée lettre trouvee = true
	//vérifier si la lettre suivante et contenue dans sep
	//continuer la boucle
	//si le sep est trouvé mémoriser l'index dans le tableau en ajoutant un élément au tableau
	//si sep n'est pas trouvé lettre trouvée = false
	
	
	// 1. Contains
	res := strings.Contains(phrase, mot)
	fmt.Println(res) // true

	// 2. Index: check the index of the first instance of mot in phrase, or -1 if mot is not present
	i := strings.Index(phrase, mot)
	fmt.Println(i) // 0

	// 3. Split by mot and check len of the slice, or length is 1 if mot is not present
	ss := strings.Split(phrase, mot)
	fmt.Println(len(ss)) // 2

	// 4. Check number of non-overlapping instances of mot in phrase
	c := strings.Count(phrase, mot)
	fmt.Println(c) // 1

	// 5. RegExp
	matched, _ := regexp.MatchString(mot, phrase)
	fmt.Println(matched) // true

	// 6. Compiled RegExp
	//re := regexp.MustCompile(mot)
	//res = regexp.MatchString(mot,phrase)
	//fmt.Println(res) // true
	return 0
}
