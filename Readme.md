<p align="center">
  <a href="https://github.com/haquenafeem/gopunch">
    <img alt="gopunch" src="https://github.com/haquenafeem/gopunch/blob/main/assets/banner.png" width="250">
  </a>
</p>

<p align="center">
  A simple http client for <a href="https://golang.org/">Golang</a>
</p>

# GoPunch
`gopunch` is a simple golang package to make http calls.

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
	var todos []Todo
	err := client.GetUnmarshal(context.Background(), "/todos", &todos)
	if err != nil {
		panic(err)
	}

	fmt.Println(todos)
}

```
<p>
  Complete Example <a href="https://github.com/haquenafeem/gopunch/tree/main/example">Here</a>
</p>