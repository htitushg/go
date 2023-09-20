package exercices

import "fmt"

// Renvoie -1 si a<b, 0 si a=b, 1 si a>b
func CompareInt(a, b int) int {
	fmt.Printf("a= %v, b= %v\n", a, b)
	if a < b {
		return 1
	} else if a > b {
		return -1
	} else {
		return 0
	}
}
