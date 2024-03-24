package turso

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	httpClient *http.Client
	token      string
	baseurl    *url.URL
}

func NewClient(client *http.Client, token, URL string) (*Client, error) {

	var baseURL = &url.URL{
		Scheme: "https",
		Host:   "api.turso.tech",
		Path:   "/",
	}
	if len(URL) > 0 {
		parsedURL, err := url.Parse(URL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL (%q): %w", URL, err)
		}
		baseURL = parsedURL
	}
	return &Client{
		httpClient: client,
		token:      token,
		baseurl:    baseURL,
	}, nil
}

func (c *Client) createRequest(ctx context.Context, method, path string, reqBody interface{}) (*http.Request, error) {
	var reqBodyReader io.Reader
	if reqBody != nil {
		data, err := json.Marshal(reqBody)
		if err != nil {
			return nil, fmt.Errorf("failed to serialize %+v: %w", reqBody, err)
		}
		reqBodyReader = bytes.NewBuffer(data)
	}
	endpoint := c.baseurl.ResolveReference(&url.URL{Path: path})

	req, err := http.NewRequestWithContext(ctx, method, endpoint.String(), reqBodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.token)

	return req, nil
}

func (c *Client) sendRequest(req *http.Request, resBody interface{}) (*http.Response, error) {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if !ok(res.StatusCode) {
		var errRes errorResponse
		err = json.Unmarshal(body, &errRes)
		if err != nil {
			return res, fmt.Errorf("failed to parse response body, status code %d: %w", res.StatusCode, err)
		}
		// if errRes.StatusCode == "" {
		// 	errRes.StatusCode = fmt.Sprint("%d", res.StatusCode)
		// }
		return res, &errRes
	}
	if resBody != nil {
		err = json.Unmarshal(body, resBody)
		if err != nil {
			return nil, fmt.Errorf("failed to parse response body: %w", err)
		}
	}
	return res, nil
}

func (c *Client) do(ctx context.Context, method, path string, reqBody interface{}, resBody interface{}) (*http.Response, error) {
	req, err := c.createRequest(ctx, method, path, reqBody)
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req, resBody)
}

func ok(code int) bool {
	return code >= 200 && code < 300
}
