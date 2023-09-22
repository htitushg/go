package main

import (
	"exercices"
	"fmt"
	"math/rand"
)

func main() {
	result := []string{"zaaa", "babab", "B", "1", "B1 MarketCom", "baba", "babaa", "baca", "b", "A", "2", "c", "C", "3", "azzz", "B1 Info"}
	fmt.Println(result)
	exercices.TriRapide(result, len(result))
	fmt.Println(result)
	exercices.SortWordArr(result)
	fmt.Println(result)
	if x := rand.Intn(1000); x%2 == 0 {
		fmt.Printf("%v est un nombre pair\n", x)
	} else {
		fmt.Printf("%v est un nombre impair\n", x)
	}
	//fmt.Println(rand.Intn(1000))
	//fmt.Println(rand.Int())

}
