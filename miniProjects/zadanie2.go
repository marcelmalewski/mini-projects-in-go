package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type playerWithResult struct {
	name             string
	numberOfAttempts int
	//date and zgadnieta liczba
}

type ByNumberOfAttempts []playerWithResult

func (a ByNumberOfAttempts) Len() int           { return len(a) }
func (a ByNumberOfAttempts) Less(i, j int) bool { return a[i].numberOfAttempts < a[j].numberOfAttempts }
func (a ByNumberOfAttempts) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func analyzeChosenNumber(chosenNumber, randomNumber int) bool {
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
				if analyzeChosenNumber(optionAsNumber, randomNumber) {
					return numberOfAttempts
				}
			}
		}
	}
}

func saveResultsToCSV(dataBase []playerWithResult, filename string) {
	//send updated records to csv
	//open file with read and write only to change it
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(file)

	//set first line in csv file
	firstLine := []string{"number", "name", "numberOfAttempts"}

	err = csvWriter.Write(firstLine)
	if err != nil {
		fmt.Println(err)
	}

	//add all records to csv
	for index, result := range dataBase {
		resultAsArray := []string{strconv.Itoa(index + 1), result.name, strconv.Itoa(result.numberOfAttempts)}

		err := csvWriter.Write(resultAsArray)
		if err != nil {
			fmt.Println(err)
		}

	}

	//flush writer
	csvWriter.Flush()

	//close file
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func printFormattedResults(dataBase []playerWithResult) {
	fmt.Println("Results:")
	for index, result := range dataBase {
		fmt.Printf("%d. name: %s, number of attempts: %d\n", index+1, result.name, result.numberOfAttempts)
	}
}

func zadanie2() {
	var name string
	var menuOption int
	dataBase := make([]playerWithResult, 0)

	dataBase = append(dataBase, playerWithResult{"janek", 15})
	dataBase = append(dataBase, playerWithResult{"kuba", 5})
	dataBase = append(dataBase, playerWithResult{"wojtek", 19})

Menu:
	for {

		fmt.Println("Menu:")
		fmt.Println("1. Start a game")
		fmt.Println("2. Quit a game")

		_, err := fmt.Scanf("%d\n", &menuOption)
		if err != nil {
			fmt.Println("error: ", err)
		}

		//przy eneter zaczyna nową gre

		switch menuOption {
		case 1:
			numberOfAttempts := game()

			fmt.Println("What is your name ?")
			_, err := fmt.Scanf("%s\n", &name)

			if err != nil {
				fmt.Println("error: ", err)
			}

			//first check if we have your name already in database
			//and check if new result is better
			//for _, result := range dataBase {
			//	if result.name == name{
			//		if result.numberOfAttempts > numberOfAttempts:
			//
			//	}
			//}

			//add result to dataBase
			dataBase = append(dataBase, playerWithResult{name, numberOfAttempts})
		case 2:
			//sort results
			sort.Sort(ByNumberOfAttempts(dataBase))

			printFormattedResults(dataBase)
			saveResultsToCSV(dataBase, "results.csv")
			break Menu
		}
	}
}
