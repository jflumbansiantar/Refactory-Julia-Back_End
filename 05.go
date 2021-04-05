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

func Counting(str string) map[string]int {

	letters := strings.Split(strings.ToLower(str), "")

		counts := make(map[string]int)
		for _, char := range letters {
			_, matched := counts[char]
			
			if matched {
				counts[char] += 1
				
			} else {
				counts[char] = 1
			}
		}
		return (counts)
	
}

func main () {
	var input = "WHite whalE"
	result := Counting(input)
	delete(result, " ")

	fmt.Println(result)
}