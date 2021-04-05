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
	make a variable that contains all the alphabetic characters (does Golang has it ?)
	lowercase all the letters
	compare part of the alphabetic characters = letters
	if something a miss from the part of the alphabetic characters, print that missing letter

	*/

}

func main() {
	var unit = "abde"

	fmt.Println(MissingLetter(unit))
}
