package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/devlopersabbir/ddos-cli/cmd"
	cnf "github.com/devlopersabbir/ddos-cli/configs"
)

func mian() {
	fmt.Println("Welcome to ddos-cli")

	endpoint, rps, err := cmd.Cli()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Endpoint: %s\n", endpoint)
	fmt.Printf("Rate: %d\n requests/sec", rps)

	matrics := &cnf.Metrics{}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
}
