// Generic GET function with caching implemented to maintain DRY code
// for all API and caching requests
package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) get(url string) ([]byte, error) {
	if val, ok := c.cache.Get(url); ok {
		return val, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode > 399 {
		return nil, fmt.Errorf("bad status code: %v", res.StatusCode)
	}

	rawData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, rawData)
	return rawData, nil
}
