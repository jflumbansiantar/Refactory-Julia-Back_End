// // ## 3. Count words

// // **Count** each of these words: `aku`, `ingin` and `dapat`.

// // ```
// // Aku ingin begini
// // Aku ingin begitu
// // Ingin ini itu banyak sekali

// // Semua semua semua
// // Dapat dikabulkan
// // Dapat dikabulkan
// // Dengan kantong ajaib

// // Aku ingin terbang bebas
// // Di angkasa
// // Hei… baling baling bambu

// // La... la... la...
// // Aku sayang sekali
// // Doraemon

// // La... la... la...
// // Aku sayang sekali
// // ```

package main 
import (
	"fmt"
	"strings"
)

func Counting(str string) map[string]int {

	inputString := strings.Fields(strings.ToLower(str))
    wordCount := make(map[string]int)
    for _, word := range inputString {
        _, matched := wordCount[word]
        if matched {
            wordCount[word] += 1
        } else {
            wordCount[word] = 1
        }
    }
    return wordCount
	
}

func main () {

	var input = "Aku ingin begini Aku ingin begitu Ingin ini itu banyak sekali Semua semua semua Dapat dikabulkan Dapat dikabulkan Dengan kantong ajaib Aku ingin terbang bebas Di angkasa Hei… baling baling bambu La... la... la... Aku sayang sekali Doraemon La... la... la... Aku sayang sekali"
	
	result := Counting(input)
	fmt.Println(`Banyak AKU adalah `,result["aku"])
	fmt.Println(`Banyak INGIN adalah `,result["ingin"])
	fmt.Println(`Banyak DAPAT adalah `,result["dapat"])
	
}
