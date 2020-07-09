package adapter

type Target interface {
	Request()
	Post()
	Get()
}
