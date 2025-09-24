package cmd

import (
	"os"
	"xanthing/cmd/server"
	"xanthing/cmd/shell"
	"xanthing/internal/service"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "xanthing",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	redisS := service.RedisS{}
	redisS.Init()

	rootCmd.AddCommand([]*cobra.Command{
		server.Cmd,
		shell.Cmd,
	}...)
}
