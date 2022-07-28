package cobra_implementation

import (
	cobra_implementation "alekseikromski.space/cobra-implmentation/pkg/cobra-implementation"
	"fmt"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Aliases: []string{"rev"},
	Short:   "Reverses a string",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := cobra_implementation.Reverse(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
