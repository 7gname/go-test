package reflect_test

import (
	"testing"
	"reflect"
	"fmt"
)

func TestTypeOf(t *testing.T)  {
	var a int
	typeOfA := reflect.TypeOf(a)
	fmt.Printf("%s\t%s\n", typeOfA.Name(), typeOfA.Kind())
}

func Test1(t *testing.T)  {
	f := func (param interface{}) error {
		p := reflect.ValueOf(param)

		if p.Kind() != reflect.Slice {
			t.Log("param is not slice")
		}
		len := p.Len()
		for i := 0; i < len; i++ {
			t.Log(p.Index(i).Interface())
		}
		return nil
	}

	p := []int{1,2,3,4}
	f(p)
}