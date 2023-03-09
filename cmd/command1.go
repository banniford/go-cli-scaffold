package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Command1Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "command1",
		Short: "Command 1 does something",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Command 1 executed")
		},
	}

	cmd.Flags().String("flag1", "", "Flag 1 does something")
	cmd.Flags().Int("flag2", 0, "Flag 2 does something")

	return cmd
}
