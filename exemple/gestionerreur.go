package main

import (
	"bufio"
	"exercices"
	"fmt"
	"os"
	"strconv"
)

func division() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entre un chiffre : ")
		scanner.Scan()
		nbr, err := strconv.Atoi(scanner.Text())
		if err != nil { // Gestion de l'erreur de la fonction  Atoi()
			fmt.Println("Vous devenez rentrer un nombre et non une chaîne de caractères !")
		} else if nbr <= 0 {
			fmt.Println("[division par zéro impossible] Votre valeur doit être supérieur ou égal à 0")
		} else {
			fmt.Println("Résultat :", 1000/nbr)
			break
		}
	}
}
func CheckedAdd(a, b int) (sum int, ok bool) {
	defer func() {
		if recover() != nil {
			sum = 0
			ok = false
		}
	}()
	return a + b, true
}

func main() {
	arguments := os.Args[1:] //La ligne de commande avec le nom de l'exécutable
	println(arguments[0])
	println(arguments[1])
	println(arguments[2])
	if !(arguments[0] == nil || arguments[1] || arguments[2]) {

		//argumentsSansExe := os.Args[1:] //La ligne de commande sans le nom de l'exécutable
		//Argument2 := os.Args[2]
		//Argument3 := os.Args[3] //Accès direct au troisième élément de la ligne de commande
		//Argument1 := os.Args[1]
		//var TArgument[3] string

		// Vérifier le nombre d'arguments passés
		//for i := 0; i < 3; i++ {
		//	print(arguments[i])
		//}
		// traiter le cas de l'addition : vérifier que tab[0] et tab[1] sont numériques
		// traiter le cas de la soustraction  idem
		// traiter le cas de la multiplication idem
		// traiter le cas de la division : idem + tab[1]!=0
		// traiter le cas du modulo: idem division

		//fmt.Println(arguments)
		//fmt.Println(argumentsSansExe)
		//fmt.Println(Argument1)
		//fmt.Println(Argument2)
		//fmt.Println(Argument3)
		//var resultat = [2]int{0, 0}
		//nombre1:=0
		//nombre2:=0
		nombre1, retour1 := exercices.StringToNumeric(arguments[0])
		nombre2, retour2 := exercices.StringToNumeric(arguments[2])
		if retour1 != "" || retour2 != "" {
			//println("Vous devez entrer deux nombres séparés par un opérateur (+,-,*,/,%)")
			//os.Exit(3)
		} else {
			// les deux paramètres numériques sont entrés
			switch arguments[1] {
			case "+":
				// Il s'agit d'une addition
				println(nombre1 + nombre2)
			case "-":
				// Il s'agit d'une soustraction
				println(nombre1 - nombre2)
			case "*":
				// Il s'agit d'une multiplication
				println(nombre1 * nombre2)
			case "/":
				// Il s'agit d'une division
				if nombre2 == 0 {
					println("Vous devez entrer un nombre différent de zéro pour le diviseur!")
				} else {
					println(nombre1 / nombre2)
				}
			case "%":
				// Il s'agit d'une opération modulo
				if nombre2 == 0 {
					println("Vous devez entrer un nombre différent de zéro pour le modulo!")
				} else {
					println(nombre1 % nombre2)
				}

			}
		}
		//fmt.Printf("la somme de arguments[0] + arguments[2] = %v\n", resultat)
		////division()
		//fmt.Println("Fin")
		//somme:=0
		//ok:=false
		//somme, ok =CheckedAdd(Argument1,Argument3)
	}
}
