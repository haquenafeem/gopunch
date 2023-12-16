# GoPunch examples

Use `example.go` file to select which example you want to run.

## JSON Place Holder example
*typeOfRequest.go*
```go
package main

type TypeOFRequest int

const (
	GetSingle TypeOFRequest = iota
	GetAll
	GetAllWithQueries
	Post
	Delete
	Put
	Patch
)
```
*jsonPlaceHolder.go*
```go
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

```

## File Download With Custom Unmarshal
```go
package main

import (
	"context"
	"io"
	"os"

	"github.com/haquenafeem/gopunch"
)

func FileDownload() {
	client := gopunch.New("https://images.unsplash.com/photo-1481349518771-20055b2a7b24?q=80&w=1000&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8NHx8cmFuZG9tfGVufDB8fDB8fHww")
	ctx := context.Background()

	file, err := os.Create("./etc/abc.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fn := func(reader io.Reader) error {
		_, err := io.Copy(file, reader)

		return err
	}

	resp := client.Get(ctx, "")
	defer resp.Close()

	err = resp.WithUnmarshal(fn)
	if err != nil {
		panic(err)
	}
}

```

## String Unmarshal With Custom Request
```go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/haquenafeem/gopunch"
)

func StringUnmarshal() {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	ctx := context.Background()

	resp := client.Custom(ctx, http.MethodGet, "/todos/1", nil)
	defer resp.Close()
	dest := ""

	fn := func(reader io.Reader) error {
		bytes, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		dest = string(bytes)

		return nil
	}

	err := resp.WithUnmarshal(fn)
	if err != nil {
		panic(err)
	}

	fmt.Println(dest)
}

```