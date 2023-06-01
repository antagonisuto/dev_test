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
End time: 16:37
Execution time:  862.798105ms - 878ms
***/

type Frequency struct {
	word []byte // word
	f    int    //frequency
}

// Execution time:  862.798105ms
func main() {
	//read file, argument $1 - filename
	start := time.Now()

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a filename as an argument.")
		return
	}

	filename := args[1]
	fmt.Printf("Filename: %s\n\n", filename)

	//store all words - 123.347µs
	unformattedSymbols, err := os.ReadFile(filename)
	check(err)

	//remove symbols, lower letters and make []byte words - 4.431796ms
	formattedWords := formatLetter(unformattedSymbols)

	//count frequency - 813.962397ms - huge process
	unsortedLibraryWord := libraryWord(formattedWords)

	//sort frequency by number
	sortedFrequency := sort(unsortedLibraryWord)

	fmt.Printf("Counted: %d\n", len(sortedFrequency))
	for i := 0; i < 20; i++ {
		fmt.Printf("%d\t %s\t \n", sortedFrequency[i].f, sortedFrequency[i].word)
	}

	// Code to measure
	fmt.Println("Execution time: ", time.Since(start))
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

// method for count the words
func libraryWord(formattedWords [][]byte) []Frequency {
	length := len(formattedWords)

	//struct for saving the word and count, it can be replaced as two variable []byte and []int
	var result []Frequency
	freq := 0
	checked := make([]int, length)

	//loop over words
	for i := 0; i < length; i++ {

		//if already checked than ignore the word
		if checked[i] == 1 {
			continue
		}

		for j := i; j < length; j++ {
			//if already checked than ignore the word
			if checked[j] == 1 {
				continue
			}

			//compare two []byte words,
			if testEq(formattedWords[i], formattedWords[j]) == true {
				//if equal then mark word as checked and add +1 top freq
				checked[j] = 1
				freq = freq + 1
				continue
			}
		}
		// add unique word with frequency
		result = append(result, Frequency{word: formattedWords[i], f: freq})

		//init var to 0 for next word
		freq = 0
	}

	fmt.Printf("All unique words: %d\n", len(result))

	return result
}

func testEq(a, b []byte) bool {
	//check if two var has diff size
	if len(a) != len(b) {
		return false
	}

	//check all stuff
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	//if two var is equal
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func formatLetter(unformattedLetter []byte) [][]byte {
	//all words
	var words [][]byte

	//temp word or collection of letter
	var formattedLetter []byte

	//loop over all symbols to find letters
	for _, byteLetter := range unformattedLetter {
		//UTF-8 Decimal format
		//if Upper letter make Lower
		if byteLetter >= 65 && byteLetter <= 90 {
			byteLetter = byteLetter + 32
		}
		//if lower letter save into formatterLetter
		if byteLetter >= 97 && byteLetter <= 122 {
			formattedLetter = append(formattedLetter, byteLetter)
		} else if len(formattedLetter) != 0 {
			//if it is not a letter, than save it as word, and make empty temp var
			words = append(words, formattedLetter)
			formattedLetter = nil
		}
	}

	//return collection of words
	return words
}
