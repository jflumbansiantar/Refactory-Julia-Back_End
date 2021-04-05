// ## 2. Censoring Words

// Censors these words: dolor, elit, quis, nisi, fugiat, proident, laborum; of this paragraph.

// ```
// Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
// Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.
// Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.
// Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
// ```

// The output should be

// ```
// Lorem ipsum ***** sit amet, consectetur adipisicing ****, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
// Ut enim ad minim veniam, **** nostrud exercitation ullamco laboris **** ut aliquip ex ea commodo consequat.
// Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu ****** nulla pariatur.
// Excepteur sint occaecat cupidatat non ******** sunt in culpa qui officia deserunt mollit anim id est *******.
// ```

package main 
import (
	"fmt"
	"strings"
	
)

func Censoring(str string, censored []string) string {
	var newSlice []string

	if len(censored) <= 0 {
		return str
	}

	strSlice := strings.Fields(str)

	for position, word := range strSlice {
		for _, censorWord := range censored {
			if test := strings.Index(strings.ToLower(word), censorWord); test >= 0 {
				replacement := strings.Repeat("*", len(word))
				strSlice[position] = replacement
				newSlice = append(strSlice[:position], strSlice[position:]...)
			}
		}
	}
	return strings.Join(newSlice, " ")
}

func main () {
	var censorWords = []string{
		"dolor", 
		"elit", 
		"quis", 
		"nisi", 
		"fugiat", 
		"proident", 
		"laborum",
	}

	var inputString = "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

	result := Censoring(inputString, censorWords)

	fmt.Println(result)
}