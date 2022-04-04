package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//Zadanie1()
	Zadanie2()
}

func analizeChosenNumber(chosenNumber, randomNumber int) bool {
	if chosenNumber == randomNumber {
		fmt.Println("You Won!")
		return true
	} else if chosenNumber < randomNumber {
		fmt.Println("random number is higher")
	} else if chosenNumber > randomNumber {
		fmt.Println("random number is lower")
	}

	return false
}

func game() int {
	numberOfAttempts := 0
	var chosenOption string
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)

	for {
		fmt.Println("Guess a number")
		//read option from user and check errors
		_, err := fmt.Scanf("%s\n", &chosenOption)
		if err != nil {
			fmt.Println("error: ", err)
		} else {
			numberOfAttempts += 1
		}

		//comend "end" end program
		if chosenOption == "end" {
			fmt.Println("the end")
			return 0
		} else {
			//get number from option
			optionAsNumber, err := strconv.Atoi(chosenOption)

			//check if chosen option is number
			if err != nil {
				fmt.Println("wrong option")
			} else {
				if analizeChosenNumber(optionAsNumber, randomNumber) {
					return numberOfAttempts
				}
			}
		}
	}
}

func saveResultToCSV(numberOfAttempts int, filename string) {
	var name string

	_, err := fmt.Scanf("%s\n", &name)
	if err != nil {
		fmt.Println("error: ", err)
	}

	//all records from csv
	records := getDataFromCSV(name, filename)

	//send updated records to csv
	sendNewDataToCSV(records, numberOfAttempts, filename)
}

func Zadanie2() {
	var menuOption, name string

Menu:
	for {
		_, err := fmt.Scanf("%s\n", &menuOption)
		if err != nil {
			fmt.Println("error: ", err)
		}

		switch menuOption {
		case "start a game":
			numberOfAttempts := game()

			fmt.Println(numberOfAttempts)
		case "quit":
			break Menu
		}
	}
}
