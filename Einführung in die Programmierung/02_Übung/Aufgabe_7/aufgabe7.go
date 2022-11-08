package main

import "fmt"

// Entwickeln Sie eine Datenstruktur für das Programm Fahrradtouren.
// Diese soll 5 Hotels sowie die Entfernung zum jeweils nächsten Hotel speichern.
// Erweitern Sie das Programm, so dass es bei Eingabe eines Hotels die Entfernung/Fahrzeiten bis zum nächsten Hotel ausgibt.

func main() {
	hotels := map[string]float32{
		"Hotel_1": 2,
		"Hotel_2": 5,
		"Hotel_3": 22,
		"Hotel_4": 40,
		"Hotel_5": 20,
	}

	var hotel string
	fmt.Printf("Enter Hotel: ")
	fmt.Scanln(&hotel)

	if hotels[hotel] == 0 {
		fmt.Println("Das Hotel konnte nicht gefunden")
	}

	nexHotel, nextHotelRange := getNextHotel(hotels, hotel)

	if nextHotelRange == 0 {
		fmt.Println("Es gibt kein weiteres Hotel.")
		return
	}

	fmt.Printf("Das nächste Hotel: '%s' ist %v km entfernt", nexHotel, nextHotelRange)
}

func getNextHotel(hotels map[string]float32, hotel string) (string, float32) {
	found := false

	for i, e := range hotels {
		if i == hotel {
			fmt.Println(i)
			found = true
			continue
		}

		if found {
			return i, e
		}
	}

	return "", 0
}
