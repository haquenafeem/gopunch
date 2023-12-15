package gopunch

import (
	"bytes"
	"context"
	"net/http"
	"time"
)

// Client
//
//	has baseURL and *http.Client
type Client struct {
	baseUrl    string
	httpClient *http.Client
}

// New
//
//	returns a new *gopunch.Client
func New(baseUrl string) *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
	}
}

// NewWithTimeOut
//
//	returns a new *gopunch.Client
//	adds time duration to http.Client for requests to complete
func NewWithTimeOut(baseUrl string, timeout time.Duration) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get
//
//	takes context, endpoint and option functions
//	returns *Response
func (c *Client) Get(ctx context.Context, endPoint string, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, completeUrl, nil)
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// GetUnmarshal
//
//	takes context, endpoint, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) GetUnmarshal(ctx context.Context, endPoint string, dest interface{}, opts ...Option) error {
	resp := c.Get(ctx, endPoint, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}

// Post
//
//	takes context, endpoint, payload and option functions
//	returns *Response
func (c *Client) Post(ctx context.Context, endPoint string, payload []byte, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, completeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// PostUnmarshal
//
//	takes context, endpoint, payload, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) PostUnmarshal(ctx context.Context, endPoint string, payload []byte, dest interface{}, opts ...Option) error {
	resp := c.Post(ctx, endPoint, payload, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}

// Delete
//
//	takes context, endpoint and option functions
//	returns *Response
func (c *Client) Delete(ctx context.Context, endPoint string, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, completeUrl, nil)
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// DeleteUnmarshal
//
//	takes context, endpoint, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) DeleteUnmarshal(ctx context.Context, endPoint string, dest interface{}, opts ...Option) error {
	resp := c.Delete(ctx, endPoint, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}

// Put
//
//	takes context, endpoint, payload and option functions
//	returns *Response
func (c *Client) Put(ctx context.Context, endPoint string, payload []byte, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, completeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// PutUnmarshal
//
//	takes context, endpoint, payload, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) PutUnmarshal(ctx context.Context, endPoint string, payload []byte, dest interface{}, opts ...Option) error {
	resp := c.Put(ctx, endPoint, payload, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}

// Patch
//
//	takes context, endpoint, payload and option functions
//	returns *Response
func (c *Client) Patch(ctx context.Context, endPoint string, payload []byte, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, completeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// PatchUnmarshal
//
//	takes context, endpoint, payload, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) PatchUnmarshal(ctx context.Context, endPoint string, payload []byte, dest interface{}, opts ...Option) error {
	resp := c.Patch(ctx, endPoint, payload, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}

// Custom
//
//	takes context, method (GET/POST....), endpoint, payload and option functions
//	returns *Response
func (c *Client) Custom(ctx context.Context, method, endPoint string, payload []byte, opts ...Option) *Response {
	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, method, completeUrl, bytes.NewBuffer(payload))
	if err != nil {
		return NewResponse(nil, err)
	}

	for _, opt := range opts {
		opt(req)
	}

	return NewResponse(c.httpClient.Do(req))
}

// CustomUnmarshal
//
//	takes context, method (GET/POST....), endpoint, payload, pointer to which the response will be unmarshalled and option functions
//	returns only error
func (c *Client) CustomUnmarshal(ctx context.Context, method, endPoint string, payload []byte, dest interface{}, opts ...Option) error {
	resp := c.Custom(ctx, method, endPoint, payload, opts...)
	defer resp.Close()

	return resp.JSONUnmarshal(dest)
}
