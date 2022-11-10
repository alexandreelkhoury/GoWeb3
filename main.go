package main

import (
	"fmt"
	//"github.com/joho/godotenv"
)

func main() {

	fmt.Println("Votre nombre de transactions est par seconde ? Tapez 'seconde' \nVotre nombre de transactions est par minute ? Tapez 'minute'\nVotre nombre de transactions est par heure ? Tapez 'heure'")
	var temps string
	fmt.Scanln(&temps)

	fmt.Println("Quel est votre nombre de transactins ? ")
	var nbTransactions int
	fmt.Scanln(&nbTransactions)

	var tempsEnSeconde int
	var tempsEnMinute int
	var tempsEnHeure int

	switch temps {
	case "seconde":
		tempsEnSeconde = nbTransactions
		tempsEnMinute = nbTransactions * 60
		tempsEnHeure = nbTransactions * 3600
		fmt.Println("Transactions par seconde : ", tempsEnSeconde, "\n Transactions par minute : ", tempsEnMinute, "\nTransactions par heure : ", tempsEnHeure)
	case "minute":
		tempsEnSeconde = nbTransactions / 60
		tempsEnMinute = nbTransactions
		tempsEnHeure = nbTransactions * 60
		fmt.Println("Transactions par seconde : ", tempsEnSeconde, "\n Transactions par minute : ", tempsEnMinute, "\nTransactions par heure : ", tempsEnHeure)
	case "heure":
		tempsEnSeconde = nbTransactions / 3600
		tempsEnMinute = nbTransactions / 60
		tempsEnHeure = nbTransactions
		fmt.Println("Transactions par seconde : ", tempsEnSeconde, "\n Transactions par minute : ", tempsEnMinute, "\nTransactions par heure : ", tempsEnHeure)
	default:
		fmt.Printf("Veuillez ne pas entrer autre chose que seconde, minute et heure")
	}
}
