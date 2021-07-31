package main

import (
	"log"

	"github.com/scriptonist/grpc-todo-example/cli/internal/cli"
	"github.com/scriptonist/grpc-todo-example/cli/internal/commands"
)

func main() {
	cli, err := cli.NewCLI("localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	cmd := commands.BuildRootCmd(cli)
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
