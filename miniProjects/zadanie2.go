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
	date             string
	//date and zgadnieta liczba
}

type ByNumberOfAttempts []playerWithResult

// rzeczy potrzebne do sortowania do uzycia sort
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
	fmt.Println(randomNumber)

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
	//open file with write only to change it
	file, err := os.OpenFile(filename, os.O_WRONLY, 0222)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(file)

	//set first line in csv file
	firstLine := []string{"number", "name", "numberOfAttempts", "date"}

	err = csvWriter.Write(firstLine)
	if err != nil {
		fmt.Println(err)
	}

	//add all records to csv
	for index, result := range dataBase {
		resultAsArray := []string{strconv.Itoa(index + 1), result.name, strconv.Itoa(result.numberOfAttempts), result.date}

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

func createDataBaseFromData(data [][]string) []playerWithResult {
	dataBase := make([]playerWithResult, 0)

	for index, row := range data {
		if index > 0 {
			numberOfAttempts, _ := strconv.Atoi(row[2])
			dataBase = append(dataBase, playerWithResult{row[1], numberOfAttempts, row[3]})
		}
	}

	return dataBase
}

func getDataBaseFromFile(filename string) []playerWithResult {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return createDataBaseFromData(data)
}

func printFormattedResults(dataBase []playerWithResult) {
	fmt.Println("Results:")
	for index, result := range dataBase {
		fmt.Printf("%d. name: %s, number of attempts: %d, date: %s\n", index+1, result.name, result.numberOfAttempts, result.date)
	}
}

func updateYourResultInDataBase(dataBase []playerWithResult, name string, numberOfAttempts int, date string) bool {
	updateWasMade := false
	var indexOfUpdatedResult int

	for index, result := range dataBase {
		if result.name == name {
			if result.numberOfAttempts > numberOfAttempts {
				updateWasMade = true
				indexOfUpdatedResult = index
			}
		}
	}

	if updateWasMade {
		dataBase[indexOfUpdatedResult] = playerWithResult{name, numberOfAttempts, date}
		return true
	}
	return false

}

func addResultToDataBase(dataBase *[]playerWithResult, name string, numberOfAttempts int, date string) {
	//zwraca falsz gdy nie uda sie z updatowac co oznacza ze nie ma go i trzeba dodac
	if !updateYourResultInDataBase(*dataBase, name, numberOfAttempts, date) {
		//add result to dataBase
		*dataBase = append(*dataBase, playerWithResult{name, numberOfAttempts, date})
	}
}

func menuAndGameControl() {
	var name string
	var menuOption int

	//pobieramy rezultaty z pliku
	dataBase := getDataBaseFromFile("results.csv")

menu:
	for {

		fmt.Println("Menu:")
		fmt.Println("1. Start a game")
		fmt.Println("2. Quit a game")

		//kontrola czy to napewno liczba
		_, err := fmt.Scanf("%d\n", &menuOption)
		if err != nil {
			fmt.Println("error: ", err)
		}

		switch menuOption {
		case 1:
			numberOfAttempts := game()

			fmt.Println("What is your name ?")
			_, err := fmt.Scanf("%s\n", &name)

			if err != nil {
				fmt.Println("error: ", err)
			}
			currentTime := time.Now().Format("01-02-2006 15:04:05")

			addResultToDataBase(&dataBase, name, numberOfAttempts, currentTime)
		case 2:
			//sort results
			sort.Sort(ByNumberOfAttempts(dataBase))
			printFormattedResults(dataBase)
			saveResultsToCSV(dataBase, "results.csv")
			break menu
		}
	}
}

func zadanie2() {
	menuAndGameControl()
}
