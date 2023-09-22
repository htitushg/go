package exercices

func SortWordArr(a []string) []string {
	// Tri à bulle
	tableautrie := false
	for i := len(a) - 1; i > -1; i-- {
		tableautrie = true
		for j := 0; j < i; j++ {
			if a[j+1] < a[j] {
				a[j], a[j+1] = a[j+1], a[j]
				tableautrie = false
			}
		}
		if tableautrie {
			return a
		}
	}
	return a
}

func ConcatSlices(S1 []string, S2 []string) []string {
	return append(S1, S2...)
}

func TriRapide(arr []string, n int) []string {

	// si la chaine est vide on retrourne la chaine sans changement
	if n <= 1 {
		return arr
	}
	// Trie les n-1 premiers elements
	TriRapide(arr, n-1)

	// Insérer le dernier élément à sa position corecte dans le tableau trié
	last := arr[n-1]
	j := n - 2

	/* Déplacer les éléments de arr[0..i-1], qui sont
	   supérieur à la clé, à une position supérieure
	   à la position actuelle */
	for j >= 0 && arr[j] > last {
		arr[j+1] = arr[j]
		j--
	}
	arr[j+1] = last
	return arr
}
