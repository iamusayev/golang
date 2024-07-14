package main

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	id        uint64
	firstname string
	lastname  string
	email     string
	age       uint
	created   time.Time
}

var users = make(map[uint64]User)
var orders = make([]uint64, 5, 5)

func main() {
	orders = append(orders, 1)
	var firstname string
	var lastname string
	var email string
	var age uint
	var number uint
	var id uint64
	for {
		fmt.Println("1.Save \n2.Get user by id \n3.Edit existing user\n4.Get all users")
		fmt.Scan(&number)

		switch number {
		case 1:
			firstname, lastname, email, age = getUserInput(firstname, lastname, email, age)
			isValidFirstnameAndLastname, isValidEmail, isValidAge := validateUserInput(firstname, lastname, email, age, users)
			validateAndSave(isValidFirstnameAndLastname, isValidEmail, isValidAge, firstname, lastname, email, age)
		case 2:
			fmt.Println("Input id\n")
			fmt.Scan(&id)
			getUserById(id)
		case 3:
			fmt.Println("Input id\n")
			fmt.Scan(&id)
			firstname, lastname, email, age := getUserInput(firstname, lastname, email, age)
			updateUserById(id, firstname, lastname, email, age)
		case 4:
			fmt.Printf("\n%v\n", users)
		}

	}
}

func validateAndSave(isValidFirstnameAndLastname bool, isValidEmail bool, isValidAge bool, firstname string, lastname string, email string, age uint) {
	if isValidFirstnameAndLastname && isValidEmail && isValidAge {
		save(firstname, lastname, email, age)
		fmt.Printf("Your info saved\n")
	} else {
		fmt.Println("Your input is not correct, try again")
	}
}

func updateUserById(id uint64, firstname string, lastname string, email string, age uint) {

	selectedUser := users[id]
	if selectedUser.id == 0 {
		fmt.Printf("User with id %v not found, try it again", id)
	} else {
		users[id] = User{
			id:        id,
			firstname: firstname,
			lastname:  lastname,
			email:     email,
			age:       age,
			created:   time.Now(),
		}
		fmt.Printf("User with id %v has been updated", id)
	}
}

func getUserById(id uint64) {
	selectedUser := users[id]
	if selectedUser.id == 0 {
		fmt.Printf("User with id %v not found, try it again", id)
	} else {
		fmt.Println("#####################################################\n")
		fmt.Printf("Selected user is %v", selectedUser)
		fmt.Println("#####################################################\n")
		return
	}
}

func save(firstname string, lastname string, email string, age uint) {
	lastID := orders[len(orders)-1]
	lastID++
	orders = append(orders, lastID)
	users[lastID] = User{
		id:        lastID,
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		age:       age,
		created:   time.Now(),
	}
}

func getUserInput(firstname string, lastname string, email string, age uint) (string, string, string, uint) {
	fmt.Println("enter your first name")
	fmt.Scan(&firstname)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastname)

	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	fmt.Println("Enter your age")
	fmt.Scan(&age)
	return firstname, lastname, email, age
}

func welcome() {
	fmt.Println("#####################################################")

	fmt.Println("Welcome to our application, please register account")

	fmt.Println("#####################################################")

}

func validateUserInput(firstName string, lastName string, email string, age uint, users map[uint64]User) (bool, bool, bool) {
	isValidFirstNameAndLastName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && !checkEmail(users, email)
	isValidAge := age > 16 && age < 90
	return isValidFirstNameAndLastName, isValidEmail, isValidAge
}

func checkEmail(users map[uint64]User, email string) bool {
	for _, user := range users {
		if user.email == email {
			return true
		}
	}
	return false
}
