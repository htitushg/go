package main

import (
	"fmt"
)

func main() {

	n1 := make([]int, 3, 4) //déclaration d'un tableau de 3 éléments avec une réservation de 4 éléments
	fmt.Println("n1:", n1)
	n2 := n1[1:2] //n2 est lié à n1 pour les éléments 1 jusqu'à 2 non compris
	fmt.Println("n2:", n2)

	fmt.Println("# n2[0] = 1")
	n2[0] = 1 // entraine que n1[1]= 1 à cause du lien
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# append(n2, 2)")
	n2 = append(n2, 2) // on ajoute 1 élémentt qui vaut 2 dans n2 cd qui entraine que n1[2]=2
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# n2[1] = 9")
	n2[1] = 9 // on change la valeur de n2[1]=9 donc n1[2]=9
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# append(n2, 3)")
	n2 = append(n2, 3) // on ajoute 1 élémentt qui vaut 3 dans n2[2] ce qui n'entraine pas n1[3]=3 puisque n1[3] n'existe pas
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# n2[1] = 8")
	n2[1] = 8  // entraine que n1[2]= 8 à cause du lien
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# append(n2, 4)")
	n2 = append(n2, 4)  // on ajoute 1 élémentt qui vaut 4 dans n2[3] ce qui n'entraine pas n1[4]=3 puisque n1[4] n'existe pas
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	fmt.Println("# n2[1] = 7")
	n2[1] = 7   // entraine que n1[2]= 7 à cause du lien
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)

	s1 := []string { "Bonjour", "Tout","le","Monde"}
	fmt.Println("s1:", s1)
	s2 := s1[1:2]
	fmt.Println("s2:", s2)
	fmt.Println(s1[0])
	phrase:=s1[0]
	fmt.Printf("Phrase = %v\n",phrase)
	phrase+=" "+s1[1]
	fmt.Printf("Phrase = %v\n",phrase)
	fmt.Println("# s2[0] = Au revoir")
	s2[0] = "Au revoir"
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# append(s2, toto)")
	s2 = append(s2, "toto")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# s2[1] = personne")
	s2[1] = "personne"
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# append(s2, 3)")
	s2 = append(s2, "tata")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# s2[1] = Henry")
	s2[1] = "Henry"
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# append(s2, bidule)")
	s2 = append(s2, "bidule")
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	fmt.Println("# s2[1] = 7")
	s2[1] = "Marie"
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	slice1 := []string{"Hello", "world", "this", "is", "Golinuxcloud"}
	slice2 := []string{"Welcome", "to", "our", "page"}
	slice3 := append(slice1, slice2...)           //append 2 slice
	slice4 := append(slice1[1:3], slice3[1:5]...) //append append overlapping slice

	fmt.Println("slice 1:", slice1)
	fmt.Println("slice 2:", slice2)
	fmt.Println("slice 3:", slice3)
	fmt.Println("slice 4:", slice4)
}