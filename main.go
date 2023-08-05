package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/eatmoreapple/openwechat"
	"github.com/robfig/cron/v3"
	"github.com/skip2/go-qrcode"
)

func ConsoleQrCode(uuid string) {
	qrCode, _ := qrcode.New(fmt.Sprintf("https://login.weixin.qq.com/l/%s", uuid), qrcode.Low)
	fmt.Println(qrCode.ToString(true))
}

func OnMessage(msg *openwechat.Message) {
	if msg.IsText() && strings.Contains(msg.Content, "hello") {
		log.Printf("this is a test")
	}
}

func main() {
	cr := cron.New(cron.WithSeconds())
	robot := openwechat.DefaultBot(openwechat.Desktop)

	robot.UUIDCallback = ConsoleQrCode

	robot.MessageHandler = OnMessage

	reloadStorage := openwechat.NewFileHotReloadStorage("storage.json")

	err := robot.HotLogin(reloadStorage)
	if err != nil {
		log.Printf("热登录失败，尝试使用普通登录 : %v\n", err)
		if err = robot.Login(); err != nil {
			log.Printf("普通登录失败, %v\n", err)
			return
		}
	}

	self, err := robot.GetCurrentUser()
	if err != nil {
		log.Printf("获取当前登录用户失败 : %v\n", err)
		return
	}

	firends, err := self.Friends()
	if err != nil {
		log.Printf("获取当前登录用户朋友列表失败 %v\n", err)
		return
	}

	fr := firends.SearchByNickName(1, "昵称")

	cr.AddFunc("*/30 * * * * * *", func() {

		fr.SendText("hello", time.Second*60)

	})

	cr.Start()

	robot.Block()

}
