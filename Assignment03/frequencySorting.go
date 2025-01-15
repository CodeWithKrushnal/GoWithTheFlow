package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Creating a sortable unit of Word and frequency Pairs
type sortabaleUnit struct {
	word  string
	count int
}

func main() {
	fmt.Println("Frequency based Sorting Program")
	fmt.Println("Kindly Enter the String of which the words need to be sorted")

	//Declaring the String to be Accepted
	var userInput string

	//Declaring the map to register the frequency of words
	var frequency = map[string]int{}

	//Scanning the Inputs from the Console including Whitespaces
	fullLineScanner := bufio.NewScanner(os.Stdin)
	fullLineScanner.Scan()
	userInput = fullLineScanner.Text()

	//Trimming the Received input to get rid of unncessary leading and laggin whitespaces
	userInput = strings.Trim(userInput, " ")

	//Splitting the trimmed input in a Slice of Words
	values := strings.Split(userInput, " ")

	//Registering the frequency of words in map
	for _, val := range values {
		if val != " " {
			olderFrequency, isPresent := frequency[val]
			if isPresent {
				frequency[val] = olderFrequency + 1
			} else {
				frequency[val] = 1
			}
		} else {
			continue
		}
	}

	//Creating a slice of type sortabaleUnit
	sortabaleSlice := []sortabaleUnit{}
	for key, value := range frequency {
		sortabaleSlice = append(sortabaleSlice, sortabaleUnit{key, value})
	}

	//Sorting the sortableSlice on the basis of frequency
	sort.Slice(sortabaleSlice, func(i, j int) bool {
		return sortabaleSlice[i].count > sortabaleSlice[j].count
	})

	highestFrequency := sortabaleSlice[0].count

	for _, element := range sortabaleSlice {
		if element.count == highestFrequency {
			fmt.Print(element.word, " ")
		} else {
			break
		}
	}

	fmt.Println("")
}
