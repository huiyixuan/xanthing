package shell

import (
	"xanthing/internal/service"

	"github.com/spf13/cobra"
)

var testShell = &cobra.Command{
	Use:   "test",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		db, _ := service.GetDb("mysql")
		a := make(map[string]interface{})
		a["content"] = "aaaaa"
		db.Table("weixin_official_message").Create(&a)

	},
}
