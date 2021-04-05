// ## 4. Play with numbers

// Create row of numbers from **0 to 1000**. Then **group that numbers** by:

// 1) even, 
// 2) odd, 
// 3) numbers multiplies by 5,
// 4) prime numbers, and
// 5) prime numbers less than 100

package main 
import (
	"fmt"
)

func GroupBy(input int) {
	var even, odd, five []int
	
	for i := 1; i <= input; i++ {
		if i%5 == 0 {
			five = append(five, i)
		}
		if i%2 == 0 {
			even = append(even, i)
		} 
		if i%2 != 0{
			odd = append(odd, i)
		}	
	}
	fmt.Println(even)
	fmt.Println(odd)
	fmt.Println(five)
}

func main () {
	GroupBy(1000)
}