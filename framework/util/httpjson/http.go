package httpjson

import (
	"encoding/json"
	"errors"
	"fmt"
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

	if resp.StatusCode > 299 {
		return *new(T), errors.New(fmt.Sprintf("get to %s failed with %d code", url, resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return *new(T), err
	}
	d := new(T)
	err = json.Unmarshal(body, &d)
	if err != nil {
		return *new(T), err
	}

	if d == nil {
		return *new(T), errors.New("failed to create T")
	}

	return *d, nil
}
