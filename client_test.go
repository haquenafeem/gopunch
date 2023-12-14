package gopunch

import (
	"context"
	"encoding/json"
	"testing"
)

const BaseURL = "https://jsonplaceholder.typicode.com"

var GetTestCases = []struct {
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

func Test_Get(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range GetTestCases {
		t.Log(testCase.Title)

		resp := client.Get(ctx, testCase.Path)

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

func Test_GetUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range GetTestCases {
		t.Log(testCase.Title)

		var resStruct struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		err := client.GetUnmarshal(ctx, testCase.Path, &resStruct)
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
	}
}

var PostTestCases = []struct {
	Title string
	Path  string
	Data  struct {
		UserID    int    `json:"userId"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
	ExpectedValues struct {
		ID int `json:"id"`
	}
}{
	{
		Title: `Sending Post Request to "https://jsonplaceholder.typicode.com/todos" should return and match expected values`,
		Path:  "/todos",
		Data: struct {
			UserID    int    `json:"userId"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			Title:     "delectus aut autem",
			Completed: false,
		},
		ExpectedValues: struct {
			ID int `json:"id"`
		}{
			ID: 201,
		},
	},
}

func Test_Post(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PostTestCases {
		t.Log(testCase.Title)
		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		resp := client.Post(ctx, testCase.Path, payloadBytes)

		var resStruct struct {
			ID int `json:"id"`
		}

		err = resp.JSONUnmarshal(&resStruct)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}

		resp.Close()
	}
}

func Test_PostUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PostTestCases {
		t.Log(testCase.Title)

		var resStruct struct {
			ID int `json:"id"`
		}

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		err = client.PostUnmarshal(ctx, testCase.Path, payloadBytes, &resStruct)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}
	}
}
