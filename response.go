package gopunch

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var ErrHttpResponseNil = errors.New("httpResponse is nil")
var ErrHttpResponseBodyNil = errors.New("httpResponse body is nil")

// Response
//
//	includes *http.Response and error
type Response struct {
	httpResponse *http.Response
	err          error
}

// NewResponse
//
//	takes *http.Response,err returns *Response
func NewResponse(httpResponse *http.Response, err error) *Response {
	return &Response{
		httpResponse: httpResponse,
		err:          err,
	}
}

// Close
//
//	closes *httpResponse body
func (r *Response) Close() error {
	if r.err != nil {
		return r.err
	}

	if r.httpResponse == nil {
		return ErrHttpResponseNil
	}

	if r.httpResponse.Body == nil {
		return ErrHttpResponseBodyNil
	}

	return r.httpResponse.Body.Close()
}

// WithUnmarshal
//
//	takes funcfunc(reader io.Reader) error
//	returns error
//	use to create custom unmarshal
func (r *Response) WithUnmarshal(fn func(reader io.Reader) error) error {
	if r.err != nil {
		return r.err
	}

	if r.httpResponse == nil {
		return ErrHttpResponseNil
	}

	if r.httpResponse.Body == nil {
		return ErrHttpResponseBodyNil
	}

	return fn(r.httpResponse.Body)
}

// JSONUnmarshal
//
//	takes pointer to destination
//	returns error
func (r *Response) JSONUnmarshal(dest interface{}) error {
	if r.err != nil {
		return r.err
	}

	fn := func(reader io.Reader) error {
		return json.NewDecoder(reader).Decode(dest)
	}

	return r.WithUnmarshal(fn)
}

// StringUnmarshal
//
//	takes pointer to destination string
//	returns error
func (r *Response) StringUnmarshal(dest *string) error {
	if r.err != nil {
		return r.err
	}

	fn := func(reader io.Reader) error {
		bytes, err := io.ReadAll(reader)
		if err != nil {
			return err
		}

		if dest == nil {
			return errors.New("nil pointer")
		}

		*dest = string(bytes)

		return nil
	}

	return r.WithUnmarshal(fn)
}
