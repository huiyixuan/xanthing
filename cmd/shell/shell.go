package shell

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use: "shell",
}

func init() {
	Cmd.AddCommand([]*cobra.Command{
		redisShell,
		testShell,
	}...)
}
