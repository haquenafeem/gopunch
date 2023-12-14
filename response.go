package gopunch

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Response struct {
	httpResponse *http.Response
	err          error
}

func NewResponse(httpResponse *http.Response, err error) *Response {
	return &Response{
		httpResponse: httpResponse,
		err:          err,
	}
}

func (r *Response) Close() error {
	return r.httpResponse.Body.Close()
}

func (r *Response) WithUnmarshal(fn func(reader io.Reader) error) error {
	if r.err != nil {
		return r.err
	}

	return fn(r.httpResponse.Body)
}

func (r *Response) JSONUnmarshal(dest interface{}) error {
	if r.err != nil {
		return r.err
	}

	fn := func(reader io.Reader) error {
		return json.NewDecoder(reader).Decode(dest)
	}

	return r.WithUnmarshal(fn)
}

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
