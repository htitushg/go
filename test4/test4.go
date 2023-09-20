package main

import (
	"fmt"
)

// Structure de l'automate
type Automate struct {
	text      string
	pattern   string
	nextState [][]int
}

// Initialisation de l'automate
func NewAutomate(text, pattern string) *Automate {
	m := len(pattern)
	nextState := make([][]int, m+1)

	for i := 0; i < m+1; i++ {
		nextState[i] = make([]int, 256)
	}

	nextState[0][pattern[0]] = 1
	x := 0
	for j := 1; j < m+1; j++ {
		for c := 0; c < 256; c++ {
			nextState[j][c] = nextState[x][c]
		}
		nextState[j][pattern[j]] = j + 1
		x = nextState[x][pattern[j]]
	}

	return &Automate{text, pattern, nextState}
}

// Recherche des occurrences
func (a *Automate) Search() {
	n := len(a.text)
	m := len(a.pattern)
	j := 0

	for i := 0; i < n; i++ {
		j = a.nextState[j][a.text[i]]
		if j == m {
			fmt.Println("Occurrence trouvée à l'index", i-m+1)
		}
	}
}

func main() {
	text := "abababab"
	pattern := "aba"

	automate := NewAutomate(text, pattern)
	automate.Search()
}
