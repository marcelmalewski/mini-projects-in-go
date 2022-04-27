package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ant struct {
	id    int
	value string
}

func newAnt(id int) *ant {
	newAnt := ant{id: id}
	newAnt.value = "m"
	//losuje sie czy jest warriarem czy nie
	return &newAnt
}

func printAnthill(anthill [20][20]ant) {
	for i := range anthill {
		fmt.Println(anthill[i])
	}
}

func createAnthill() (anthill [20][20]ant) {
	var row [20]ant //assign

	for i := range anthill { //assign
		anthill[i] = row
	}

	return anthill
}

func checkIfSpaceIsEmpty(row int, col int, anthill [20][20]ant) bool {
	if anthill[row][col] == "-" {
		return true
	} else {
		return false
	}
}

func checkNumberOfThingsOnAnthill(anthill [20][20]ant, thing string) int {
	sum := 0

	for _, row := range anthill {
		for _, space := range row {
			if space == thing {
				sum++
			}
		}
	}

	return sum
}

func addSomeThingOnAnthill(anthill *[20][20]string, thing string, howMany int) {
	rand.Seed(time.Now().UnixNano())

loop:
	for {
		row := rand.Intn(20)
		col := rand.Intn(20)

		if checkIfSpaceIsEmpty(row, col, *anthill) {
			anthill[row][col] = thing
		}

		if checkNumberOfThingsOnAnthill(*anthill, thing) == howMany {
			break loop
		}
	}
}

func main() {
	antHill := createAnthill()

	addSomeThingOnAnthill(&antHill, "m", 20)
	//addSomeThingOnAnthill(&antHill, "l", 30)

	printAnthill(antHill)
}
