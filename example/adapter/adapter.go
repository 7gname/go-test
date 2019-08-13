package adapter

import (
	"go-test/example/adapter/xiaomi"
	"go-test/example/adapter/huawei"
)

type Adaptee interface {
	BefreSpecificRequest()
	SpecificRequest()
	AfterSpecificRequest()
}

func NewAdaptee(t string) Adaptee {
	switch t {
	case "XiaoMi":
		return xiaomi.XiaoMiAdaptee{}
	case "HuaWei":
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

func (a Adapter) Request()  {
	a.BefreSpecificRequest()
	a.SpecificRequest()
	a.AfterSpecificRequest()
}