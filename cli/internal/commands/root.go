package commands

import (
	"github.com/scriptonist/todo/cli/internal/cli"
	"github.com/spf13/cobra"
)

func BuildRootCmd(c *cli.CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use: "todo",
	}
	cmd.AddCommand(buildCreateCmd(c))
	cmd.AddCommand(buildListCmd(c))
	return cmd
}
