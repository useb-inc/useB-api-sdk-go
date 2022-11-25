package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Base struct {
	Host string
}

func callAPIMethod[T, R interface{}](base *Base, auth *Auth, method, path string) func(*T) (*R, error) {
	url := makeUrl([]string{base.Host, path})
	return func(param *T) (*R, error) {
		reqBody, err := json.Marshal(param)
		if err != nil {
			return nil, err
		}

		req, err := http.NewRequest(method, url, bytes.NewReader(reqBody))
		if err != nil {
			return nil, err
		}

		req.Header.Set(HeaderKeyContentType, HeaderValueContentTypeJson)
		req.Header.Set(HeaderKeyUserAgent, HeaderValueUserAgent)

		if auth != nil {
			token, err := auth.getToken()
			if err != nil {
				return nil, err
			}
			req.Header.Set(HeaderKeyAuthorization, token)
		}

		// send request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("status code: %d, body: %s", res.StatusCode, string(resBody))
		}

		var response R
		json.Unmarshal(resBody, &response)

		return &response, nil
	}
}

func makeUrl(paths []string) string {
	// loop through paths and remove leading and trailing slashes
	for i, path := range paths {
		paths[i] = strings.Trim(path, "/")
	}
	return strings.Join(paths, "/")
}
