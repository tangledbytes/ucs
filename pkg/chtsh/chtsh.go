package chtsh

import (
	"fmt"
	"io"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/utkarsh-pro/ucs/pkg/utils"
)

const (
	host = "https://cht.sh/"
)

// Get takes a query and returns the result
//
// The query is directly passed to cht.sh
//
// NOTE: Offline mode hasn't been implemented yet
func Get(arg string) (string, error) {
	data, err := getOnline(arg)
	if err == nil {
		return data, nil
	}

	logrus.Error("failed to reach the source, trying offline mode: ", err)

	return getOffline(arg)
}

// getOnline takes a query and returns the result from cht.sh
func getOnline(arg string) (string, error) {
	endpoint := host + prepareQuery(arg)

	resp, err := utils.HTTPGet(endpoint, map[string]string{"User-Agent": "curl/7.54.0"})
	if err != nil {
		return "", fmt.Errorf("failed to reach the source: %s", err)
	}

	if resp.StatusCode == 200 {
		result, err := io.ReadAll(resp.Body)
		defer utils.LogIfError(resp.Body.Close())
		if err != nil {
			return "", fmt.Errorf("failed to read the response: %s", err)
		}

		return string(result), nil
	}

	return "", fmt.Errorf("failed to reach the source: status code => %s", resp.Status)
}

// getOffline takes a query and returns the result from local cached copy
//
// NOTE: This function is not implemented yet
func getOffline(arg string) (string, error) {
	return "", fmt.Errorf("offline mode not implemented")
}

// prepareQuery takes a query and transforms it into a valid cht.sh query
func prepareQuery(query string) string {
	parsed := strings.SplitN(query, " ", 2)
	if len(parsed) == 1 {
		return parsed[0]
	}

	return parsed[0] + "/" + strings.ReplaceAll(parsed[1], " ", "+")
}
