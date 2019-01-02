package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

type todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type user struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Todos    []todo `json:"todos"`
}

func (u *user) GetTodos(wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos?userId=%d", u.ID)
	res, reqErr := http.Get(url)

	if reqErr != nil {
		panic(reqErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		panic(readErr)
	}

	todos := []todo{}
	jsonErr := json.Unmarshal(body, &todos)

	if jsonErr != nil {
		panic(jsonErr)
	}

	u.Todos = todos
}

func getUsers(ch chan<- []user) {
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
