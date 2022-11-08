package main

import "fmt"

type User struct {
	username string
	password string
}

// Entwickeln Sie ein Programm, welches einen Usernamen und ein Passwd einliest und überprüft.
// Hinterlegen Sie hierzu mindestens 5 Usernamen und zugehörige Passwords.
// Verwenden Sie hierzu die Datentypen Struct und Array.
// Zusatz: Verwenden Sie hierfür den Datentyp Map.
func main() {
	users := map[string]string{
		"Hans":   "123",
		"Peter":  "234",
		"Franz":  "456",
		"Josef":  "789",
		"Ingolf": "2423",
	}

	var username, password string

	fmt.Printf("Enter username: ")
	fmt.Scanln(&username)
	fmt.Printf("Enter password: ")
	fmt.Scanln(&password)

	if users[username] == password {
		fmt.Println("Login erfolgreich")
	}
}
