package messages

import (
	"fmt"
	"log"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/skip2/go-qrcode"
)

type UserMessageHandler struct{}

// 确保UserMessageHandler实现了MessageInterface
var _ MessageInterface = (*UserMessageHandler)(nil)

func (u *UserMessageHandler) handler(msg *openwechat.Message) error {
	if msg.IsText() {
		return u.ReplyText(msg)
	}
	return nil
}

func (u *UserMessageHandler) ReplyText(msg *openwechat.Message) error {
	sender, _ := msg.Sender()

	log.Printf("来自用户 %s 的消息 : %v\n", sender.NickName, msg.Context())

	text := strings.Trim(msg.Content, "\n")

	_, err := msg.ReplyText(text)
	if err != nil {
		return err
	}

	return nil
}

func ConsoleQrCode(uuid string) {
	qrCode, _ := qrcode.New(fmt.Sprintf("https://login.weixin.qq.com/l/%s", uuid), qrcode.Low)
	fmt.Println(qrCode.ToString(true))
}
