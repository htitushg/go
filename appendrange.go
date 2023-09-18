package exercices

func AppendRange(min, max int) []int{
	if max<min{
		return nil
	}else{
		var n1 []int
		for i:=min;i<max;i++{
			n1 = append(n1, min)
			min++
		}
		return n1
	}	
	
}