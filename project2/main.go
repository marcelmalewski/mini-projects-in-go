package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

type Player struct {
	points     int
	rejections int
}

func main() {
	var menuOption string
	//menu
menu:
	for {
		fmt.Println("Options: \n 1 - start new game\n 2 - exit")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case "1":
			//one game
			winner := OneGame()
			//check if you won
			//if you won save your result to csv file
			saveResultToCSV(winner, "results.csv")
		case "2":
			break menu
		}
	}
}

func saveResultToCSV(winner string, filename string) {
	var yourName string

	if winner == "you" {
		fmt.Println("You won, so What is your name?")
		fmt.Scanln(&yourName)

		//all records from csv
		records := getDataFromCSV(yourName, filename)

		//send updated records to csv
		sendNewDataToCSV(records, filename)
	} else {
		//if you didnt win then nothing will happen
		fmt.Println("You lost so we dont need your name.")
	}
}

func getDataFromCSV(yourName, filename string) (records [][]string) {
	// open file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(file)

	//check if you are already in results or not
	areYouInResults := false
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		//if you are then add one point to your nickname
		if record[0] == yourName {
			areYouInResults = true

			//get your points as int ant add one
			currentScore, _ := strconv.Atoi(record[1])
			newScore := currentScore + 1
			//change points to string
			newScoreString := strconv.Itoa(newScore)
			updatedRecord := []string{yourName, newScoreString}
			//append to all recaords changed record with your nickname
			records = append(records, updatedRecord)
		} else {
			//else we append unchanged record
			records = append(records, record)
		}
	}

	//if your record wasnt found add new one with 1 point
	if areYouInResults == false {
		newRecord := []string{yourName, "1"}
		records = append(records, newRecord)
	}

	//close file
	file.Close()
	fmt.Println("Now your point is saved")
	return
}

func sendNewDataToCSV(records [][]string, filename string) {
	//open file with read and write only to change it
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := csv.NewWriter(file)

	//add all records to csv
	err = csvWriter.WriteAll(records)

	if err != nil {
		log.Fatal(err)
	}

	//close file
	file.Close()
}

func OneGame() (winner string) {
	player := Player{0, 0}
	bot := Player{0, 0}
	var youDrawOrNotDrawCard, howManyPointsFromAs string

	//card on the top of deck
	topOfTheDeckOfCards := rand.Intn(10) + 2

	for {
		//Check both cant move anymore
		if (player.rejections == 3 || player.points > 21) && (bot.rejections == 3 || bot.points > 21) {
			fmt.Printf("Number of your points: %d\n", player.points)
			fmt.Printf("Number of bot's points: %d\n", bot.points)
			//my ditance to 21
			playerDistance := math.Abs(float64(21 - player.points))
			//bot's distance to 21
			botDistance := math.Abs(float64(21 - bot.points))

			if playerDistance < botDistance {
				fmt.Println("You Won!!!")
				return "you"
			} else if playerDistance > botDistance {
				fmt.Println("Bot Won!!!")
				return "bot"
			} else {
				fmt.Println("Draw!!!")
				return "draw"
			}
		}

		//check if person can do sth
		if player.rejections != 3 && player.points < 21 {
			fmt.Println("---------------------------")
			//options
			fmt.Println("Options: \n 1 - draw a card\n 2 - do not draw a card")
			fmt.Scanln(&youDrawOrNotDrawCard)

			switch youDrawOrNotDrawCard {
			case "1":
				//what if player draw an as
				if topOfTheDeckOfCards == 11 {
					fmt.Println("Your card is As")
					fmt.Println("Options: \n 1 - As = 1 point\n 2 - As = 11 points")
					fmt.Scanln(&howManyPointsFromAs)
					switch howManyPointsFromAs {
					case "1":
						topOfTheDeckOfCards = 1
					case "2":
						topOfTheDeckOfCards = 11
					}

				}
				//add points
				player.points += topOfTheDeckOfCards
				fmt.Printf("Value of drawn card: %d\n", topOfTheDeckOfCards)
				fmt.Printf("Number of your points: %d\n", player.points)
				//if player have 21 points he wins
				if player.points == 21 {
					return "you"
				}
				//get new card on the top of deck
				topOfTheDeckOfCards = rand.Intn(10) + 2
			case "2":
				//rejection
				player.rejections++
				fmt.Printf("Number of rejections: %d\n", player.rejections)
				fmt.Printf("Number of your points: %d\n", player.points)
				//check if player rejected 3 times already
				if player.rejections == 3 {
					fmt.Printf("Your final number of points: %d\n", player.points)
				}
			default:
				fmt.Println("Wrong youDrawOrNotDrawCard")
			}
		}

		//check if bot can still play
		if bot.rejections != 3 && bot.points < 21 {
			//bot have 25% to draw card
			botDrawOrNotDrawCard := rand.Intn(4)

			if botDrawOrNotDrawCard > 0 {
				//add points
				bot.points += topOfTheDeckOfCards
				fmt.Printf("Value of drawn card by bot: %d\n", topOfTheDeckOfCards)
				fmt.Printf("Number of bot points: %d\n", bot.points)

				if bot.points == 21 {
					return "bot"
				} else if bot.points >= 16 {
					// what if bot points are higher than 16
					randOption := rand.Intn(1)

					if randOption == 1 {
						bot.rejections = 3
					}
				}
				//new card on the top of the deck
				topOfTheDeckOfCards = rand.Intn(10) + 2
			} else {
				bot.rejections++
				fmt.Printf("Number of bot's rejections: %d\n", bot.rejections)
				fmt.Printf("Number of bot's points: %d\n", bot.points)
				if bot.rejections == 3 {
					fmt.Printf("Bot's final number of points: %d\n", bot.points)
				}
			}

		}

	}
}
