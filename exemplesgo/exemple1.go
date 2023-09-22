package main

import "fmt"

func main() {

	// test des fonction de tri
	//result := []string{"zaaa", "B", "1", "B1 MarketCom", "b", "A", "2", "c", "C", "3", "azzz", "B1 Info"}
	//fmt.Println(result)
	//exercices.TriRapide(result, len(result))
	//fmt.Println(result)
	//exercices.SortWordArr(result)
	//fmt.Println(result)

	////Test de la condition if
	//if x := rand.Intn(1000); x%2 == 0 {
	//	fmt.Printf("%v est un nombre pair\n", x)
	//} else {
	//	fmt.Printf("%v est un nombre impair\n", x)
	//}
	////fmt.Println(rand.Intn(1000))
	////fmt.Println(rand.Int())

	////Boucle for 1
	//for i := 0; i < 5; i++ {
	//	fmt.Println(i)
	//}

	////Boucle for 2
	//x := 0
	//for x < 5 {
	//	fmt.Println(x)
	//	x++
	//}

	////Boucle for 3
	//y := 0
	//for {
	//	if y > 5 {
	//		break
	//	}
	//	fmt.Println(y)
	//	y++
	//}

	////Boucle for 4
	//z := 0
	//for ; z <= 10; z++ {
	//	if z%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(z)
	//}

	// Les tableaux
	//var list [3]int
	//list[0] = 10
	//list[1] = 20
	//list[2] = 30
	////for u, range(list){
	////	fmt.Println(list[u])
	////}
	//newlist := [...]int{10, 20, 30}
	//for k, v := range newlist {
	//	fmt.Printf("Position %d est égal à %d\n", k, v)
	//}
	//
	//// Les fonctions
	//func sayBye(name string) (string,error){
	//	if name == ""{
	//
	//	}
	//
	//}

	divide(10, 5)
	divide(20, 0)
	divide(30, 10)

}

func divide(x, y int) {
	if y == 0 {
		panic("You can not divide a number by Zero")

	} else {
		fmt.Printf("%d / %d = %d\n", x, y, x/y)
	}
}
