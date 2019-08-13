package pubsub

import (
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	var channel string
	var name string
	var subscriber Subscribe
	var publisher Publish

	channel = "sport"
	name = "xiaoming"

	subscriber = NewSubscriber(name)
	subscriber.Subscribe(channel, func(content Content) {
		t.Logf("%v", content)
	})

	channel = "food"
	name = "huahua"
	subscriber = NewSubscriber(name)
	subscriber.Subscribe(channel, func(content Content) {
		t.Logf("%v", content)
	})

	time.Sleep(time.Second * 3)

	channel = "sport"
	name = "laowang"

	publisher = NewPubliser(name)
	publisher.Publish(channel, "恒大3:1国安")

}
