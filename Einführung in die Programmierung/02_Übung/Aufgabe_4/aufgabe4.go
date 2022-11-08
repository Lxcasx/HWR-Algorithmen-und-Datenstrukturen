package main

import "fmt"

type User struct {
	username string
	password string
}

// Entwickeln Sie ein Programm, welches einen Usernamen und ein Passwd einliest und überprüft.
// Hinterlegen Sie hierzu mindestens 5 Usernamen und zugehörige Passwords.
// Verwenden Sie hierzu die Datentypen Struct und Array.
func main() {
	users := [5]User{
		{username: "Hans", password: "123"},
		{username: "Peter", password: "234"},
		{username: "Franz", password: "456"},
		{username: "Josef", password: "789"},
		{username: "Ingolf", password: "2423"},
	}

	var username, password string

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	for _, e := range users {
		if e.username == username && e.password == password {
			fmt.Println("Login erfolgreich")
		}
	}
}
