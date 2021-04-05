// ## 1. Play with CSV and object

// We have a collection of product and price with CSV formats:

// ```csv
// NAME, CATEGORY, PRICE
// Xiaomi Redmi 5A, Smartphone, 1199000
// Macbook Air, Laptop, 13775000
// Samsung Galaxy J7, Smartphone, 3549000
// DELL XPS 13, Laptop, 26799000
// Xiaomi Mi 6, Smartphone, 5399000
// LG V30 Plus, Smartphone, 10499000
// ```

// **Sort by price** and **transforms** those collection into object within an array like this format:

// ```json
// [
//   {
//     "name": "Macbook Air",
//     "price": "Rp13.775.000",
//     "category": "Laptop"
//   }
// ]
// ```

package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"fmt"
	"strings"
)


func main() {

	type Product struct {
      	Name string
      	Category string
      	Price string
   	}
   
   	productline, err := csv.NewReader(CsvFile).ReadAll()
   	if err != nil {
      	log.Fatal(err)
   	}

   	prods := []Product{}

   	for i, prod := range productline {
    	if i == 0 {
			continue
      	}
		p := Product{prod[0],prod[1],prod[2]}
      	prods = append(prods,[]Product{p}...)
   }

   	res, _ := json.MarshalIndent(prods,"","   ")
   
   	fmt.Printf("%s\n",res)
}

var CsvFile = strings.NewReader(`Name,Category,Price
Xiaomi Redmi 5A,Smartphone,1199000
Macbook Air,Laptop,13775000
Samsung Galaxy J7,Smartphone,3549000
DELL XPS 13,Laptop,26799000
Xiaomi Mi 6,Smartphone,5399000
LG V30 Plus,Smartphone,10499000`)