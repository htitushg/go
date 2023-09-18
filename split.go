package exercices


func Split(s, sep string) [] string {
	var tableau [] string
	motencours:=false
	var mot string
	premièrelettretrouvée:=false
	for i:=0 ; i < len(s); i++	{
		if s[i]==sep[0]{
			premièrelettretrouvée=true
		}
		if ((s[i] == sep[1]&&premièrelettretrouvée)||i== len(s)-1)&& mot!=""{
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
