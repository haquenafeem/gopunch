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
		Title: "Given x,y = a,b as queries; when fetching queries x,y key should return a,b",
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
		Title: "Given key1,key2 = value1,value2 as queries; when fetching queries key1,key2 key should return value1,value2",
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

var HeaderTestCases = []struct {
	Title           string
	GivenHeaders    map[string]string
	ExpectedHeaders map[string]string
}{
	{
		Title: "Given x,y = a,b as headers; when fetching headers x,y key should return a,b",
		GivenHeaders: map[string]string{
			"x": "a",
			"y": "b",
		},
		ExpectedHeaders: map[string]string{
			"x": "a",
			"y": "b",
		},
	},
	{
		Title: "Given key1,key2 = value1,value2 as headers; when fetching headers key1,key2 key should return value1,value2",
		GivenHeaders: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		ExpectedHeaders: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
	},
}

func Test_WithHeaders(t *testing.T) {
	for _, testCase := range HeaderTestCases {
		req, err := http.NewRequest("", "", nil)
		if err != nil {
			panic(err)
		}

		WithHeaders := WithHeaders(testCase.GivenHeaders)
		WithHeaders(req)

		for key, expected := range testCase.ExpectedHeaders {
			value := req.Header.Get(key)
			if value != expected {
				t.Log("value missmatch")
				t.Fail()
			}
		}
	}
}
