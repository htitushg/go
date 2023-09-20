package main

import (
	"fmt"
	"strconv"
)

func main() {
	numStr := "123"
	numInt, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error occurred :", err)
		return
	}
	fmt.Println("Converted number :", numInt)
}
