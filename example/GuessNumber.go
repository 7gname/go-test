package example

import (
	"time"
	"strconv"
	"os"
	"math/rand"
	"fmt"
)

//猜数字小游戏
func GuessNumber() {
	rand.Seed(time.Now().UnixNano())
	n, _ := strconv.Atoi(os.Args[1])
	r := rand.Intn(n)
	fmt.Print("请输入小于", n, "的数字：")
	var a int
	var c int
	for {
		c ++
		fmt.Scan(&a)
		if a < r {
			fmt.Print("小了\n请重新输入：")
		} else if a > r {
			fmt.Print("大了\n请重新输入：")
		} else {
			fmt.Println("正确，你太棒了")
			break;
		}
	}
	fmt.Println("猜对数字用了", c, "次")
}
