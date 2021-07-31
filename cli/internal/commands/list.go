package commands

import (
	"github.com/scriptonist/grpc-todo-example/cli/internal/cli"
	"github.com/spf13/cobra"
)

type ListCmdOpts struct{}

func buildListCmd(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListTodos(cli.ListTodoOpts{})
		},
	}

	return cmd
}
