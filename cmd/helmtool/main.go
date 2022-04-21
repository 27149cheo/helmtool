package main

import (
	"log"

	"github.com/27149cheo/helmtool/pkg/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("Failed to run helmtool, error: %s", err)
	}
}
