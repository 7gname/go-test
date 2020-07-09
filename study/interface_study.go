package study

import (
	"fmt"
	"reflect"
)

type Action interface {
	Call()
}

type Cat struct {
	Voice string
}

func NewCat() (c *Cat) {
	c = &Cat{Voice: "喵喵喵..."}
	return
}

func (c *Cat) Call() {
	fmt.Println(c.Voice)
}

func (c *Cat) Living() {
	fmt.Println("你们这些猫奴，听好了：我是神！")
}

type Dog struct {
	Voice string
}

func NewDog() (d *Dog) {
	d = &Dog{Voice: "汪汪汪..."}
	return
}

func (d *Dog) Call() {
	fmt.Println(d.Voice)
}

type JinMaoDog struct {
	Dog
	Skin string
}

func NewJinMaoDog() (jm *JinMaoDog) {
	jm = &JinMaoDog{Dog{Voice: "汪汪汪"}, "金黄色"}
	return
}

func (jm *JinMaoDog) Look() {
	fmt.Println(jm.Skin)
}

func StudyInterface() {
	cat := NewCat()
	cat.Call()
	cat.Living()

	dog := NewDog()
	dog.Call()

	jmDog := NewJinMaoDog()
	jmDog.Call()
	jmDog.Look()

	t := []Action{cat, dog, jmDog}
	for _, item := range t {
		fmt.Println(reflect.TypeOf(item).String())
		item.Call()
	}
}
