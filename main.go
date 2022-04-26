package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/utkarsh-pro/ucs/pkg/chtsh"
)

func main() {
	args := strings.Join(os.Args[1:], " ")

	res, err := chtsh.Get(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)
}
