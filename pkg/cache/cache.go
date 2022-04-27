package cache

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/utkarsh-pro/ucs/pkg/utils"
)

var (
	path string
)

type Key interface {
	GetNamespace() string
	GetQuery() string
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path = filepath.Join(home, ".cache", "ucs")

	utils.PanicIfError(os.MkdirAll(path, 0755))
}

// Get takes the key and returns the value stored in the cached file
func Get(key Key) (string, error) {
	cached := filepath.Join(path, key.GetNamespace(), processPath(key.GetQuery()))

	data, err := os.ReadFile(cached)
	if err != nil {
		return "", fmt.Errorf("failed to read the cached file: %s", err)
	}

	return string(data), nil
}

// Set takes the key and the value and creates a cached file
// with key as the filename and value as the content
func Set(key Key, value string) error {
	cached := filepath.Join(path, key.GetNamespace(), processPath(key.GetQuery()))

	if err := os.MkdirAll(filepath.Join(path, key.GetNamespace()), 0755); err != nil {
		return fmt.Errorf("failed to create the cache directory: %s", err)
	}

	if err := os.WriteFile(cached, []byte(value), 0644); err != nil {
		return fmt.Errorf("failed to write the cached file: %s", err)
	}

	return nil
}

func processPath(path string) string {
	return strings.ReplaceAll(path, "/", "$$")
}
