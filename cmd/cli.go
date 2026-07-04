package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Cli() (string, int64, error) {
	fmt.Println("Wellcome to ddos-cli")
	fmt.Println("This is a cli tool to do ddos attacks")
	fmt.Println("It will be a simple cli tool to do ddos attacks")

	if len(os.Args) < 3 {
		fmt.Println("Usage: ddos-cli <target> <port>")

		return "", 0, errors.New("invalid arguments: missing target or port")
	}

	endpoint := os.Args[1]

	if strings.TrimSpace(endpoint) == "" {
		return "", 0, errors.New("invalid arguments: missing target")
	}

	rps, err := strconv.ParseInt(os.Args[2], 10, 64)

	if err != nil {
		return "", 0, fmt.Errorf("invalid rate: %s must be a number", os.Args[2])
	}
	if rps < 0 || rps > 100000 {
		return "", 0, errors.New("invalid rate: must be between 0 and 100000")
	}

	return endpoint, rps, nil
}
