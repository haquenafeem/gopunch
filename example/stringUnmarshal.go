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
