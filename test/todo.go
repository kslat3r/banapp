package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func listTodos(ch chan<- []todo) {
	res, reqErr := http.Get("https://jsonplaceholder.typicode.com/todos")

	if reqErr != nil {
		panic(reqErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		panic(readErr)
	}

	list := []todo{}
	jsonErr := json.Unmarshal(body, &list)

	if jsonErr != nil {
		panic(jsonErr)
	}

	ch <- list
}
