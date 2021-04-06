// ## 6. Missing Letter

// If there is a list of ordered letters, and one of them is missing. Find out the missing letter!

// Example :

// ```javascript
// const list_letters_first = ["c","d","e","g","h"]
// const list_letters_second = ["X","Z"]
// ```

// Expected output :

// ```
// list_letters_first = f
// list_letters_second = Y
// ```

package main

import (
	"fmt"
)

func MissingLetter(s string) {
	/**PSEUDOCODE
	make a variable that contains all the alphabetic characters 
	lowercase all the letters
	search first lettercase in alphabetic order
	compare part of the alphabetic characters = letters
	if something a miss from the part of the alphabetic characters, print that missing letter

	*/
	var alphabetic = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var missing []string;

	newString = strings.ToLower(s)

	for i, letter := range newString {
		for j, chars := range alphabetic {
			if letter[i] != chars[j] {
				missing = append(missing, v[j])
			}
		}
	}
	return missing
}

func main() {
	var unit = ["a", "b", "d", "e"]

	fmt.Println(MissingLetter(unit))
}
