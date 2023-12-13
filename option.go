package gopunch

import (
	"net/http"
)

type Option func(req *http.Request)

func WithQueries(queries map[string]string) Option {
	return func(req *http.Request) {
		query := req.URL.Query()
		for key, value := range queries {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}
}

func WithHeaders(headers map[string]string) Option {
	return func(req *http.Request) {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
}
