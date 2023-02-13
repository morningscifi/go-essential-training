// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string
	Name     string
	NumRepos int `json:"public_repos"`
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	u := fmt.Sprintf("https://api.github.com/users/%s", login)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatalf("GET failed")
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	var user User
	if err := dec.Decode(&user); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}

	fmt.Printf("got: %+v\n", user)

	return &user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
