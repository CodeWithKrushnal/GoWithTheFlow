package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Roman to Integer Converter")
	fmt.Println("Kindly Enter the Roman Number on following Line:")

	//Decaring and initializing the map for values
	var values = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	//Declaring variable  and Accepting the input into variable roman
	var roman string
	fmt.Scanln(&roman)
	roman = strings.ToUpper(roman)

	//Trimming the Input to Exclude Excessive unnecessory spaces
	roman = strings.Trim(roman, " ")

	//Regex pattern
	validRomanPattern := `^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$`

	//Regex Compilation
	pattern := regexp.MustCompile(validRomanPattern)

	//Validation
	if !pattern.MatchString(roman) {
		fmt.Println("Invalid Roman Number")
		return
	}

	var Total int = 0

	var length int = len(roman)

	for i := 0; i < length; i++ {
		current := string(roman[i])
		extractedValue := values[current]
		if extractedValue == 0 {
			fmt.Print("Error for the Invalid input ", current)
			break
		} else {
			if i+1 < length && extractedValue < values[string(roman[i+1])] {
				Total -= extractedValue
			} else {
				Total += extractedValue
			}
		}
	}

	fmt.Println("The Integer value for roman value", roman, "is", Total)
}

