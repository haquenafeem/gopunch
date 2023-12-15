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
