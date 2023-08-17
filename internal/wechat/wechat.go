package wechat

import (
	"log"

	"github.com/eatmoreapple/openwechat"
	"github.com/seaung/wechat_monitor/internal/messages"
)

func RunWechat() {
	robot := openwechat.DefaultBot(openwechat.Desktop)

	robot.UUIDCallback = messages.ConsoleQrCode

	robot.MessageHandler = messages.OnMessage

	storage := openwechat.NewFileHotReloadStorage("storage.json")

	err := robot.HotLogin(storage)
	if err != nil {
		log.Printf("热登录失败，尝试使用普通登录 : %v\n", err)
		if err = robot.Login(); err != nil {
			log.Printf("尝试使用普通登录失败 : %v\n", err)
			return
		}
	}

	robot.Block()
}
