package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Printf("Combien d'éléments contient votre liste ? ")
	var nbElements int
	fmt.Scanln(&nbElements)

	var liste []string

	for i := 0; i < nbElements; i++ {
		fmt.Println("Entrez l'élément", i+1, "de votre liste : ")
		var element string
		fmt.Scanln(&element)

		liste = append(liste, element)
	}
	sort.Strings(liste)
	fmt.Println(liste)
}
