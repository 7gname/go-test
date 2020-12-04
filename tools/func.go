package tools

import (
	"math/rand"
	"time"
)

func RandSlice(s []interface{}) (rs []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(s))
	for _, item := range perm {
		rs = append(rs, s[item])
	}
	return
}

func RandIntSlice(s []int) (rs []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(s))
	for _, item := range perm {
		rs = append(rs, s[item])
	}
	return
}
