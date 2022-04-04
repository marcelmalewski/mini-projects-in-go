package main

import "fmt"

//funkcja ktora sprawdza czy slice posiada dany element
func contains(elements []int, element int) bool {
	for _, value := range elements {
		if element == value {
			return true
		}
	}
	return false
}

//z malej czyli prywatna funkcja
func Zadanie1() {
	//slice z wybranymi indexami
	chosenElements := make([]int, 0)
	var option, chosenIndex int
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Wybierz ktore miejsca usunąć")
Loop:
	//menu
	for {
		fmt.Println("1 - dalej wybieram.")
		fmt.Println("2 - usuwam.")
		fmt.Scanf("%d\n", &option)

		switch option {
		case 1:
			//kontunuje wybieranie
			fmt.Println("Choose index:")
			fmt.Scanf("%d\n", &chosenIndex)
			if chosenIndex < len(slice1) {
				chosenElements = append(chosenElements, chosenIndex)
			}
		case 2:
			//wychodze z pętli
			break Loop
		}
	}

	//tmp to slice1 po filtrowaniu
	tmp := make([]int, 0)
	for i, v := range slice1 {
		//sprawdzamy czy index jest w wybranych indexach
		if contains(chosenElements, i) == false {
			tmp = append(tmp, v)
		}
	}
	slice1 = tmp
	fmt.Println(slice1)
	fmt.Println(chosenElements)
}
