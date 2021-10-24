package github

import "net/http"

type httpClient struct {
	client      http.Client
	bearerToken string
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+c.bearerToken)
	return c.client.Do(req)
}
