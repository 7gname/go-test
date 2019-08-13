package adapter

import "testing"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee("XiaoMi")
	adapter := NewAdapter(adaptee)

	adapter.Request()

	adaptee = NewAdaptee("HuaWei")
	adapter = NewAdapter(adaptee)

	adapter.Request()

	target := Target(NewAdapter(NewAdaptee("XiaoMi")))
	target.Request()
}
