package adapter

import (
	"go-test/example/adapter/huawei"
	"go-test/example/adapter/xiaomi"
	"strings"
)

type Adaptee interface {
	BefreSpecificRequest()
	SpecificRequest()
	AfterSpecificRequest()
}

func NewAdaptee(adt string) Adaptee {
	switch strings.ToLower(adt) {
	case "xiaomi":
		return xiaomi.XiaoMiAdaptee{}
	case "huawei":
		return huawei.HuaWeiAdaptee{}
	default:
		return nil
	}
}

type Adapter struct {
	Adaptee
}

func NewAdapter(adaptee Adaptee) Adapter {
	return Adapter{
		Adaptee: adaptee,
	}
}

func (a Adapter) Request() {
	a.BefreSpecificRequest()
	a.SpecificRequest()
	a.AfterSpecificRequest()
}

func (a Adapter) Post() {

}

func (a Adapter) Get() {

}
