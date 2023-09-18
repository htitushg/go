package exercices


func BasicJoin(elems [] string) string {
	phrase:=""
	for i := 0; i < len(elems); i++	{
			phrase +=  elems[i]
	}
	return phrase
}
