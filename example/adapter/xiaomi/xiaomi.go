package xiaomi

import "fmt"

type XiaoMiAdaptee struct {
}

func (xm XiaoMiAdaptee) BefreSpecificRequest() {

}

func (xm XiaoMiAdaptee) SpecificRequest() {
	fmt.Println("This is a specific request from Xiao Mi")
}

func (xm XiaoMiAdaptee) AfterSpecificRequest() {

}
