package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ihaliti03/pokedex/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

var httpClient = &http.Client{Timeout: 10 * time.Second}

// fetchCached returns the decoded JSON at url, using the cache.
// On a cache hit: unmarshal the cached bytes.
// On a miss: HTTP GET, decode, and cache raw bytes only after success.
func fetchCached[T any](cache *pokecache.Cache, url string) (*T, error) {
	if val, ok := cache.Get(url); ok {
		var v T
		if err := json.Unmarshal(val, &v); err != nil {
			return nil, fmt.Errorf("decode cached %s: %w", url, err)
		}
		return &v, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("build request %s: %w", url, err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(res.Body, 512))
		return nil, fmt.Errorf("GET %s: status %d: %s", url, res.StatusCode, body)
	}

	raw, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", url, err)
	}

	var v T
	if err := json.Unmarshal(raw, &v); err != nil {
		return nil, fmt.Errorf("decode %s: %w", url, err)
	}

	cache.Add(url, raw)
	return &v, nil
}