package xiaomi

import (
	"go-test/tools"
	"math/rand"
	"sort"
	"time"
)

//砍价
//originPrice 原价 targetPric砍价的目标价格 peoples砍价人数
func Bargaining(originPrice, targetPrice, peoples int) (rs []int) {
	bargainingPrice := originPrice - targetPrice

	i := 1
	max := bargainingPrice
	total := 0
	prices := make([]int, 0)
	for i <= peoples {
		price := 0
		if i < peoples {
			rand.Seed(time.Now().UnixNano())
			price = rand.Intn(max / (peoples - i))
			total = total + price
			max = max - price
		} else {
			price = bargainingPrice - total
		}
		prices = append(prices, price)
		i++
	}

	sort.Slice(prices, func(i, j int) bool {
		return prices[i] > prices[j]
	})

	len := len(prices)
	cnt := len / 10
	j := 0
	for j < cnt {
		offset := j * 10
		end := (j + 1) * 10
		if end >= len {
			end = len - 1
		}
		s := prices[offset:end]
		ns := tools.RandIntSlice(s)
		rs = append(rs, ns...)
		j++
	}

	return
}
