package huawei

import "fmt"

type HuaWeiAdaptee struct {
	
}

func (hw HuaWeiAdaptee) BefreSpecificRequest()  {

}

func (hw HuaWeiAdaptee) SpecificRequest()  {
	fmt.Println("This is a specific request from Hua Wei")
}

func (hw HuaWeiAdaptee) AfterSpecificRequest()  {

}