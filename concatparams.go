package exercices

// Renvoie dans une string les strings contenus dans le tableau args, séparés par "\n"
func ConcatParams(args []string) string {
	return Join(args, "\n")
}
