package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("ogniem_i_mieczem.txt")

	if err != nil {
		log.Fatal(err)
	}

	slicedFile := strings.Fields(string(content))

	for i := 0; i < len(slicedFile); i++ {
		wordToLower := strings.ToLower(slicedFile[i])
		wordToLowerSplit := strings.Split(wordToLower, "")
		var wordToLowerSplitFiltered []string

		for i := range wordToLowerSplit {
			if isLetterOrNumber(wordToLowerSplit[i]) {
				wordToLowerSplitFiltered = append(wordToLowerSplitFiltered, wordToLowerSplit[i])
			}
		}

		//fmt.Println(filteredWord)
		fmt.Println(wordToLowerSplitFiltered)
	}
}

func isLetterOrNumber(val string) bool {
	matched, _ := regexp.MatchString(`[a-z0-9ąćęłńóśżź]`, val)
	return matched
}
