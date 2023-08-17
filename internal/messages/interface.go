package messages

import "github.com/eatmoreapple/openwechat"

type MessageInterface interface {
	handler(*openwechat.Message) error
	ReplyText(*openwechat.Message) error
}

func OnMessage(msg *openwechat.Message) {}
