package utils

import "net/http"

// HTTPGet takes a host and headers and returns a response
//
// If headers is nil then no headers are sent
func HTTPGet(host string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return http.DefaultClient.Do(req)
}
