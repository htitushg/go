package exercices


func Join(elems [] string, sep string) string {
	phrase:=elems[0]
	for i := 1; i < len(elems); i++	{
			phrase +=  sep+elems[i]
	}
	return phrase
}
