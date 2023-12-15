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

func Test_GetUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range GetTestCases {
		t.Log(testCase.Title)

		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		var resStruct struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		err := client.GetUnmarshal(ctx, testCase.Path, &resStruct, opt)
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
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		resp := client.Post(ctx, testCase.Path, payloadBytes, opt)

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
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		var resStruct struct {
			ID int `json:"id"`
		}

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		err = client.PostUnmarshal(ctx, testCase.Path, payloadBytes, &resStruct, opt)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}
	}
}

var DeleteTestCases = []struct {
	Title          string
	Path           string
	ExpectedString string
}{
	{
		Title:          `Sending Delete Request to "https://jsonplaceholder.typicode.com/todos/1" would return empty json "{}"`,
		Path:           "/todos/1",
		ExpectedString: "{}",
	},
}

func Test_Delete(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range DeleteTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		resp := client.Delete(ctx, testCase.Path, opt)

		str := ""
		err := resp.StringUnmarshal(&str)
		if err != nil {
			t.Fatal(err)
		}

		if str != testCase.ExpectedString {
			t.Fail()
		}

		resp.Close()
	}
}

var DeleteUnmarshalTestCases = []struct {
	Title string
	Path  string
}{
	{
		Title: `Sending Delete Request to "https://jsonplaceholder.typicode.com/todos/1" and trying to unmarshal to a
		map would return empty map without error`,
		Path: "/todos/1",
	},
}

func Test_DeleteUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range DeleteUnmarshalTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		var m map[string]interface{}
		err := client.DeleteUnmarshal(ctx, testCase.Path, &m, opt)

		if len(m) != 0 {
			t.Fail()
		}

		if err != nil {
			t.Fail()
		}
	}
}

var PutTestCases = []struct {
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
		Title: `Sending Put Request to "https://jsonplaceholder.typicode.com/todos/13" should return and match expected values`,
		Path:  "/todos/13",
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
			ID: 13,
		},
	},
}

func Test_Put(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PutTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		resp := client.Put(ctx, testCase.Path, payloadBytes, opt)

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

func Test_PutUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PutTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		var resStruct struct {
			ID int `json:"id"`
		}

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		err = client.PutUnmarshal(ctx, testCase.Path, payloadBytes, &resStruct, opt)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}
	}
}
