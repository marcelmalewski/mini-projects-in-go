package main

import (
	"fmt"
	"math/rand"
	"time"
)

func printAnthill(anthill [20][20]string) {
	for i := range anthill {
		fmt.Println(anthill[i])
	}
}

func createAnthill() (anthill [20][20]string) {
	var row = [...]string{"-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"} //assign

	for i := range anthill { //assign
		anthill[i] = row
	}

	return anthill
}

func checkIfSpaceIsEmpty(row int, col int, anthill [20][20]string) bool {
	if anthill[row][col] == "-" {
		return true
	} else {
		return false
	}
}

func checkNumberOfAntsOnAnthill(anthill [20][20]string) int {
	sum := 0

	for _, row := range anthill {
		for _, space := range row {
			if space == "1" {
				sum++
			}
		}
	}

	return sum
}

func checkNumberOfLeavesOnAnthill(anthill [20][20]string) int {
	sum := 0

	for _, row := range anthill {
		for _, space := range row {
			if space == "2" {
				sum++
			}
		}
	}

	return sum
}

func createAntsOnAnthill(anthill *[20][20]string) {
	rand.Seed(time.Now().UnixNano())

loop:
	for {
		row := rand.Intn(20)
		col := rand.Intn(20)

		if checkIfSpaceIsEmpty(row, col, *anthill) {
			anthill[row][col] = "1"
		}

		if checkNumberOfAntsOnAnthill(*anthill) == 20 {
			break loop
		}
	}
}

func createLeavesOnAnthill(anthill *[20][20]string) {
	//z tejk i tej wyzej mozna zrobic jenno funkcje ktora w parametrze przyjmuje co wstawia
	rand.Seed(time.Now().UnixNano())

loop:
	for {
		row := rand.Intn(20)
		col := rand.Intn(20)

		if checkIfSpaceIsEmpty(row, col, *anthill) {
			anthill[row][col] = "2"
		}

		if checkNumberOfLeavesOnAnthill(*anthill) == 20 {
			break loop
		}
	}
}

func main() {
	antHill := createAnthill()

	createAntsOnAnthill(&antHill)
	createLeavesOnAnthill(&antHill)

	printAnthill(antHill)
}
