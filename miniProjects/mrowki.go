package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type object struct {
	name        string
	value       string
	coordinates [2]int
}

func newObject(name string, value string, coordinates [2]int) *object {
	object := object{name: name}
	object.value = value
	object.coordinates = coordinates
	//losuje sie czy jest warriarem czy nie
	return &object
}

type objectWithNewCoordinates struct {
	object            *object
	coordinatesOfMove [2]int
}

func printAnthill(anthill [30][30]*object) {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}

	for _, row := range anthill {
		for _, space := range row {
			fmt.Print(space.value)
		}
		fmt.Println()
	}
}

func createAnthill() (anthill [30][30]*object) {
	//anthill gdzie wszędzie jest trawa
	var row [30]*object

	for i := range anthill {
		for j := range row {
			row[j] = &object{"grass", "-", [2]int{i, j}}
		}
		anthill[i] = row
	}

	return anthill
}

func checkIfSpaceIsEmpty(coordinates [2]int, anthill *[30][30]*object) bool {
	if coordinates[0] < 0 || coordinates[0] >= 30 || coordinates[1] < 0 || coordinates[1] >= 30 {
		return false
	}

	if anthill[coordinates[0]][coordinates[1]].value == "-" {
		return true
	} else {
		return false
	}
}

func putLeavesOnAnthill(anthill *[30][30]*object, howMany int) {

	numberOfLeavesPutOnAnthill := 0
loop:
	for {
		rand.Seed(time.Now().UnixNano())
		row := rand.Intn(30)
		col := rand.Intn(30)

		if checkIfSpaceIsEmpty([2]int{row, col}, anthill) {
			// spawn
			anthill[row][col] = &object{"leaf", "l", [2]int{row, col}}
			numberOfLeavesPutOnAnthill++
		}

		if numberOfLeavesPutOnAnthill == howMany {
			break loop
		}
	}
}

func generateRandomNumber(seedsChannel chan int64) int {
	newSeed := <-seedsChannel

	rand.Seed(newSeed)
	return rand.Intn(4)
}

func generateRandomCoordinates() [2]int {
	rand.Seed(time.Now().UnixNano())

	row := rand.Intn(30)
	col := rand.Intn(30)

	return [2]int{row, col}
}

func lifeOfAnAnt(ant *object, seedsChannel chan int64, antsChannel chan objectWithNewCoordinates, anthill *[30][30]*object) {
	for {
		time.Sleep(100 * time.Millisecond)
		// czy możemy wysłać ruch
		if len(antsChannel) < 30 {
			// czy nie jestem jeszcze na mapie
			if ant.coordinates == [2]int{-1, -1} {
				spawnOption := generateRandomCoordinates()
				antWithMoveOption := objectWithNewCoordinates{ant, spawnOption}

				if checkIfSpaceIsEmpty(spawnOption, anthill) {

					antsChannel <- antWithMoveOption
				}
			} else {
				//// jestem juz na mapie i sie probuje ruszyć
				moveOption := generateRandomNumber(seedsChannel)
				switch moveOption {
				case 0:
					spawnOption := [2]int{ant.coordinates[0] - 1, ant.coordinates[1]}
					antWithMoveOption := objectWithNewCoordinates{ant, spawnOption}

					if checkIfSpaceIsEmpty(spawnOption, anthill) {
						antsChannel <- antWithMoveOption
					}
				case 1:
					spawnOption := [2]int{ant.coordinates[0], ant.coordinates[1] + 1}
					antWithMoveOption := objectWithNewCoordinates{ant, spawnOption}

					if checkIfSpaceIsEmpty(spawnOption, anthill) {
						antsChannel <- antWithMoveOption
					}
				case 2:
					spawnOption := [2]int{ant.coordinates[0] + 1, ant.coordinates[1]}
					antWithMoveOption := objectWithNewCoordinates{ant, spawnOption}

					if checkIfSpaceIsEmpty(spawnOption, anthill) {
						antsChannel <- antWithMoveOption
					}
				case 3:
					spawnOption := [2]int{ant.coordinates[0], ant.coordinates[1] - 1}
					antWithMoveOption := objectWithNewCoordinates{ant, spawnOption}

					if checkIfSpaceIsEmpty(spawnOption, anthill) {
						antsChannel <- antWithMoveOption
					}
				}

			}
		}
	}
}

func checkIfLeafIsNextToNewCoordinates(coordinates [2]int, anthill *[30][30]*object) bool {
	row := coordinates[0]
	col := coordinates[1]

	if col+1 < 30 {
		if anthill[row][col+1].value == "l" {
			return true
		}
	}

	if col-1 >= 0 {
		if anthill[row][col-1].value == "l" {
			return true
		}
	}

	if row+1 < 30 {
		if anthill[row+1][col].value == "l" {
			return true
		}
	}

	if row-1 >= 0 {
		if anthill[row-1][col].value == "l" {
			return true
		}
	}

	return false
}

func getCoordinatesOfLeafNextToAnt(coordinates [2]int, anthill *[30][30]*object) [2]int {
	row := coordinates[0]
	col := coordinates[1]

	if col+1 < 30 {
		if anthill[row][col+1].value == "l" {
			return [2]int{row, col + 1}
		}
	}

	if col-1 >= 0 {
		if anthill[row][col-1].value == "l" {
			return [2]int{row, col - 1}
		}
	}

	if row+1 < 30 {
		if anthill[row+1][col].value == "l" {
			return [2]int{row + 1, col}
		}
	}

	if row-1 >= 0 {
		if anthill[row-1][col].value == "l" {
			return [2]int{row - 1, col}
		}
	}

	return [2]int{-1, -1}
}

func putObjectOnAnthillOnCoordinates(antWithMoveOption objectWithNewCoordinates, anthill *[30][30]*object) {
	prevRow := antWithMoveOption.object.coordinates[0]
	prevCol := antWithMoveOption.object.coordinates[1]

	newRow := antWithMoveOption.coordinatesOfMove[0]
	newCol := antWithMoveOption.coordinatesOfMove[1]

	if prevRow == -1 {
		// spawn
		// dajemy jej nowe koordynany
		antWithMoveOption.object.coordinates = [2]int{newRow, newCol}

		//a tu sie pojawia
		anthill[newRow][newCol] = antWithMoveOption.object

	} else {
		if checkIfLeafIsNextToNewCoordinates(antWithMoveOption.coordinatesOfMove, anthill) {
			if antWithMoveOption.object.value == "m" {
				coordinatesOfLeaf := getCoordinatesOfLeafNextToAnt(antWithMoveOption.coordinatesOfMove, anthill)
				if coordinatesOfLeaf[0] != -1 {
					//bez liscia
					//czyli go podnosi
					antWithMoveOption.object.value = "e"

					//mrowka ma nowe koordyanty
					antWithMoveOption.object.coordinates = [2]int{newRow, newCol}

					//tam gdzie byla trawa teraz jest mrowka
					anthill[newRow][newCol] = antWithMoveOption.object

					//tam gzie stala mrowka jest teraz trawa
					anthill[prevRow][prevCol] = &object{"grass", "-", [2]int{prevRow, prevCol}}

					//znaleziony lisc znika
					anthill[coordinatesOfLeaf[0]][coordinatesOfLeaf[1]] = &object{"grass", "-", [2]int{coordinatesOfLeaf[0], coordinatesOfLeaf[1]}}
				}
			} else {
				//niesie juz liscia i go zostawia na pole ruchu
				antWithMoveOption.object.value = "m"

				//tam gdzie byla trawa teraz jest lisc
				anthill[newRow][newCol] = &object{"leaf", "l", [2]int{newRow, newCol}}
			}
		} else {
			//move
			antWithMoveOption.object.coordinates = [2]int{newRow, newCol}

			//tam gzie stala mrowka jest teraz trawa
			anthill[prevRow][prevCol] = &object{"grass", "-", [2]int{prevRow, prevCol}}

			//tam gdzie byla trawa teraz jest mrowka
			anthill[newRow][newCol] = antWithMoveOption.object
		}
	}
}

func anthillController(anthill *[30][30]*object, antsChannel chan objectWithNewCoordinates) {
	for {
		if len(antsChannel) != 0 {
			antWithMoveOption := <-antsChannel

			//pytanie czy te miejsce dalej jest wolne
			if checkIfSpaceIsEmpty(antWithMoveOption.coordinatesOfMove, anthill) { //prezsylam wartosc
				putObjectOnAnthillOnCoordinates(antWithMoveOption, anthill) // przesylam pointer
			}
		}
	}
}

func anthillShow(anthill *[30][30]*object) {
	for {
		printAnthill(*anthill)
	}
}

func makeSeeds(seedsChannel chan int64) {
	var number int64 = 0
	for {
		if len(seedsChannel) < 30 {
			seedsChannel <- number
			number++
		}
	}
}

func main() {
	seedsChannel := make(chan int64, 30)
	go makeSeeds(seedsChannel)

	antsChannel := make(chan objectWithNewCoordinates, 30)

	anthill := createAnthill()
	putLeavesOnAnthill(&anthill, 70)

	// koordyna (-1, -1) oznaczają ze mrowka nie jest na planszy
	for i := 0; i < 40; i++ {
		go lifeOfAnAnt(newObject("ant", "m", [2]int{-1, -1}), seedsChannel, antsChannel, &anthill)
	}

	go anthillShow(&anthill)

	anthillController(&anthill, antsChannel)
}
