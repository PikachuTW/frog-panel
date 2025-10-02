package utils

import (
	"fmt"

	"resty.dev/v3"
)

func FetchData[T any](url string, queryParams ...map[string]string) (T, error) {
	client := resty.New()
	var result T

	req := client.R().SetResult(&result)

	if len(queryParams) > 0 {
		for key, value := range queryParams[0] {
			req.SetQueryParam(key, value)
		}
	}

	resp, err := req.Get(url)

	if err != nil {
		return *new(T), err
	}

	if !resp.IsSuccess() {
		return *new(T), fmt.Errorf("HTTP %d: %s", resp.StatusCode(), resp.Status())
	}

	return result, nil
}
