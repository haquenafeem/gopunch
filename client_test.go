package gopunch

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"
	"time"
)

const BaseURL = "https://jsonplaceholder.typicode.com"

func Test_NewWithTimeOut(t *testing.T) {
	client := NewWithTimeOut(BaseURL, time.Millisecond)
	ctx := context.Background()
	expectedErrString := `Get "https://jsonplaceholder.typicode.com/todos/1": context deadline exceeded (Client.Timeout exceeded while awaiting headers)`

	t.Log(`Creating client with "NewWithTimeOut" and providing very little timeout; the request should fail and return
context deadline exceeded releted error`)

	var m map[string]interface{}
	err := client.GetUnmarshal(ctx, "/todos/1", &m)
	if err == nil {
		t.Fail()
	}

	if err.Error() != expectedErrString {
		t.Fail()
	}
}

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

func Test_Get_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Get(nil, "/todos/1")
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
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

func Test_Post_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Post(nil, "/todos/1", nil)
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
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
	Title string
	Path  string
}{
	{
		Title: `Sending Delete Request to "https://jsonplaceholder.typicode.com/todos/1" would return empty response; 
but without error and with status code ok`,
		Path: "/todos/1",
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

		if resp.err != nil {
			t.Fatal(resp.err)
		}

		if resp.httpResponse.StatusCode != http.StatusOK {
			t.Fail()
		}

		resp.Close()
	}
}

func Test_Delete_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Delete(nil, "/todos/1")
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
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

func Test_Put_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Put(nil, "/todos/1", nil)
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
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

var PatchTestCases = []struct {
	Title string
	Path  string
	Data  struct {
		Title string `json:"title"`
	}
	ExpectedValues struct {
		UserID    int    `json:"userId"`
		ID        int    `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
}{
	{
		Title: `Sending Patch Request for title to "https://jsonplaceholder.typicode.com/todos/1" should return and match expected values`,
		Path:  "/todos/1",
		Data: struct {
			Title string `json:"title"`
		}{
			Title: "title updated",
		},
		ExpectedValues: struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}{
			UserID:    1,
			ID:        1,
			Title:     "title updated",
			Completed: false,
		},
	},
}

func Test_Patch(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PatchTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		resp := client.Patch(ctx, testCase.Path, payloadBytes, opt)

		var resStruct struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}

		err = resp.JSONUnmarshal(&resStruct)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}

		if resStruct.UserID != testCase.ExpectedValues.UserID {
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

func Test_Patch_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Patch(nil, "/todos/1", nil)
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
	}
}

func Test_PatchUnmarshal(t *testing.T) {
	client := New(BaseURL)
	ctx := context.Background()
	for _, testCase := range PatchTestCases {
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

		payloadBytes, err := json.Marshal(&testCase.Data)
		if err != nil {
			t.Fatal(err)
		}

		err = client.PatchUnmarshal(ctx, testCase.Path, payloadBytes, &resStruct, opt)
		if err != nil {
			t.Fatal(err)
		}

		if resStruct.ID != testCase.ExpectedValues.ID {
			t.Fail()
		}

		if resStruct.UserID != testCase.ExpectedValues.UserID {
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

var CustomTestCases = []struct {
	Title  string
	Values struct {
		Path    string
		Context context.Context
		Method  string
		Payload map[string]interface{}
	}
	ExpectedErrContains string
	CheckErr            bool
	CheckValues         bool
	ExpectedValues      map[string]interface{}
}{
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/1" providing context should not return err; and 
expected values should match`,
		Values: struct {
			Path    string
			Context context.Context
			Method  string
			Payload map[string]interface{}
		}{
			Path:    "/todos/1",
			Context: context.Background(),
			Method:  http.MethodGet,
			Payload: map[string]interface{}{},
		},
		ExpectedErrContains: "",
		CheckErr:            false,
		ExpectedValues: map[string]interface{}{
			"userId":    float64(1),
			"id":        float64(1),
			"title":     `delectus aut autem`,
			"completed": false,
		},
		CheckValues: true,
	},
	{
		Title: `Sending Get Request to "https://jsonplaceholder.typicode.com/todos/1" providing nil context should return err`,
		Values: struct {
			Path    string
			Context context.Context
			Method  string
			Payload map[string]interface{}
		}{
			Path:    "/todos/1",
			Context: nil,
			Method:  http.MethodGet,
			Payload: map[string]interface{}{},
		},
		ExpectedErrContains: "nil Context",
		CheckErr:            true,
		ExpectedValues:      map[string]interface{}{},
		CheckValues:         false,
	},
}

func Test_Custom(t *testing.T) {
	client := New(BaseURL)
	for _, testCase := range CustomTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		payloadBytes, err := json.Marshal(&testCase.Values.Payload)
		if err != nil {
			t.Fatal(err)
		}

		resp := client.Custom(
			testCase.Values.Context,
			testCase.Values.Method,
			testCase.Values.Path,
			payloadBytes,
			opt)

		if testCase.CheckErr {
			if resp.err == nil {
				t.Fail()
			}

			if !strings.Contains(resp.err.Error(), testCase.ExpectedErrContains) {
				t.Fail()
			}
		}

		if testCase.CheckValues {
			var m map[string]interface{}
			err := resp.JSONUnmarshal(&m)
			if err != nil {
				t.Fatal(err)
			}

			if len(m) != len(testCase.ExpectedValues) {
				t.Fail()
			}

			for key, value := range m {
				if testCase.ExpectedValues[key] != value {
					t.Fail()
				}
			}
		}
	}
}

func Test_Custom_Without_Context(t *testing.T) {
	client := New(BaseURL)
	t.Log("Given context provided as nil; request would fail with error")
	response := client.Custom(nil, http.MethodGet, "/todos/1", nil)
	if response.err == nil {
		t.Fail()
	}

	if !strings.Contains(response.err.Error(), "nil Context") {
		t.Fail()
	}
}

func Test_CustomUnmarshal(t *testing.T) {
	client := New(BaseURL)
	for _, testCase := range CustomTestCases {
		t.Log(testCase.Title)
		opt := WithHeaders(map[string]string{
			"Content-Type": "application/json",
		})

		payloadBytes, err := json.Marshal(&testCase.Values.Payload)
		if err != nil {
			t.Fatal(err)
		}

		var m map[string]interface{}

		err = client.CustomUnmarshal(
			testCase.Values.Context,
			testCase.Values.Method,
			testCase.Values.Path,
			payloadBytes,
			&m,
			opt)

		if testCase.CheckErr {
			if err == nil {
				t.Fail()
			}

			if !strings.Contains(err.Error(), testCase.ExpectedErrContains) {
				t.Fail()
			}
		}

		if testCase.CheckValues {
			if len(m) != len(testCase.ExpectedValues) {
				t.Fail()
			}

			for key, value := range m {
				if testCase.ExpectedValues[key] != value {
					t.Fail()
				}
			}
		}
	}
}
