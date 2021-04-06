// ## 10. Rest API

// Fetch data from `https://jsonplaceholder.typicode.com/posts` and 
// `https://jsonplaceholder.typicode.com/users`. Then combine `posts` and `users` based on `userId`.

// Expected output:

// ```json
// [
//   {
//     "userId": 1,
//     "id": 1,
//     "title": "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
//     "body": "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
//     "user": {
//       "id": 1,
//       "name": "Leanne Graham",
//       "username": "Bret",
//       "email": "Sincere@april.biz",
//       "address": {
//         "street": "Kulas Light",
//         "suite": "Apt. 556",
//         "city": "Gwenborough",
//         "zipcode": "92998-3874",
//         "geo": {
//           "lat": "-37.3159",
//           "lng": "81.1496"
//         }
//       },
//       "phone": "1-770-736-8031 x56442",
//       "website": "hildegard.org",
//       "company": {
//         "name": "Romaguera-Crona",
//         "catchPhrase": "Multi-layered client-server neural-net",
//         "bs": "harness real-time e-markets"
//       }
//     }
//   },
// ```

package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	// "encoding/json"
	"os"
	// "bytes"
)


func main() {
	
	response1, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData1, err := ioutil.ReadAll(response1.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	data1 := string(responseData1)
	
	response2, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}
	data2 := string(responseData2)



	fmt.Println(data1)
	fmt.Println(data2)
	
	
}
	















	//PSEUDOCODE-nya
	/**
	if responseData1.userId == responseData2.id {
		data1.push(data2)
		} 
		*/
		
		// if strings.Contains(responseData1.userId) = string(responseData2.id) {
			// 	return fmt.Sprin)
			// }
			
			// bMap := make(map[string]responseData2)
			
			// for _, b := range responseData2 {
				// 	bMap[b.id] = b
	// }

	// merged := make([]responseData1, len(responseData1))
	// for i, a := range responseData1 {
	// 	if b, ok := bMap[a.userId]; ok {
	// 		a.userId = b.id
	// 	}
	// 	merged[i] = a
	// }
	// return merged
	
// }