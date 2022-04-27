package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/utkarsh-pro/ucs/pkg/chtsh"
)

func init() {
	// Setup logger
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout())
}

func main() {
	args := strings.Join(os.Args[1:], " ")

	res, err := chtsh.Get(args)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(res)
}
