// twitter_status_json.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Status struct {
	Text string
}

type User struct {
	Status Status
}

func main() {
	/* perform an HTTP request for the twitter status of user: Googland */
	res, _ := http.Get("http://twitter.com/users/Googland.json")/*http://twitter.com/users/Googland.json頁面不存在了*/
	/* initialize the structure of the JSON response */
	user := User{Status{""}}
	/* unmarshal the JSON into our structures */
	temp, _ := ioutil.ReadAll(res.Body)
	body := []byte(temp)
	json.Unmarshal(body, &user)
	fmt.Printf("status: %s", user.Status.Text)
}

/* Output:
status: Robot cars invade California, on orders from Google:
Google has been testing self-driving cars ... http://bit.ly/cbtpUN http://retwt.me/97p
*/