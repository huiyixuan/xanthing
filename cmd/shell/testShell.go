package shell

import (
	"github.com/spf13/cobra"
)

var testShell = &cobra.Command{
	Use:   "test",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
