package main

import "fmt"

func updateA(number int) int {
	number = 5
	return number
}
func updateB(item map[string]int) {
	item["bonbon"] = 4
	item["salade"] = 10
	//return number
}

func main() {

	// Pass par valeur
	number := 10
	fmt.Printf("Le nombre number est : %v  \n", number)
	number = updateA(number)
	fmt.Printf("Le nombre number est : %v  \n", number)

	// B: maps, functions
	superMarketPrice := map[string]int{
		"prince": 8,
		"eau":    2,
		"viande": 6,
	}
	fmt.Println(superMarketPrice)
	updateB(superMarketPrice)
	fmt.Println(superMarketPrice)

	// Les pointeurs
}
