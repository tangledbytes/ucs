package utils

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

// HTTPGet takes a host and headers and returns a response
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

// LogIfError logs the error if it's not nil
func LogIfError(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

// PanicIfError panics if the error is not nil
func PanicIfError(err error) {
	if err != nil {
		logrus.Panic(err)
	}
}
