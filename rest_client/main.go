package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	UserIdenNo int `json:"userId"`
	IdenNo int `json:"id"`
	TextTitle string `json:"title"`
	IsCompleted bool `json:"completed"`
}

func main() {

	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/4")

	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	log.Println(string(body))

	var todo Todo
	json.Unmarshal(body, &todo)
	
	log.Println(todo)
}