package cobra_implementation

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-implementation",
	Short: "cobra-implementation - it's just simple implementation of cobra pkg",
	Long:  "cobra-implementation - it's just simple implementation of cobra pkg",

	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Something is wrong in a process of working CLI")
		os.Exit(1)
	}
}
