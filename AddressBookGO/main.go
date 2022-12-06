package main

import (
	"fmt"

	addressbook "package.com/addressbook/addressBook"
)

func main() {
	fmt.Print("Welcome to addressBook in go")
	addressbook.CreateContact()
	addressbook.DisplayContact()
	addressbook.UpdateContact()
}
