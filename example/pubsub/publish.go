package pubsub

import (
	"go-test/constant"
	"time"
)

type Publish interface {
	Publish(channel string, msg string) bool
}

type Publisher struct {
	Name string
}

var publishChanMap map[string]chan Content

func (p Publisher) Publish(channel string, msg string) bool {
	content := Content{
		Channel:     channel,
		Message:     msg,
		Publisher:   p,
		PublishTime: time.Now().Format(constant.TIME_FORMAT),
	}
	ch := make(chan Content)
	ch <- content
	publishChanMap[channel] = ch
	return true
}

func NewPubliser(name string) Publish {
	return Publisher{name}
}

func init() {
	publishChanMap = make(map[string]chan Content)
}
