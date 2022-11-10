package main

import (
	"fmt"
	"sort"
	//"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Combien d'éléments contient votre liste ? ")
	var nbElements int
	fmt.Scanln(&nbElements)

	var liste []string

	for i := 0; i < nbElements; i++ {
		fmt.Println("Entrez l'élément", i, " de votre liste : ")
		var element string
		fmt.Scanln("%s", &element)

		liste = append(liste, element)
	}

	sort.Strings(liste)
	fmt.Println(liste)

}
