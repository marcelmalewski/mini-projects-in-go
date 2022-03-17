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

//import (

//)

//tworzenie nowego typu
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }

//potrzebne do sortowania
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

var data = []string{"a", "string", "list"}

func main() {
	content, err := ioutil.ReadFile("ogniem_i_mieczem.txt")

	if err != nil {
		log.Fatal(err)
	}

	//zmienia string na array słow usuwa wszystkie spacje i entery
	slicedFile := strings.Fields(string(content))
	//mapy do liczenia slow dla Pauliny i Marcela
	mapNumberOfWordsFilteredForPaulina := make(map[string]int)
	mapNumberOfWordsFilteredForMarcel := make(map[string]int)

	for i := 0; i < len(slicedFile); i++ {
		//zeby nie przejmowac sie wielkoscią liter
		wordToLower := strings.ToLower(slicedFile[i])
		wordToLowerSplit := strings.Split(wordToLower, "")
		var wordToLowerFiltered string

		//czyscimy stringa ze znakow interpunkcyjnych i spacji itd.
		for i := range wordToLowerSplit {
			if isLetterOrNumber(wordToLowerSplit[i]) {
				wordToLowerFiltered += wordToLowerSplit[i]
			}
		}

		if wordToLowerFiltered != "" {
			//zestaw dla Pauliny
			if isAcceptedByPaulina(wordToLowerFiltered) {
				mapNumberOfWordsFilteredForPaulina[wordToLowerFiltered] += 1
			}

			//zestaw dla Marcela
			if isAcceptedByMarcel(wordToLowerFiltered) {
				mapNumberOfWordsFilteredForMarcel[wordToLowerFiltered] += 1
			}
		}

	}

	//tworzenie zestawu Pauliny
	resultForPaulina := make(PairList, len(mapNumberOfWordsFilteredForPaulina))

	i := 0
	for key, value := range mapNumberOfWordsFilteredForPaulina {
		resultForPaulina[i] = Pair{key, value}
		i++
	}

	//sortowanie
	sort.Sort(sort.Reverse(resultForPaulina))
	resultForPaulinaSlice := make([]string, len(mapNumberOfWordsFilteredForPaulina))

	i = 0
	for _, k := range resultForPaulina {
		stringVal := strconv.Itoa(k.Value)
		resultForPaulinaSlice[i] = k.Key + " - " + stringVal
		i++
	}

	//tworzenie zestawu Marcela
	resultForMarcel := make(PairList, len(mapNumberOfWordsFilteredForMarcel))

	i = 0
	for key, value := range mapNumberOfWordsFilteredForMarcel {
		resultForMarcel[i] = Pair{key, value}
		i++
	}

	//sortowanie
	sort.Sort(sort.Reverse(resultForMarcel))
	resultForMarcelSlice := make([]string, len(mapNumberOfWordsFilteredForMarcel))

	i = 0
	for _, k := range resultForMarcel {
		stringVal := strconv.Itoa(k.Value)
		resultForMarcelSlice[i] = k.Key + " - " + stringVal
		i++
	}

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
	//listMarcel.Resize(fyne.NewSize(300, 300))

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
	//listPaulina.Resize(fyne.NewSize(300, 300))

	//test := container.New(layout.NewGridLayout(2), title, list, title, list)
	vBox := container.New(layout.NewGridWrapLayout(fyne.NewSize(300, 300)), titleMarcel, listMarcel, titlePaulina, listPaulina)

	myWindow.SetContent(vBox)
	myWindow.ShowAndRun()
}

func isLetterOrNumber(val string) bool {
	matched, _ := regexp.MatchString(`[a-z0-9ąćęłńóśżź]`, val)
	return matched
}

func isAcceptedByPaulina(val string) bool {
	matched := strings.Contains(val, "nie")
	return !matched
}

func isAcceptedByMarcel(val string) bool {
	matched := strings.ContainsAny(val, "ąćęłńóśżź")
	return !matched
}
