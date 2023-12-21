package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/haquenafeem/gopunch"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func RunAllJSONPlaceHolder() {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	ctx := context.Background()

	getAll(ctx, client)
	getSingle(ctx, client)
	delete(ctx, client)
	post(ctx, client)
	put(ctx, client)
	patch(ctx, client)
}

func RunJSONPlaceHolder(reqType TypeOFRequest) {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	ctx := context.Background()

	switch reqType {
	case GetAll:
		getAll(ctx, client)
		return
	case GetAllWithQueries:
		getAllWithQueries(ctx, client)
		return
	case GetSingle:
		getSingle(ctx, client)
		return
	case Post:
		post(ctx, client)
		return
	case Delete:
		delete(ctx, client)
		return
	case Put:
		put(ctx, client)
		return
	case Patch:
		patch(ctx, client)
	}
}

func getAll(ctx context.Context, client *gopunch.Client) {
	fmt.Println("getAll--->")
	var todos []Todo
	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	err := client.GetUnmarshal(ctx, "/todos", &todos, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)
}

func getAllWithQueries(ctx context.Context, client *gopunch.Client) {
	fmt.Println("getAllWithQueries--->")
	var todos []Todo
	opt1 := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	opt2 := gopunch.WithQueries(map[string]string{
		"userId":    "1",
		"completed": "false",
	})

	err := client.GetUnmarshal(ctx, "/todos", &todos, opt1, opt2)
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)
}

func getSingle(ctx context.Context, client *gopunch.Client) {
	fmt.Println("getSingle--->")
	var todo Todo

	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	err := client.GetUnmarshal(ctx, "/todos/1", &todo, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(todo)
}

func post(ctx context.Context, client *gopunch.Client) {
	fmt.Println("post--->")
	req := &Todo{
		UserID:    6,
		Title:     "xxxx",
		Completed: false,
	}

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	var m map[string]interface{}
	err = client.PostUnmarshal(ctx, "/todos", payloadBytes, &m, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func delete(ctx context.Context, client *gopunch.Client) {
	fmt.Println("delete--->")
	var m map[string]interface{}
	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	err := client.DeleteUnmarshal(ctx, "/todos/1", &m, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func put(ctx context.Context, client *gopunch.Client) {
	fmt.Println("put--->")
	req := &Todo{
		UserID:    6,
		Title:     "xxxx",
		Completed: false,
	}

	payloadBytes, err := json.Marshal(req)
	if err != nil {
		panic(err)
	}

	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	var m map[string]interface{}
	err = client.PutUnmarshal(ctx, "/todos/13", payloadBytes, &m, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}

func patch(ctx context.Context, client *gopunch.Client) {
	fmt.Println("patch--->")
	req := struct {
		Title string `json:"title"`
	}{
		Title: "abc",
	}

	payloadBytes, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	opt := gopunch.WithHeaders(map[string]string{
		"Content-Type": "application/json",
	})

	err = client.PatchUnmarshal(ctx, "/todos/1", payloadBytes, &m, opt)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)
}
