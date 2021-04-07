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
	"math"
)

func GroupBy(input int) {
	var even, odd, five, prime, primeLess []int
	
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
		if i <= input {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				prime = append(prime, i)
			}
		}
		if i <= 100 {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primeLess = append(primeLess, i)
			}
		}
	}
	fmt.Println(even, " ")
	fmt.Println(odd, " ")
	fmt.Println(five, " ")
	fmt.Println(prime, " ")
	fmt.Println(primeLess, " ", "prime less 100")
}

func main () {
	GroupBy(10)
}