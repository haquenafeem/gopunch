package gopunch

import (
	"context"
	"encoding/json"
	"io"
	"testing"
)

var WithUnmarshalTestCases = []struct {
	Title          string
	Path           string
	ExpectedValues struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
}{
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/1" should return and match expected values`,
		Path:  "/todos/1",
		ExpectedValues: struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			ID:        1,
			Title:     "delectus aut autem",
			Completed: false,
		},
	},
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/2" should return and match expected values`,
		Path:  "/todos/2",
		ExpectedValues: struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			ID:        2,
			Title:     "quis ut nam facilis et officia qui",
			Completed: false,
		},
	},
}

func Test_WithUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range WithUnmarshalTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		resp := client.Get(ctx, testCase.Path, opt)

		var resStruct struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		fn := func(reader io.Reader) error {
			return json.NewDecoder(reader).Decode(&resStruct)
		}

		err := resp.WithUnmarshal(fn)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.UserID != testCase.ExpectedValues.UserID {
			t.Fail()
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}

		if resStruct.Title != testCase.ExpectedValues.Title {
			t.Fail()
		}

		if resStruct.Completed != testCase.ExpectedValues.Completed {
			t.Fail()
		}

		resp.Close()
	}
}

var JSONUnmarshalTestCases = []struct {
	Title          string
	Path           string
	ExpectedValues struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
}{
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/1" should return and match expected values`,
		Path:  "/todos/1",
		ExpectedValues: struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			ID:        1,
			Title:     "delectus aut autem",
			Completed: false,
		},
	},
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/2" should return and match expected values`,
		Path:  "/todos/2",
		ExpectedValues: struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			ID:        2,
			Title:     "quis ut nam facilis et officia qui",
			Completed: false,
		},
	},
}

func Test_JSONUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range JSONUnmarshalTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		resp := client.Get(ctx, testCase.Path, opt)

		var resStruct struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		err := resp.JSONUnmarshal(&resStruct)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.UserID != testCase.ExpectedValues.UserID {
			t.Fail()
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}

		if resStruct.Title != testCase.ExpectedValues.Title {
			t.Fail()
		}

		if resStruct.Completed != testCase.ExpectedValues.Completed {
			t.Fail()
		}

		resp.Close()
	}
}

var StringUnmarshalTestCases = []struct {
	Title         string
	Path          string
	ExpectedValue string
}{
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/1" should return and match expected value`,
		Path:  "/todos/1",
		ExpectedValue: `{
 "userId": 1,
 "id": 1,
 "title": "delectus aut autem",
 "completed": false
}`,
	},
}

func Test_StringUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range StringUnmarshalTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		resp := client.Get(ctx, testCase.Path, opt)

		dest := ""

		err := resp.StringUnmarshal(&dest)
		if err != nil {
			t.Fatal(err)
		}

		var d, e map[string]interface{}
		err = json.Unmarshal([]byte(dest), &d)
		if err != nil {
			t.Fatal(err)
		}

		err = json.Unmarshal([]byte(testCase.ExpectedValue), &e)
		if err != nil {
			t.Fatal(err)
		}

		if len(d) != len(e) {
			t.Fail()
		}

		for key, value := range d {
			if e[key] != value {
				t.Fail()
			}
		}

		resp.Close()
	}
}
