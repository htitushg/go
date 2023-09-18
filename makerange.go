package exercices

//func remove(s []int, i int) []int {
//    s[i] = s[len(s)-1]
//    return s[:len(s)-1]
//}

func MakeRange(min, max int) []int{
	if max<min{
		return nil
	}else{
		n1 := make([]int, max-min)
		minimum:=min
		//const taille = 50
		//var n1 [taille]int
		for i:=min;i<max;i++{
			n1[i-minimum] = min
			min++
		}
		return n1
	}	
	
}