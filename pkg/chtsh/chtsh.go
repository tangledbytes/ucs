package chtsh

import (
	"fmt"
	"io"

	"github.com/utkarsh-pro/ucs/pkg/cache"
	"github.com/utkarsh-pro/ucs/pkg/utils"
)

const (
	host = "https://cht.sh/"
)

type ChtShtQuery interface {
	GetNamespace() string
	GetQuery() string
}

type ChtSht struct {
	query ChtShtQuery
}

func New(query ChtShtQuery) ChtSht {
	return ChtSht{query: query}
}

func (c ChtSht) Find() (string, error) {
	data, err := c.findOnline()
	utils.LogIfError(err)
	if err != nil {
		return c.findOffline()
	}

	return data, nil
}

func (c ChtSht) findOnline() (string, error) {
	endpoint := host + c.query.GetNamespace() + "/" + c.query.GetQuery()

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

		utils.LogIfError(cache.Set(c.query, string(result)))
		return string(result), nil
	}

	return "", fmt.Errorf("failed to reach the source: status code => %s", resp.Status)
}

func (c ChtSht) findOffline() (string, error) {
	data, err := cache.Get(c.query)
	if err != nil {
		return "", fmt.Errorf("failed to read the cached data: %s", err)
	}

	return data, nil
}
