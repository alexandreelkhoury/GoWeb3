package main

import (
	"fmt"
	//"github.com/joho/godotenv"
)

func main() {

	type Ville struct {
		nombreHabitant int
		codePostal     int
		superficie     int
	}

	var Villes map[string]Ville

	i := 1
	for i < 3 {

		fmt.Println("Enter the first city name: ")
		var cityName string
		fmt.Scanln(&cityName)

		fmt.Println("Enter the number of citizens: ")
		var nombreHabitant int
		fmt.Scanln(&nombreHabitant)

		fmt.Println("Enter the postal code: ")
		var codePostal int
		fmt.Scanln(&codePostal)

		fmt.Println("Enter the superficy: ")
		var superficie int
		fmt.Scanln(&superficie)

		ville := Ville{nombreHabitant, codePostal, superficie}
		fmt.Println(ville)

		Villes = make(map[string]Ville)
		Villes[cityName] = ville
		fmt.Println(Villes[cityName])

		i += i
	}

	fmt.Println(i)
	fmt.Println(Villes)

}
