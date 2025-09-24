package shell

import (
	"github.com/spf13/cobra"
)

var redisShell = &cobra.Command{
	Use:   "redis",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
