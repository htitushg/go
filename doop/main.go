package main

import (
	"fmt"
	"os"
)

func printRunes(runes ...rune) {
	chars := make([]byte, len(runes))
	for index, r := range runes {
		chars[index] = byte(r)
	}
	os.Stdout.Write(chars)
}

func isMathOperator(op string) bool {
	for _, char := range op {
		if char == 37 || char == 42 || char == 43 || char == 45 || char == 47 {
			return true
		}
	}
	return false
}

func isNumeric(s string) bool {
	for index, char := range s {
		if index == 0 {
			if char == 43 || char == 45 {
				continue
			}
		}
		if char < 48 || char > 57 {
			return false
		}
	}
	return true
}

func trimAtoi(s string) int {
	if s != "" {
		var intInStr string
		for _, char := range s {
			if isNumeric(string(char)) {
				intInStr += string(char)
			}
		}
		return strToInt(intInStr)
	}
	return 0
}

func recursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	}
	return nb * recursivePower(nb, power-1)
}

func strToInt(s string) int {
	if s != "" {
		var isNegative bool
		if s[0] == byte('-') {
			isNegative = true
			s = s[1:]
		}
		if isNumeric(s) {
			var result int
			for index, digit := range s {
				result += (int(digit) - 48) * recursivePower(10, len(s)-1-index)
			}
			if isNegative {
				result = -result
			}
			return result
		}
	}
	return 0
}

func printNbr(n int) {
	if n < 0 {
		printRunes('-')
		n = -n
	}
	if n == 0 {
		printRunes('0')
	} else {
		var str string
		for i := 0; n != 0; i++ {
			str += string(rune(48 + n%10))
			n /= 10
		}
		chars := make([]rune, len(str))
		for index, char := range str {
			chars[len(str)-1-index] = char
		}
		for _, char := range chars {
			printRunes(char)
		}
	}
}

func main() {
	fmt.Printf("Nombre d'arguments : %v\n %#v\n", len(os.Args), os.Args)
	// fmt.Println(isNumeric(os.Args[1]))
	// fmt.Println(isMathOperator(os.Args[2]))
	// fmt.Println(isNumeric(os.Args[3]))
	// fmt.Println(len(os.Args[1]) + len(os.Args[3]))
	if len(os.Args) == 4 {
		if isNumeric(os.Args[1]) && isMathOperator(os.Args[2]) && isNumeric(os.Args[3]) {
			if len(os.Args[1])+len(os.Args[3]) < 16 {
				if os.Args[2] == "+" {
					printNbr(trimAtoi(os.Args[1]) + trimAtoi(os.Args[3]))
				} else if os.Args[2] == "-" {
					printNbr(trimAtoi(os.Args[1]) - trimAtoi(os.Args[3]))
				} else if os.Args[2] == "*" {
					printNbr(trimAtoi(os.Args[1]) * trimAtoi(os.Args[3]))
				} else if os.Args[3] == "0" {
					os.Stdout.Write([]byte("No division by 0 | no modulo by 0"))
				} else {
					printNbr(trimAtoi(os.Args[1]) / trimAtoi(os.Args[3]))
				}
				printRunes('\n')
			}
		}
	}
}
