package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/remnikmart/pokedexcli/internal/pokecache"
)

// URLs
func FormUrl(sourceUrl string, path string, params map[string]string, stripOldParams bool) (string, error) {
	u, err := url.Parse(sourceUrl)
	if err != nil {
		return "", err
	}
	if stripOldParams {
		u.RawQuery = ""
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	res, errRes := url.JoinPath(u.String(), path)
	if errRes != nil {
		return "", errRes
	}
	return res, nil
}

// Pages
func GetPage[T any](url string, cache *pokecache.Cache) (T, error) {
	var res T
	body, exists := cache.Get(url)
	if !exists {
		req, errReq := http.NewRequest("GET", url, nil)
		if errReq != nil {
			return res, fmt.Errorf("creating GET request: %w", errReq)
		}
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, errResp := client.Do(req)
		if errResp != nil {
			return res, fmt.Errorf("getting response: %w", errResp)
		}
		defer resp.Body.Close()
		var errBody error
		body, errBody = io.ReadAll(resp.Body)
		if errBody != nil {
			return res, fmt.Errorf("extracting body: %w", errBody)
		}
		cache.Add(url, body)
	}
	err := json.Unmarshal(body, &res)
	if err != nil {
		return res, fmt.Errorf("parsing json: %w", err)
	}
	return res, nil

}
