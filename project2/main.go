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

var menuOption, yourName string

//deklarowanie czy make czy bez make , slice ?
//doczytanie o flagach
//exit nie dziala

func main() {
	for {
		//menu
		fmt.Println("Options: \n 1 - start new game\n 2 - exit")
		fmt.Scanln(&menuOption)

		switch menuOption {
		case "1":
			//dzieje sie gra
			winner := OneGame()
			//jezeli wygrales podajesz imie by dodac do tego imienia punkt
			//albo dodac imie do wynikow
			if winner == "you" {
				fmt.Println("You won, so What is your name?")
				fmt.Scanln(&yourName)

				//zczytujemy rekory z csv wynikow
				//records = nil
				var records [][]string

				file, err := os.Open("results.csv")
				if err != nil {
					log.Fatal(err)
				}
				csvReader := csv.NewReader(file)

				areYouInResults := false
				for {
					record, err := csvReader.Read()
					if err == io.EOF {
						break
					}
					if err != nil {
						log.Fatal(err)
					}

					//jezlie znajdziemy imie to je updejtujemy o jeden
					if record[0] == yourName {
						areYouInResults = true
						currentScore, _ := strconv.Atoi(record[1])
						newScore := currentScore + 1
						newScoreString := strconv.Itoa(newScore)

						updatedRecord := []string{yourName, newScoreString}
						records = append(records, updatedRecord)
					} else {
						records = append(records, record)
					}
				}

				//jezeli nie to dodajemy nowy rekord
				if areYouInResults == false {
					newRecord := []string{yourName, "1"}
					records = append(records, newRecord)
				}

				file.Close()

				//read and write only, permisie maja wplyw tylko jak tworzymy plik
				file, err = os.OpenFile("results.csv", os.O_RDWR, 0644)
				if err != nil {
					log.Fatal(err)
				}
				csvWriter := csv.NewWriter(file)
				err = csvWriter.WriteAll(records)

				if err != nil {
					log.Fatal(err)
				}

				file.Close()

				fmt.Println("Now your point is saved")
			} else {
				fmt.Println("You lost so we dont need your name.")
			}
		case "2":
			break
		}
	}
}

func OneGame() (winner string) {
	var yourPoints, yourRejections, botPoints, botRejections, botDrawOrNotDrawCard int
	var youDrawOrNotDrawCard, howManyPointsFromAs string

	//karta na wierzchu decku
	topOfTheDeckOfCards := rand.Intn(10) + 2

	for {
		//sprawdzamy czy osoba moze wykonac ruch
		if yourRejections != 3 && yourPoints < 21 {
			fmt.Println("---------------------------")
			//opcje
			fmt.Println("Options: \n 1 - draw a card\n 2 - do not draw a card")
			fmt.Scanln(&youDrawOrNotDrawCard)

			switch youDrawOrNotDrawCard {
			case "1":
				//co gdy wyloswal 11 wybiera ile punktow za asa chcemy
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
				//dodawanie punktow
				yourPoints += topOfTheDeckOfCards
				fmt.Printf("Value of drawn card: %d\n", topOfTheDeckOfCards)
				fmt.Printf("Number of your points: %d\n", yourPoints)
				//jak 21 to od razu wygrana
				if yourPoints == 21 {
					winner = "you"
				}
				//losujemy nową karte na top decku
				topOfTheDeckOfCards = rand.Intn(10) + 2
			case "2":
				//odrzucamy dobieranie
				yourRejections++
				fmt.Printf("Number of rejections: %d\n", yourRejections)
				fmt.Printf("Number of your points: %d\n", yourPoints)
				//3 odrzucenia to koniec dobierania ogolnie
				if yourRejections == 3 {
					fmt.Printf("Your final number of points: %d\n", yourPoints)
				}
			default:
				fmt.Println("Wrong youDrawOrNotDrawCard")
			}
		}

		//po dobieraniu sprawdzamy czy gracz wygral
		if winner == "you" {
			println("Your Won!!!")
			return
		}

		//sprawdzamy czy bot dalej gra
		if botRejections != 3 && botPoints < 21 {
			//w poleceniu jest 0 lub 1
			//25% szans ze bot nie dobierze karty
			botDrawOrNotDrawCard = rand.Intn(4)

			if botDrawOrNotDrawCard > 0 {
				botPoints += topOfTheDeckOfCards
				fmt.Printf("Value of drawn card by bot: %d\n", topOfTheDeckOfCards)
				fmt.Printf("Number of bot points: %d\n", botPoints)

				if botPoints == 21 {
					winner = "bot"
				} else if botPoints >= 16 {
					// co gdy bot zblizy sie do 16 punktow
					//losuje czy przestaje dobierac czy nie
					randOption := rand.Intn(1)

					if randOption == 1 {
						botRejections = 3
					}
				}
				//losujemy nową karte na top decku
				topOfTheDeckOfCards = rand.Intn(10) + 2
			} else {
				botRejections++
				fmt.Printf("Number of bot's rejections: %d\n", botRejections)
				fmt.Printf("Number of bot's points: %d\n", botPoints)
				if botRejections == 3 {
					fmt.Printf("Bot's final number of points: %d\n", botPoints)
				}
			}

		}

		//po dobieraniu sprawdamy czy bot wygral
		if winner == "bot" {
			println("Bot Won!!!")
			return
		}
		//sprawdzamy czy gra się nie skonczyla
		//jezlie obaj mają 3 odmowy albo ponad 21 TO SPRAWDZAMY KTO WYGRAL
		if (yourRejections == 3 || yourPoints > 21) && (botRejections == 3 || botPoints > 21) {
			fmt.Printf("Number of your points: %d\n", yourPoints)
			fmt.Printf("Number of bot's points: %d\n", botPoints)
			//moja odlegosc od 21
			yourDistance := math.Abs(float64(21 - yourPoints))
			//bota odlegosc od 21
			botDistance := math.Abs(float64(21 - botPoints))

			if yourDistance < botDistance {
				fmt.Println("You Won!!!")
				return "you"
			} else if yourDistance > botDistance {
				fmt.Println("Bot Won!!!")
				return "bot"
			} else {
				fmt.Println("Draw!!!")
				return "draw"
			}
		}
	}
}
