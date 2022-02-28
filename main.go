package main

import (
	"unicode"
)

/*
STRING SPECS

Given a string that is a sequence of numbers followed by dash followed by text, eg: 23-ab-48-caba-56-haha

-The numbers can be assumed to be unsigned integers
-The strings can be assumed to be ASCII sequences of arbitrary length larger than 0 (empty strings not allowed).

TASK 1

-Before implementing the needed function, please estimate the difficulty of the functions and report the
estimated and the used time to implement the functions in the function comment block.

-Write a function testValidity that takes the string as an input, and returns boolean flag true if the given
string complies with the format, or false if the string does not comply

-Write a function averageNumber that takes the string, and returns the average number from all the numbers

-Write a function wholeStory that takes the string, and returns a text that is composed from all the text
 words separated by spaces, e.g. story called for the string 1-hello-2-world should return text: "hello world"

 -Write a function storyStats that returns four things:
	-the shortest word
	-the longest word
	-the average word length
	-the list (or empty list) of all words from the story that have the length the same as the average length rounded up and down.
*/

//-Write a function testValidity that takes the string as an input, and returns boolean flag true if the given
//string complies with the format, or false if the string does not comply

//Estimated difficulty testValidity: low. Estimated time: 15 minutes

//testValidity takes a string and returns true or false depending on whether it does or does not comply
//with the format number-text[(optional)-number-text-...]
//Check stores the last two values in the string (0 for numbers, 1 for - and 2 for letters), it´s used
//to validate if the string format fits the required
func testValidity(s string) bool {
	//Check represents the last two chars.
	//It is 0 when the char was an int, 1 when it was -, 2 when it was a letter
	var check string = "00"
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) && (check == "21" || check == "00" || check == "10") {
			if check == "21" {
				check = "10"
			} else {
				check = "00"
			}
			if i == len(s)-1 {
				return true
			}
		} else if string(s[i]) == `-` && (check == "00" || check == "10" || check == "22" || check == "12") {
			if check == "00" || check == "10" {
				check = "01"
			} else if check == "22" || check == "12" {
				check = "21"
			}
		} else if unicode.IsLetter(rune(s[i])) && (check == "22" || check == "12" || check == "01") {
			if check == "22" || check == "12" {
				check = "22"
			} else {
				check = "12"
			}
			if i == len(s)-1 {
				return true
			}
		} else {
			return false
		}

	}
	return false
}

//-Write a function averageNumber that takes the string, and returns the average number from all the numbers

//Function difficulty: low. Expected time: 10 minutes or less.

//averageNumber takes a string and returns the average number of all the numbers in the
//string (assuming that "each number" refers to each digit in the string, not to each whole number
//e.g: 123-as-1-d would return 1.75)
func averageNumber(s string) float64 {
	var cantNums float64 = 0
	var sum float64 = 0
	for i := 0; i < len(s); i++ {
		if unicode.IsNumber(rune(s[i])) {
			cantNums += 1
			sum += float64(s[i] - '0')
		}
	}
	return float64(sum / cantNums)
}

//-Write a function wholeStory that takes the string, and returns a text that is composed
//from all the text words separated by spaces, e.g. story called for the string 1-hello-2-world
//should return text: "hello world"

//Difficulty: Low. Expected time: 10 minutes or less

//wholeStory takes a string and returns the words in it with a space between different words
//check is used to test whether it is or not required to add a space before the next char
func wholeStory(s string) string {
	//check is 0 while the last character was a letter and 1 when it wasn´t
	check := 0
	str := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			if check == 1 {
				str = str + " "
			}
			str = str + string(s[i])
			check = 0
		} else {
			check = 1
		}
	}
	return str
}

/*
-Write a function storyStats that returns four things:
	-the shortest word
	-the longest word
	-the average word length
	-the list (or empty list) of all words from the story that have
	the length the same as the average length rounded up and down.
*/

//Difficulty: Low. Estimated time:20 minutes

//storyStats takes a string and returns the shortest word, the longest word,
//the average word length and the list of all words from the string.
//If the string does not contain text, it returns an empty string for
//shortest and longest, 0 for the average word length and an empty list of strings
func storyStats(s string) (string, string, float64, []string) {
	var words []string
	check := 1
	str := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsLetter(rune(s[i])) {
			str = str + string(s[i])
			check = 0
		} else {
			if check == 0 {
				words = append(words, str)
				str = ""
			}
			check = 1
		}
	}
	if str != "" {
		words = append(words, str)
	}

	if len(words) > 0 {
		shortest := words[0]
		longest := words[0]
		var length float64 = 0
		for i := 0; i < len(words); i++ {
			if len(words[i]) > len(longest) {
				longest = words[i]
			}
			if len(words[i]) < len(shortest) {
				shortest = words[i]
			}
			length += float64(len(words[i]))
		}
		return shortest, longest, length / float64(len(words)), words
	}
	return "", "", 0, words

}

func main() {
}
