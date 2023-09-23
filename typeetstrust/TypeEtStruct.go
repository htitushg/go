package main

import "fmt"

func main() {

	myContact := newContact("Alex", 30, map[string]string{"Fixe": "10.14.25.36.25", "Portable": "17.54.14.25.62"})
	//fmt.Println(myContact)
	//fmt.Println(displayContact())
	fmt.Println(myContact.displayContact())

	myNewContact := newContact("Justine", 28, map[string]string{"Fixe": "10.14.25.44.32", "Portable": "56.33.14.25.62"})
	//fmt.Println(myNewContact)
	//fmt.Println(displayContact())
	fmt.Println(myNewContact.displayContact())

	myNewContact2 := newContact("Henry", 73, map[string]string{"Fixe": "05.56.25.44.32", "Portable": "06.64.14.25.62"})
	//fmt.Println(myNewContact2)
	//fmt.Println(displayContact())
	fmt.Println(myNewContact2.displayContact())
}
