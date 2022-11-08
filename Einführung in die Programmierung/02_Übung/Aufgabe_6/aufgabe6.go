package main

import "fmt"

// Entwickeln Sie ein Programm, welches einen Usernamen und ein Passwd einliest und überprüft.
// Hinterlegen Sie hierzu mindestens 5 Usernamen und zugehörige Passwords.
// Verwenden Sie hierzu die Datentypen Struct und Array.
// Zusatz: Erweitern Sie das Programm, so dass es maximal 3 fehlerhafte Passwd‐Abfragen erlaubt und dann mit einer Fehlermeldung terminiert.
func main() {
	users := map[string]string{
		"Hans":   "123",
		"Peter":  "234",
		"Franz":  "456",
		"Josef":  "789",
		"Ingolf": "2423",
	}

	for i := 0; i < 3; i++ {
		var username, password string

		fmt.Printf("Enter username: ")
		fmt.Scanln(&username)
		fmt.Printf("Enter password: ")
		fmt.Scanln(&password)

		if users[username] == password {
			fmt.Println("Login erfolgreich")
			break
		} else {
			fmt.Println("Falsche Login Daten. Probiere es erneut.")
		}
	}
}
