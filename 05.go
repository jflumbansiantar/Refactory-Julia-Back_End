// ## 5 Char Counter

// Given a string and count how many occurrence of each character of the string. The output is a dictionary data type, key for chars and value is sum of the chars. The character only letter, not space, or any symbols.

// Example :

// ```javascript
// const text_1 = "Mammals"
// const text_2 = "Bruiser build"
// ```

// Expected output :

// ```
// {'m': '***', 'a': '**', 'l': '*', 's': '*'}
// {'b': '**', 'r': '**', 'u': '**', 'i': '**', 's': '*', 'e': '*', 'l': '*', 'd': '*'}
// ```

package main 
import (
	"fmt"
	"strings"
)

func Counting(str string) {
	// var newSlice []string

	strSlice := strings.Split(str, ",")
	fmt.Println(strSlice, "slice")

	for _, word := range strSlice {

		fmt.Println(word, "slice2")
		// for _, matched := range 
	}
}

func main () {
	var input = "mammals"

	Counting(input)
}