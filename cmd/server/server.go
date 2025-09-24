package server

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "server",
}

func init() {
	Cmd.AddCommand([]*cobra.Command{
		httpServerCmd,
	}...)
}
