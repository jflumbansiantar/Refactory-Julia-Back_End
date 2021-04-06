// ## 7. Sorting Odd Numbers

// Given a list of even and odd numbers. Sort the odd numbers, without change even numbers position.
// p.s : zero is an even number.

// Example :

// ```js
// const numbers = [9,4,2,4,1,5,3,0]
// sort_odd(numbers)
// ```

// Expected output :

// ```
// [1, 4, 2, 4, 3, 5, 9, 0]

package main
import "fmt"
func SortOdd(numbers []int) []int {

	for i := 0; i < len(numbers); i++ {
      	for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				if numbers[j-1] % 2 != 0 {
					intermediate := numbers[j]
					numbers[j] = numbers[j-1]
					numbers[j-1] = intermediate
				}
		   }
		} 
    }
   return numbers
 }

func main() {
    a := []int{9,4,2,4,1,5,3,0}
    fmt.Println(SortOdd(a))
}



/** PSEUDOCODE
for i = 0, i < len(numbers), i ++
	for j = 1, j < i, j ++ 
		if numbers[j] % 2 == 0 
			numbers[j-1], numbers[j] = numbers[j-1], numbers[j]
		else 	 
			if numbers[j-1] > numbers[j] {
				   intermediate := numbers[j]
				   numbers[j] = numbers[j-1]
				   numbers[j-1] = intermediate
			   }
*/