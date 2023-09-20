package main

import (
	"exercices"
	"fmt"
)

func main() {
	//exercices.PrintComb2()
	//exercices.PrintNbr(-123)
	//exercices.PrintNbr(0)
	//exercices.PrintNbr(123)
	//exercices.IsNegative(0)
	//exercices.IsNegative(-1)
	//arg := 120
	//fmt.Println(exercices.RecursiveFactorial(arg))
	//fmt.Println(exercices.IterativeFactorial(arg))
	//fmt.Println(exercices.RecursivePower(2, 4)
	//fmt.Println(exercices.Fibonacci(4))
	//fmt.Println(exercices.Fibonacci(0))
	//fmt.Println(exercices.Fibonacci(1))
	//fmt.Println(exercices.Fibonacci(2))
	//fmt.Println(exercices.Fibonacci(3))
	//fmt.Println(exercices.Fibonacci(-1))

	//fmt.Println(exercices.Sqrt(9))
	//fmt.Println(exercices.Sqrt(4))
	//fmt.Println(exercices.Sqrt(3))
	//fmt.Println(exercices.IsPrime(0))
	//fmt.Println(exercices.IsPrime(1))
	//fmt.Println(exercices.IsPrime(2))
	//fmt.Println(exercices.IsPrime(3))
	//fmt.Println(exercices.IsPrime(4))
	//fmt.Println(exercices.IsPrime(5))
	//fmt.Println(exercices.IsPrime(6))
	//fmt.Println(exercices.IsPrime(7))
	//fmt.Println(exercices.IsPrime(9))
	//fmt.Println(exercices.IsPrime(11))
	//fmt.Println(exercices.FindNextPrime(5))
	//fmt.Println(exercices.FindNextPrime(4))
	//fmt.Println(exercices.FindNextPrime(8))
	//fmt.Println(exercices.Index("Hello", "l"))
	//fmt.Println(exercices.Index("Salut", "alu"))
	//fmt.Println(exercices.Index("Salut", "au"))
	//fmt.Println(exercices.Index("Salutations de notre part", "notre"))
	//fmt.Println(exercices.Concat("Salutations", " de notre part"))
	//fmt.Println(exercices.IsLower("salutations de notre part"))
	//fmt.Println(exercices.ToUpper("Salutations de notre part"))
	//fmt.Println(exercices.ToUpper("salutations de notre part"))
	//mt.Println(exercices.IsPrintable("Salutations de notre part"))
	//mt.Println(exercices.IsPrintable("Salutations de notre part\n"))
	//fmt.Println(exercices.Capitalize("Hello! How are you? How+are+things+4you?"))
	//fmt.Println(exercices.Capitalize("salut tout le-monde"))
	//elems:=[]string{"Salutations", "de ","notre","part"}
	//fmt.Println(exercices.Join(elems, "=="))
	//fmt.Println(exercices.TrimAtoi("12345"))
	//fmt.Println(exercices.TrimAtoi("str123ing45"))
	//fmt.Println(exercices.TrimAtoi("012 345"))
	//fmt.Println(exercices.TrimAtoi("Hello World!"))
	//fmt.Println(exercices.TrimAtoi("sd+x1fa2w3s4"))
	//fmt.Println(exercices.TrimAtoi("sd-x1fa2w3s4"))
	//fmt.Println(exercices.TrimAtoi("sdx1-fa2w3s4"))
	//fmt.Println(exercices.TrimAtoi("sdx1+fa2w3s4"))
	//fmt.Println(exercices.AppendRange(5, 10 ))
	//fmt.Println(exercices.AppendRange(10, 5 ))
	//fmt.Println(exercices.MakeRange(5, 10 ))
	//fmt.Println(exercices.MakeRange(10, 5 ))
	//test:=[]string{"Hello","how","are","you?"}
	//fmt.Println(exercices.ConcatParams(test))
	//test1:=[]string{"Hello how are you?"}
	//fmt.Println(test1)
	//test2:="Hello how are you? bonjour"
	//fmt.Println(test2)
	//exercices.PrinWordsTables(exercices.SplitWhiteSpaces(test2))
	//s := "HelloHAhowHAareHAyou?HABonjour"
	//fmt.Printf("s[23:29]: %v", s[23:29])
	//fmt.Printf("%#v\n", exercices.Index(s[23:29], "HA"))
	//fmt.Printf("%#v\n", exercices.Split(s, "HA"))

	//a := []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(exercices.Map(exercices.IsPrime, a))

	//a1 := []string{"Hello", "how", "are", "you"}
	//a2 := []string{"This", "is", "4", "you"}
	//fmt.Println(exercices.Any(exercices.IsNumeric, a1))
	//fmt.Println(exercices.Any(exercices.IsNumeric, a2))

	//tab1 := []string{"Hello", "how", "are", "you"}
	//tab2 := []string{"This", "1", "is", "4", "you"}
	//fmt.Println(exercices.CountIf(exercices.IsNumeric, tab1))
	//fmt.Println(exercices.CountIf(exercices.IsNumeric, tab2))

	a3 := []int{1, 2, 3, 4, 5}
	a4 := []int{0, 2, 1, 3}
	result1 := exercices.IsSorted(exercices.CompareInt, a3)
	result2 := exercices.IsSorted(exercices.CompareInt, a4)
	fmt.Println(result1)
	fmt.Println(result2)

}
