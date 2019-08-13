package tools

import (
	"fmt"
	"sort"
)

type MapItem struct {
	Key string
	Val interface{}
}

type MapKeySorter []MapItem

func MapKeySort(p interface{}) (MapKeySorter, error) {
	mp, ok := p.(map[string]interface{})
	if !ok {

	}
	ms := make(MapKeySorter, 0, len(mp))
	for k, v := range mp {
		ms = append(ms, MapItem{k, v})
	}
	sort.Sort(ms)
	return ms, nil
}

func (mks MapKeySorter) Len() int {
	return len(mks)
}

func (mks MapKeySorter) Swap(i, j int) {
	mks[i], mks[j] = mks[j], mks[i]
}

func (mks MapKeySorter) Less(i, j int) bool {
	return mks[i].Key < mks[j].Key
}

func init() {
	p := map[string]string{
		"ali":       "china",
		"apple":     "america",
		"microsoft": "america",
		"tencent":   "china",
		"google":    "america",
	}

	ps, _ := MapKeySort(p)
	fmt.Printf("%v", ps)
}
