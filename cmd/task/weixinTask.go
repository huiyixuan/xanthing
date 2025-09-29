package task

import (
	"time"
	"xanthing/internal/model"
	"xanthing/internal/service"
	"xanthing/pkg/wechat"

	"github.com/spf13/cobra"
)

var weixinTaskCmd = &cobra.Command{
	Use:   "weixinTask",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		db, _ := service.GetDb("mysql")

		weixinAccount := wechat.Account{
			AppID:     "wxa35c2d70e6ddf7ab",
			AppSecret: "391de8f3516e7914c2d39f295d5f7ff6",
			//AccessToken: "96_F647bdsfqE89ejjaGbZtYFRQh82s3H-krlmzdP2MPJWhnTqYJmuVCaQ_hLPpg5FNt8wIkWfBZWhfnORyP43loTNDCMaSoW5NumkeEXsqp7b5sPj6PIraUp0ZFg4YMTgABARFM",
		}
		_, _ = weixinAccount.GetAccessToken()
		nextOpenid := ""

		tablename := "weixin_user"
		for {
			response, _ := weixinAccount.GetUsers(nextOpenid)
			openids := response.Data["openid"]
			db.Table(tablename).Where("id>0").Update("is_subscribe", 0)
			if response.Count == 0 {
				break
			}
			nextOpenid = openids[len(openids)-1]
			for _, openid := range openids {
				var weixinUser model.WeixinUser
				record := db.Table(tablename).Where("openid = ?", openid).First(&weixinUser)
				if record.RowsAffected == 0 {
					weixinUser.Openid = openid
					weixinUser.WeixinAccount = "gh_12345"
					weixinUser.AddTime = time.Now().Unix()
					weixinUser.UpdateTime = time.Now().Unix()
					weixinUser.IsSubscribe = 1
					db.Table("weixin_user").Create(&weixinUser)
				} else {
					db.Table("weixin_user").Where("openid = ?", weixinUser.Openid).Updates(map[string]any{
						"is_subscribe": 1,
						"update_time":  time.Now().Unix(),
					})
				}
			}

			break
		}

	},
}
