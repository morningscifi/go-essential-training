// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"user"`
	Name     string `json:"name"`
	NumRepos int    `json:"public_repos"`
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	/* TODO:
	Call the github API for a given login
	e.g. https://api.github.com/users/tebeka


	And return User struct
	*/
	resp, err := http.Get("https://api.github.com/users/tebeka")
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
