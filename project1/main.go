package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

// Swap and Less and Len we need this to sort results by concurrency
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Len() int           { return len(p) }

func main() {
	words := getWordsFromFile("ogniem_i_mieczem.txt")

	clearedWordsWithOccurrenceForMarcel, clearedWordsWithOccurrenceForPaulina := getClearedWordsWithOccurrence(words)
	sortedResultForMarcelToPrint, sortedResultForPaulinaToPrint := getSortedResultToPrintForMarcelAndPaulina(clearedWordsWithOccurrenceForMarcel, clearedWordsWithOccurrenceForPaulina)

	//makeshift window with results
	printResultsInWindow(sortedResultForMarcelToPrint, sortedResultForPaulinaToPrint)
}

func printResultsInWindow(resultForMarcelSlice, resultForPaulinaSlice []string) {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Widget")
	myWindow.Resize(fyne.NewSize(800, 700))

	titleMarcel := canvas.NewText("Zestaw dla Marcela", color.Black)
	titleMarcel.TextSize = 20

	titlePaulina := canvas.NewText("Zestaw dla Pauliny", color.Black)
	titlePaulina.TextSize = 20

	listMarcel := widget.NewList(
		func() int {
			return len(resultForMarcelSlice)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("yes")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(resultForMarcelSlice[i])
		})

	listPaulina := widget.NewList(
		func() int {
			return len(resultForMarcelSlice)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("yes")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(resultForPaulinaSlice[i])
		})

	vBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 300)), titleMarcel, listMarcel, titlePaulina, listPaulina)

	myWindow.SetContent(vBox)
	myWindow.ShowAndRun()
}

func getWordsFromFile(filename string) (slicedFile []string) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	//we change string of words to slice of words
	//and we make every letter small
	slicedFile = strings.Fields(strings.ToLower(string(content)))
	return
}

func getClearedWordsWithOccurrence(words []string) (clearedWordsWithOccurrenceForMarcel, clearedWordsWithOccurrenceForPaulina map[string]int) {
	//maps to count Occurrence for Marcel and Paulina
	clearedWordsWithOccurrenceForPaulina = make(map[string]int)
	clearedWordsWithOccurrenceForMarcel = make(map[string]int)

	//regex to clear every literal which is not a letter nor number
	reg, _ := regexp.Compile("[^a-zżźćńółęąś0-9]+")

	for _, word := range words {
		//clear word
		wordFiltered := reg.ReplaceAllLiteralString(word, "")

		//if word after clearing is not empty
		if wordFiltered != "" {
			//filter for Paulina
			if isAcceptedByPaulina(wordFiltered) {
				clearedWordsWithOccurrenceForPaulina[wordFiltered] += 1
			}

			//filter for Marcel
			if isAcceptedByMarcel(wordFiltered) {
				clearedWordsWithOccurrenceForMarcel[wordFiltered] += 1
			}
		}

	}

	return
}

func getSortedResultToPrintForMarcelAndPaulina(clearedWordsWithOccurrenceForMarcel, clearedWordsWithOccurrenceForPaulina map[string]int) (resultForMarcelToPrint, resultForPaulinaToPrint []string) {
	//we store our data in PairList to be able to use Sort

	//creating result for Paulina
	sortedResultForPaulina := make(PairList, len(clearedWordsWithOccurrenceForPaulina))

	i := 0
	for key, value := range clearedWordsWithOccurrenceForPaulina {
		sortedResultForPaulina[i] = Pair{key, value}
		i++
	}

	//sort
	sort.Sort(sort.Reverse(sortedResultForPaulina))
	resultForPaulinaToPrint = make([]string, len(clearedWordsWithOccurrenceForPaulina))

	i = 0
	for _, k := range sortedResultForPaulina {
		stringVal := strconv.Itoa(k.Value)
		resultForPaulinaToPrint[i] = k.Key + " - " + stringVal
		i++
	}

	//creating result for Marcel
	sortedResultForMarcel := make(PairList, len(clearedWordsWithOccurrenceForMarcel))

	i = 0
	for key, value := range clearedWordsWithOccurrenceForMarcel {
		sortedResultForMarcel[i] = Pair{key, value}
		i++
	}

	//sort
	sort.Sort(sort.Reverse(sortedResultForMarcel))
	resultForMarcelToPrint = make([]string, len(clearedWordsWithOccurrenceForMarcel))

	i = 0
	for _, k := range sortedResultForMarcel {
		stringVal := strconv.Itoa(k.Value)
		resultForMarcelToPrint[i] = k.Key + " - " + stringVal
		i++
	}

	return
}

func isAcceptedByPaulina(val string) bool {
	matched := strings.Contains(val, "nie")
	return !matched
}

func isAcceptedByMarcel(val string) bool {
	matched := strings.ContainsAny(val, "ąćęłńóśżź")
	return !matched
}
