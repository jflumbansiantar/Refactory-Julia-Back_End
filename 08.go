// Given a string "23dn3ir30fd2eddd", please change all the characters with * (asterisk) except last 3 characters. Write a function to solve this.

// Example :

// ```javascript
// const secret_text = "23dn3ir30fd2eddd"
// masking(secret_text)
// ```

// Expected output :

// ```
// *************ddd
// ```

package main 
import (
	"fmt"
)

func Masking(str string) string{
	 re := []rune(str)

	 for i := 0; i < len(re)-3; i++ {
		 re[i] = '*'
	 }

	 return string(re)
}

func main () {
	var input = "23dn3ir30fd2eddd"
	fmt.Println(Masking(input))
}