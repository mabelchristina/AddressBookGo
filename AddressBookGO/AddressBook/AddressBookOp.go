package addressbook

import (
	"fmt"
)

var AddressArray []Contact


func CreateContact() {
	var length int
	fmt.Println("Enter the name of contact add")
	fmt.Scanln(&length)
for a,_ := range length()
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

func DisplayContact() {
	fmt.Printf("%v", AddressArray)
}

func UpdateContact() {
	var name string
	fmt.Println("Enter the name of person to edit contact")
	fmt.Scanln(&name)

	for i := range AddressArray {
		if AddressArray[i].FirstName == name {
			fmt.Println("Choose the field to update:")
			var options int
			fmt.Println("1.LastName\n2.Email\n3.Phone")
			fmt.Scanln(&options)

			switch options {
			case 1:
				fmt.Println("Enter Last Name :")
				var lastName string
				fmt.Scanln(&lastName)
				AddressArray[i].LastName = lastName
			case 2:
				fmt.Println("Enter Email :")
				var email string
				fmt.Scanln(&email)
				AddressArray[i].Email = email

			case 3:
				fmt.Println("Enter Phone Number:")
				var phone string
				fmt.Scanln(&phone)
				AddressArray[i].Phone = phone

			}

		}
		fmt.Println("No person found")
	}
}

