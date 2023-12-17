<p align="center">
  <a href="https://github.com/haquenafeem/gopunch">
    <img alt="gopunch" src="https://github.com/haquenafeem/gopunch/blob/main/assets/banner.png" width="250">
  </a>
</p>

# GoPunch
`gopunch` is a simple golang package to make http calls. 

![workflow](https://github.com/haquenafeem/gopunch/actions/workflows/go.yml/badge.svg) [![Go Reference](https://pkg.go.dev/badge/github.com/haquenafeem/gopunch.svg)](https://pkg.go.dev/github.com/haquenafeem/gopunch) [![Go Report Card](https://goreportcard.com/badge/github.com/haquenafeem/gopunch)](https://goreportcard.com/report/github.com/haquenafeem/gopunch)

# Features
- Easy To Use
- Use Own Unmarshal Logic
- Create With/Without Timer
- Easy To Use
- Default JSON Oriented
- Request/Respose Modification 
- Examples To Get Your Started
- All Tests/Examples Based On `JSON Place Holder`
- Tests Passing


## Usage
```go
package main

import (
	"context"
	"fmt"

	"github.com/haquenafeem/gopunch"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	ctx := context.Background()

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

```
<p>
  Complete Example <a href="https://github.com/haquenafeem/gopunch/tree/main/example">Here</a>
</p>