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

func RechercheMot(phrase, mot string) int {
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
