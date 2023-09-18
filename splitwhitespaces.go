package exercices


func SplitWhiteSpaces(s string) [] string {
	var tableau [] string
	motencours:=false
	var mot string
	for i:=0 ; i < len(s); i++	{
		if (s[i] == 32||i== len(s)-1)&& mot!=""{
			if i== len(s)-1{
				mot+=string(s[i])
			}
			if motencours{
				tableau = append(tableau,mot)
				mot=""
			}
		}else{
			mot+=string(s[i])
			motencours=true
		}	
	}
	return tableau
}
