package gopunch

import (
	"net/http"
	"testing"
)

var QueryTestCases = []struct {
	Title           string
	GivenQueries    map[string]string
	ExpectedQueries map[string]string
}{
	{
		Title: "",
		GivenQueries: map[string]string{
			"x": "a",
			"y": "b",
		},
		ExpectedQueries: map[string]string{
			"x": "a",
			"y": "b",
		},
	},
	{
		Title: "",
		GivenQueries: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		ExpectedQueries: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	},
}

func Test_WithQueries(t *testing.T) {
	for _, testCase := range QueryTestCases {
		req, err := http.NewRequest("", "", nil)
		if err != nil {
			panic(err)
		}

		withQueris := WithQueries(testCase.GivenQueries)
		withQueris(req)

		for key, expected := range testCase.ExpectedQueries {
			value := req.URL.Query().Get(key)
			if value != expected {
				t.Log("value missmatch")
				t.Fail()
			}
		}
	}
}
