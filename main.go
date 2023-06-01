package main

import (
	"fmt"
	"os"
	"time"
)

/***

The Problem

Given the attached text file as an argument,
your program will read the file, and output
the 20 most frequently used words in the file in order,
along with their frequency.

Start time: 12:10PM

***/

type Frequency struct {
	word []byte // word
	f    int    //frequency
}

// Execution time:  862.798105ms
func main() {
	//read file, argument $1 - filename
	start := time.Now()

	// show := 150
	args := os.Args
	var filename string
	if len(args) == 4 {
		filename = args[3]
	} else {
		filename = args[1]
	}
	fmt.Printf("Filename: %s\n\n", filename)

	//store all words
	unformattedSymbols := readFile(filename)
	// fmt.Printf("unformattedSymbols: %s\n\n", unformattedSymbols[:show])

	//remove symbols and make words
	formattedWords := formatLetter(unformattedSymbols)
	// fmt.Printf("formattedWords: %s - %d\n\n", formattedWords[:show], len(formattedWords))

	//unique words
	// libraryWords := uniqueWords(formattedWords)

	//count
	unsortedLibraryWord := libraryWord(formattedWords)
	// fmt.Printf("unformattedSymbols: %s - %d\n\n", countedWords[0], len(countedWords))
	// //print

	sortedFrequency := sort(unsortedLibraryWord)

	fmt.Printf("Counted: %d\n", len(sortedFrequency))
	for idx, f := range sortedFrequency {
		if idx == 20 {
			break
		}
		fmt.Printf("%d\t %s\t \n", f.f, f.word)
	}

	// Code to measure
	duration := time.Since(start)
	fmt.Println("Execution time: ", duration)
}

func sort(arr []Frequency) []Frequency {
	for i := 0; i <= len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j].f < arr[j+1].f {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// func printFrequency(countedWords []Frequency) {

// }

func libraryWord(formattedWords [][]byte) []Frequency {
	length := len(formattedWords)

	var result []Frequency
	freq := 0
	checked := make([]int, length)

	for i := 0; i < len(formattedWords); i++ {
		if checked[i] == 1 {
			continue
		}

		for j := i; j < len(formattedWords); j++ {
			if checked[j] == 1 {
				continue
			}
			if testEq(formattedWords[i], formattedWords[j]) == true {
				checked[j] = 1
				freq = freq + 1
				continue
			}
		}
		result = append(result, Frequency{word: formattedWords[i], f: freq})
		freq = 0
	}

	fmt.Printf("All unique words: %d\n", len(result))

	return result
}

func testEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(filename string) []byte {
	words, err := os.ReadFile(filename)
	check(err)
	return words
}

func formatLetter(unformattedLetter []byte) [][]byte {
	var words [][]byte
	var formattedLetter []byte
	for _, byteLetter := range unformattedLetter {
		// if (byteLetter == 10 || byteLetter == 32) && len(formattedLetter) != 0 {
		// 	words = append(words, formattedLetter)
		// 	formattedLetter = nil
		// }
		if byteLetter >= 65 && byteLetter <= 90 {
			byteLetter = byteLetter + 32
		}
		if byteLetter >= 97 && byteLetter <= 122 {
			formattedLetter = append(formattedLetter, byteLetter)
		} else if len(formattedLetter) != 0 {
			words = append(words, formattedLetter)
			formattedLetter = nil
		}
	}
	return words
}
