package worker

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "worker",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.AddCommand([]*cobra.Command{
		weixinWorkerCmd,
	}...)
}
