package exercices
//import "fmt"

func TrimAtoi(s string) int {
	resultat:=0
	premierchiffrerencontre:=false
	negatif:=false
	//var caractere string
	//index:=0
	var phrase string
	for i := 0; i < len(s); i++	{
		//caractere:=s[i]
		//fmt.Printf("string(s[i])= %v\n", string(s[i]))
		if string(s[i])=="-" && !premierchiffrerencontre{
			negatif=true
		}
		if IsNumeric(string(s[i])){
			phrase+=string(s[i])
			premierchiffrerencontre=true
			//fmt.Printf("Le caractere numero %v est %v\n", i, string(s[i]))
		}
	}
	//fmt.Printf("La phrase est %v\n", phrase)
	//fmt.Printf("La phrase est %v sa longueur est : %v\n", s,len(phrase))
	if(len(phrase)==0)||!premierchiffrerencontre{
		//fmt.Printf("Pas de nombre à afficher\n")
		return 0
	}else{
		for i:=len(phrase)-1;i>-1;i--{

			//fmt.Printf("le chiffre traité est : %v l'index est %v la puissance de 10 est : %d\n", int(phrase[i]-48),i,len(phrase)-1-i)
			//temp:=(int(phrase[i]-48))*RecursivePower(10, len(phrase)-1-i)
			//fmt.Printf("Le chifre traité %v avec sa puissance est %v, la puissance calculée est %v\n",int(phrase[i]-48),temp,RecursivePower(10, len(phrase)-1-i))
			resultat+=(int(phrase[i]-48))*RecursivePower(10, len(phrase)-1-i)
			//fmt.Printf("Le résultat provisoire est : %v \n",resultat)
		}
	}
	if negatif{
		resultat=-resultat
	}
	return resultat
}
