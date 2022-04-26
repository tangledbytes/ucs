package chtsh

import (
	"fmt"
	"io"
	"strings"

	"github.com/utkarsh-pro/ucs/pkg/utils"
)

const (
	host = "cht.sh"
)

func Get(arg string) (string, error) {
	data, err := getOnline(arg)
	if err == nil {
		return data, nil
	}
	println(err.Error())

	return getOffline(arg)
}

func getOnline(arg string) (string, error) {
	endpoint := "http://" + host + "/" + prepareQuery(arg)
	resp, err := utils.HTTPGet(endpoint, map[string]string{"User-Agent": "curl/7.54.0"})
	if err != nil {
		return "", fmt.Errorf("failed to reach the source: %s", err)
	}

	if resp.StatusCode == 200 {
		result, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return "", fmt.Errorf("failed to read the response: %s", err)
		}

		return string(result), nil
	}

	return "", fmt.Errorf("failed to reach the source: status code => %s", resp.Status)
}

func getOffline(arg string) (string, error) {
	return "", fmt.Errorf("offline mode not implemented")
}

func prepareQuery(query string) string {
	parsed := strings.SplitN(query, " ", 2)
	if len(parsed) == 1 {
		return parsed[0]
	}

	return parsed[0] + "/" + strings.ReplaceAll(parsed[1], " ", "+")
}
