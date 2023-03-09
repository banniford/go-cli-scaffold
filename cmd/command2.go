package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command2Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command2",
		Short: "Command 2 does something",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Command 2 executed")
		},
	}

	cmd.Flags().String("flag3", "", "Flag 3 does something")
	cmd.Flags().Int("flag4", 0, "Flag 4 does something")

	return cmd
}
