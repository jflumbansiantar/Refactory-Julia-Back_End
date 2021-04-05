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
	// "time"
)

func Counting(str string, countWord []string) string {
	
	inputStr := strings.Fields(str)

	if len(countWord) <= 0 {
		return str
	}

    for _, word := range inputStr {
        for _, count := range countWord {
			if test := strings.Index(strings.ToLower(word), count); test >= 0 {
				_, matched := countWord[]
				if matched {
					countWord[count] += 1
				}
			}
		}
	}
	return countWord
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
	var input = "Aku ingin begini Aku ingin begitu Ingin ini itu banyak sekali Semua semua semua Dapat dikabulkan Dapat dikabulkan Dengan kantong ajaib Aku ingin terbang bebas Di angkasa Hei… baling baling bambu La... la... la... Aku sayang sekali Doraemon La... la... la... Aku sayang sekali"
	
	result := Counting(input, censorWords)

	fmt.Println(result)
}
