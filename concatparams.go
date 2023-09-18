package exercices
import "fmt"

// Recherche la 1ère occurence d'une lettre dans s et renvoie l'index de cette lettre
func ConcatParams(args [] string) string {
	fmt.Printf("Le tableau entré est : %v\n",args)
	return Join(args, "\n")
}
