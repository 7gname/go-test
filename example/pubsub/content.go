package pubsub

type Content struct {
	Channel string
	Message string
	Publisher
	PublishTime string
}
