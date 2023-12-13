package main

import (
	"context"
	"fmt"

	"github.com/haquenafeem/gopunch"
)

type TypeOFRequest int

const (
	GetSingle TypeOFRequest = iota
	GetAll
	Post
	Delete
)

func main() {
	TestJSONPlaceHolder(GetAll)
}

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func TestJSONPlaceHolder(reqType TypeOFRequest) {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	switch reqType {
	case GetAll:
		getAll(client)
		return
	case GetSingle:
		getSingle(client)
		return
	case Post:
		post(client)
		return
	case Delete:
		delete(client)
		return
	}
}

func getAll(client *gopunch.Client) {
	query := gopunch.WithQueries(map[string]string{
		"id":        "1",
		"completed": "false",
	})

	var todos []Todo
	err := client.GetUnmarshal(context.Background(), "/todos", &todos, query)
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)
}

func getSingle(client *gopunch.Client) {
	var todo Todo
	err := client.GetUnmarshal(context.Background(), "/todos/1", &todo)
	if err != nil {
		panic(err)
	}

	fmt.Println(todo)
}

func post(client *gopunch.Client) {
	req := &Todo{
		UserID:    6,
		Title:     "xxxx",
		Completed: false,
	}

	var m map[string]interface{}
	err := client.PostUnmarshal(context.Background(), "/todos", req, &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func delete(client *gopunch.Client) {
	var m map[string]interface{}
	err := client.DeleteUnmarshal(context.Background(), "/todos/1", &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}
