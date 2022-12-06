package main

import (
	"fmt"

	addressbook "package.com/addressbook/addressBook"
)

func main() {
	fmt.Print("Welcome to addressBook in go")
	var option int
	fmt.Println("Enter 1 to add contact\n ")
	fmt.Println("Enter your option:")
	fmt.Scanln(&option)

	switch option {

	case 1:
		addressbook.CreateContact()
		addressbook.DisplayContact()
	default:
		fmt.Println("enter the valid option")
	}
}
