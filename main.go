package main

import (
	"fmt"

	"github.com/go-cli-scaffold/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "My CLI Application",
		Long:  "My CLI Application is a tool for doing stuff",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, world!")
		},
	}

	rootCmd.AddCommand(cmd.Command1Cmd())
	rootCmd.AddCommand(cmd.Command2Cmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
