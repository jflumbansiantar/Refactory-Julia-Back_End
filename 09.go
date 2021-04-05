// ## 9. Array Sort

// Sort the array of number from min to max

// Example :

// ```javascript
// const arrayList = [3, 12, 4, 5, 8, 9]
// const sortMethod =(items)=>{
// 	// your code
// }
// sortMethod(arrayList)
// ```

// Expected output :

// ```
// [3, 4, 5, 8, 9, 12]
// ```

package main 
import (
	"fmt"
	"sort"
)

func main () {
	numbers := []int{3, 12, 4, 5, 8, 9}
	sort.Ints(numbers)
	fmt.Println(numbers)
}