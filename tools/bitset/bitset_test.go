package bitset

import (
	"testing"
	"fmt"
)

func TestBitSets(t *testing.T) {
	bs := NewBitSets(128)
	fmt.Printf("bs[%+v]\n", bs)
	//bs.Set(129)
	//fmt.Printf("bs[%+v]\n", bs)
	b := bs.IsSet(129)
	fmt.Printf("b[%t]\n", b)
	bs.Unset(129)
	b = bs.IsSet(129)
	fmt.Printf("b[%t]\n", b)
}
