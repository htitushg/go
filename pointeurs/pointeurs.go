package main

import "fmt"

func update(pointerOfNumber *int, value int) {
	*pointerOfNumber = value
}

func main() {

	// a: string, int, bool, float, array
	// B: maps, functions
	number := 10
	fmt.Printf("Le nombre number est : %v  \n", number)
	fmt.Printf("Adresse mémoire de number est : %v  \n", &number)
	myPointer := &number
	fmt.Printf("Adresse mémoire de number est : %v  \n", myPointer)
	fmt.Printf("La valeur de l'adresse mémoire  %v  est %d\n", myPointer, *myPointer)
	update(myPointer, 100)
	fmt.Printf("Le nombre number est : %v  \n", number)

}
