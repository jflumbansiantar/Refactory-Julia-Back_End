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
//   {
//     "userId": 1,
//     "id": 2,
//     "title": "qui est esse",
//     "body": "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
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
//   {
//     "userId": 6,
//     "id": 53,
//     "title": "ut quo aut ducimus alias",
//     "body": "minima harum praesentium eum rerum illo dolore\nquasi exercitationem rerum nam\nporro quis neque quo\nconsequatur minus dolor quidem veritatis sunt non explicabo similique",
//     "user": {
//       "id": 6,
//       "name": "Mrs. Dennis Schulist",
//       "username": "Leopoldo_Corkery",
//       "email": "Karley_Dach@jasper.info",
//       "address": {
//         "street": "Norberto Crossing",
//         "suite": "Apt. 950",
//         "city": "South Christy",
//         "zipcode": "23505-1337",
//         "geo": {
//           "lat": "-71.4197",
//           "lng": "71.7478"
//         }
//       },
//       "phone": "1-477-935-8478 x6430",
//       "website": "ola.org",
//       "company": {
//         "name": "Considine-Lockman",
//         "catchPhrase": "Synchronised bottom-line interface",
//         "bs": "e-enable innovative applications"
//       }
//     }
//   }
// ]
// ```

package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Data struct {
	userId string
	id string
	title string
	body string
	user string
}

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
	fmt.Println(string(responseData1))

	response2, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData2, err := ioutil.ReadAll(response2.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData2))

	//PSEUDOCODE-nya
	/**
	if responseData1.userId = responseData2.id {
		populate["user"]
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
	

}