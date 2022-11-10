package main

import (
	"fmt"
)

type Ville struct {
	nombreHabitant int
	codePostal     int
	superficie     int
}

func main() {

	var Villes = map[string]Ville{}

	fmt.Println("Enter the number of city : ")
	var nbOfCity int
	for i := 0; i < nbOfCity; i++ {

		fmt.Println("Enter infos for city number ", i+1)

		fmt.Printf("Enter the first city name: ")
		var cityName string
		fmt.Scanln(&cityName)

		fmt.Printf("Enter the number of citizens: ")
		var nombreHabitant int
		fmt.Scanln("%d", &nombreHabitant)

		fmt.Printf("Enter the postal code: ")
		var codePostal int
		fmt.Scanln("%d", &codePostal)

		fmt.Printf("Enter the superficy: ")
		var superficie int
		fmt.Scanln("%d", &superficie)

		Villes[cityName] = Ville{
			nombreHabitant: nombreHabitant,
			codePostal:     codePostal,
			superficie:     superficie,
		}
	}
	fmt.Println(Villes)
}
