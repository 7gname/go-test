package pubsub

type Subscribe interface {
	Subscribe(channel string, cb Cb) bool
}

type Cb func(content Content)

type Subscriber struct {
	Name string
}

func (s Subscriber) Subscribe(channel string, cb Cb) bool {
	go func() {
		for {
			select {
			case content := <-publishChanMap[channel]:
				cb(content)
			}
		}
	}()
	return true
}

func NewSubscriber(name string) Subscribe {
	return Subscriber{name}
}
