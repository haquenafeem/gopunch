package main

import (
	"context"
	"fmt"
	"io"

	"github.com/haquenafeem/gopunch"
)

// func StringUnmarshal(dest *string) error {

// 	dest := ""

// 	fn := func(reader io.Reader) error {
// 		bytes, err := io.ReadAll(reader)
// 		if err != nil {
// 			return err
// 		}

// 		if dest == nil {
// 			return errors.New("nil pointer")
// 		}

// 		*dest = string(bytes)

// 		return nil
// 	}

// 	return r.WithUnmarshal(fn)
// }

func StringUnmarshal() {
	client := gopunch.New("https://jsonplaceholder.typicode.com")
	ctx := context.Background()

	resp := client.Get(ctx, "/todos/1")
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
