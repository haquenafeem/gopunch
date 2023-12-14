package gopunch

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// Client
// has baseURL and *http.Client
type Client struct {
	baseUrl    string
	httpClient *http.Client
}

// New
// returns a new *gopunch.Client
func New(baseUrl string) *Client {
	return &Client{
		baseUrl:    baseUrl,
		httpClient: &http.Client{},
	}
}

// NewWithTimeOut
// returns a new *gopunch.Client
// adds time duration to http.Client for requests to complete
func NewWithTimeOut(baseUrl string, timeout time.Duration) *Client {
	return &Client{
		baseUrl: baseUrl,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get
// takes context, endpoint and option functions
// returns *http.Response and error
func (c *Client) Get(ctx context.Context, endPoint string, opts ...Option) (*http.Response, error) {
	completeUrl := c.baseUrl + endPoint

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, completeUrl, nil)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// GetUnmarshal
// takes context, endpoint, pointer to which the response will be unmarshalled and option functions
// returns only error
func (c *Client) GetUnmarshal(ctx context.Context, endPoint string, dest interface{}, opts ...Option) error {
	resp, err := c.Get(ctx, endPoint, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}

// Post
// takes context, endpoint, payload pointer and option functions
// returns *http.Response and error
func (c *Client) Post(ctx context.Context, endPoint string, payload interface{}, opts ...Option) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	completeUrl := c.baseUrl + endPoint

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, completeUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// PostUnmarshal
// takes context, endpoint, payload pointer, pointer to which the response will be unmarshalled and option functions
// returns only error
func (c *Client) PostUnmarshal(ctx context.Context, endPoint string, payload interface{}, dest interface{}, opts ...Option) error {
	resp, err := c.Post(ctx, endPoint, payload, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}

// Delete
// takes context, endpoint and option functions
// returns *http.Response and error
func (c *Client) Delete(ctx context.Context, endPoint string, opts ...Option) (*http.Response, error) {
	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, completeUrl, nil)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// DeleteUnmarshal
// takes context, endpoint, pointer to which the response will be unmarshalled and option functions
// returns only error
func (c *Client) DeleteUnmarshal(ctx context.Context, endPoint string, dest interface{}, opts ...Option) error {
	resp, err := c.Delete(ctx, endPoint, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}

// Put
// takes context, endpoint, payload pointer and option functions
// returns *http.Response and error
func (c *Client) Put(ctx context.Context, endPoint string, payload interface{}, opts ...Option) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, completeUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// PutUnmarshal
// takes context, endpoint, payload pointer, pointer to which the response will be unmarshalled and option functions
// returns only error
func (c *Client) PutUnmarshal(ctx context.Context, endPoint string, payload interface{}, dest interface{}, opts ...Option) error {
	resp, err := c.Put(ctx, endPoint, payload, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}

// Patch
// takes context, endpoint, payload pointer and option functions
// returns *http.Response and error
func (c *Client) Patch(ctx context.Context, endPoint string, payload interface{}, opts ...Option) (*http.Response, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	completeUrl := c.baseUrl + endPoint
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, completeUrl, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		opt(req)
	}

	return c.httpClient.Do(req)
}

// PatchUnmarshal
// takes context, endpoint, payload pointer, pointer to which the response will be unmarshalled and option functions
// returns only error
func (c *Client) PatchUnmarshal(ctx context.Context, endPoint string, payload interface{}, dest interface{}, opts ...Option) error {
	resp, err := c.Patch(ctx, endPoint, payload, opts...)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}
