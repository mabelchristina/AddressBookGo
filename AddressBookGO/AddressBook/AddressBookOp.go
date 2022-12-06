package addressbook

import (
	"fmt"
)

var AddressArray []Contact

func CreateContact() {
	var firstname, lastname, email, phone string

	fmt.Println("Enter First name")
	fmt.Scan(&firstname)

	fmt.Println("Enter Last name")
	fmt.Scan(&lastname)

	fmt.Println("Enter email address")
	fmt.Scan(&email)

	fmt.Println("Enter phone number")
	fmt.Scan(&phone)

	Contact := New(firstname, lastname, email, phone)

	AddressArray = append(AddressArray, Contact)
}

func DisplayContact(){
	fmt.Printf("%v", AddressArray)
}


