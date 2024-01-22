package httpjson

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

var client *http.Client
var once sync.Once

func getClient() *http.Client {
	once.Do(func() {
		tr := &http.Transport{
			MaxIdleConns:          10,
			IdleConnTimeout:       15 * time.Second,
			ResponseHeaderTimeout: 15 * time.Second,
			DisableKeepAlives:     false,
		}
		httpClient := &http.Client{
			Transport: tr,
			// do not follow redirects
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
		client = httpClient
	})

	return client
}

func Get[T any](url string) (T, error) {
	resp, err := getClient().Get(url)
	if err != nil {
		return *new(T), err
	}

	defer func() {
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(T), err
	}
	d := new(T)
	err = json.Unmarshal(body, &d)
	if err != nil {
		return *new(T), err
	}
	return *d, nil
}
