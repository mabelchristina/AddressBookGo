package addressbook

import (
	"fmt"
)

var AddressArray []Contact
var contactMap map[string]Contact

func AddressBookOperation() {
	var input int
	for ok := true; ok; ok = (input != 2) {
		var choice int
		fmt.Println(" Enter 1 for create contact \n Enter 2  for display contact \n Enter 3 for updating contact\n Enter 4 for deleting contact\n enter 5 for entering multi contact\nenter 6 for exiting loop")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			CreateContact()
		case 2:
			DisplayContact()
		case 3:
			UpdateContact()
		case 4:
			DeleteContact()
		case 5:
			MultiContact()
		case 6:
			fmt.Println("Exting the loop")
			input = 2
		}

	}

}

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
			DisplayContact()
		}
	}
}

func DeleteContact() {
	var name string
	fmt.Println("Enter the name of person to delete")
	fmt.Scanln(&name)

	for i := range AddressArray {
		if AddressArray[i].FirstName == name {
			AddressArray = nil
			fmt.Println("Record deleted contact is", AddressArray)
		}
	}
}
func MultiContact() {

	var num int
	fmt.Println("Enter No of contacts you need to add")
	fmt.Scanln(&num)
	for num > 0 {
		CreateContact()
		num--
	}

}
