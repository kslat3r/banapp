package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Todos    []todo `json:"todos"`
}

func listUsers(ch chan<- []user) {
	res, reqErr := http.Get("https://jsonplaceholder.typicode.com/users")

	if reqErr != nil {
		panic(reqErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		panic(readErr)
	}

	list := []user{}
	jsonErr := json.Unmarshal(body, &list)

	if jsonErr != nil {
		panic(jsonErr)
	}

	ch <- list
}
