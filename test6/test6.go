package main

func main() {
	//result := []string{"zaaa", "babab", "B", "1", "B1 MarketCom", "baba", "babaa", "baca", "b", "A", "2", "c", "C", "3", "azzz", "B1 Info"}
	//fmt.Println(result)
	//exercices.TriRapide(result, len(result))
	//fmt.Println(result)
	//exercices.SortWordArr(result)
	//fmt.Println(result)
	//if x := rand.Intn(1000); x%2 == 0 {
	//	fmt.Printf("%v est un nombre pair\n", x)
	//} else {
	//	fmt.Printf("%v est un nombre impair\n", x)
	//}
	////fmt.Println(rand.Intn(1000))
	////fmt.Println(rand.Int())
	//phrase := " Bonjour vous êtes sur le point de me trouver, mais j'aimerais être insaisissable"
	//mot := "êtes"
	//fmt.Printf("Le mot recherché : %v a été trouvé : %v\n", mot, exercices.RechercheMot(phrase, mot))

	//// les map
	//superMarketPrice := map[string]int{
	//	"prince": 8,
	//	"eau":    2,
	//	"viande": 6,
	//}
	//fmt.Println("prince")
	//
	//for key, value := range superMarketPrice {
	//	fmt.Println(key, value)
	//}
	//for _, value := range superMarketPrice {
	//	fmt.Println(value)
	//}
	//for key, _ := range superMarketPrice {
	//	fmt.Println(key)
	//}
	//listeMois := map[string]int{
	//	"Janvier":   31,
	//	"Février":   28,
	//	"Mars":      31,
	//	"Avril":     30,
	//	"Mai":       31,
	//	"Juin":      30,
	//	"Juillet":   31,
	//	"Aout":      31,
	//	"Septembre": 30,
	//	"Octobre":   31,
	//	"Novembre":  30,
	//	"Décembre":  31,
	//}
	//fmt.Println("listeMois")
	//// Notez que cet affichage ne donne pas les mois dans l'ordre
	//for key, value := range listeMois {
	//	fmt.Println(key, value)
	//}
	//nbJoursDansAnnee := 0
	//for _, value := range listeMois {
	//	nbJoursDansAnnee = nbJoursDansAnnee + value
	//}
	//fmt.Println(nbJoursDansAnnee)

	//// Les fonctions anonymes
	//
	//w := func() {
	//	fmt.Println("Je suis une fonction anonyme sans valeur return")
	//}
	//w()
	//z := func() string {
	//	fmt.Println("Je suis une fonction anonyme avec valeur return")
	//	return "Alex"
	//}()
	//fmt.Println(z)
	//
	//x, y := func() (int, int) {
	//	fmt.Println("Aucun paramètre. Deux return")
	//	return 5, 7
	//}()
	//fmt.Println(x)
	//fmt.Println(y)
	//
	//func(a, b int) {
	//	fmt.Println("A * A + B * B", a*a+b*b)
	//}(x, y)

	//// La fonction fmt.Printf
	//fmt.Printf("Salut à tous!\n")
	//x, y := 50, 50
	//fmt.Printf("Mon nombre (default) est %v\n", x)
	//fmt.Printf("Mon nombre (base 2)  est %b\n", x)
	//fmt.Printf("Mon nombre (base 8) est %o\n", x)
	//fmt.Printf("Mon nombre (base 10) est %d\n", x)
	//fmt.Printf("Mon nombre (base 16) est %X\n", x)
	//
	//fmt.Printf("La valeur de X est égale à la valeur de Y %t\n", x > y)
	//fmt.Printf("L'écriture scientifique de 1258,97682 -> %E\n", 1258.97682)
}

//// Pass par valeur
//func updateA(number int) {
//	number = 5
//}
