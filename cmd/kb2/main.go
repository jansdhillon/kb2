package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {

	ctx := context.Background()

	cmd := &cli.Command{
		Name:     "kb2",
		Version:  "v0.0.1",
		Usage:    "Root entry point for kb2",
		Commands: []*cli.Command{},
	}

	if err := cmd.Run(ctx, os.Args); err != nil {
		log.Fatal(err)
	}
}
