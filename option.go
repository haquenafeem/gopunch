package gopunch

import (
	"net/http"
)

// Option
//
//	this can be used to customize each request
type Option func(req *http.Request)

// WithQueries
//
//	takes a query map and adds it as queries to the request
func WithQueries(queries map[string]string) Option {
	return func(req *http.Request) {
		query := req.URL.Query()
		for key, value := range queries {
			query.Add(key, value)
		}
		req.URL.RawQuery = query.Encode()
	}
}

// WithHeaders
//
//	takes a headers map and adds it as headers to the request
func WithHeaders(headers map[string]string) Option {
	return func(req *http.Request) {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}
}
