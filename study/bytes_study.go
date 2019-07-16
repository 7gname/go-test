package study

import (
	"fmt"
	"bytes"
	"strings"
)
//bytes.Compare 函数返回一个整数表示两个[]byte切片按字典序比较的结果（类同C的strcmp）。如果a==b返回0；如果a<b返回-1；否则返回+1。nil参数视为空切片。
func compare() {
	var a, b []byte
	fmt.Println(bytes.Compare(a, b))

	a = []byte("b")
	fmt.Println(bytes.Compare(a, b))

	b = []byte("bcd")
	fmt.Println(bytes.Compare(a, b))

	c := [][]byte{
		[]byte("abc"),
		[]byte("123"),
		[]byte("bcd"),
	}

	for _, item := range c {
		if 0 == bytes.Compare(b, item) {
			fmt.Println(string(item))
		}
	}
}
//bytes.Equal 判断两个切片的内容是否完全相同
func equal() {
	var a, b []byte
	fmt.Println(bytes.Equal(a, b))

	a = []byte("b")
	fmt.Println(bytes.Equal(a, b))

	b = []byte("bcd")
	fmt.Println(bytes.Equal(a, b))

	c := [][]byte{
		[]byte("abc"),
		[]byte("123"),
		[]byte("bcd"),
	}

	for _, item := range c {
		if bytes.Equal(b, item) {
			fmt.Println(string(item))
			break
		}
	}
}
//bytes.EqualFold 判断两个utf-8编码切片（将unicode大写、小写、标题三种格式字符视为相同）是否相同
func equalFold()  {
	var a, b []byte
	fmt.Println(bytes.EqualFold(a, b))

	a = []byte("b")
	fmt.Println(bytes.EqualFold(a, b))

	b = []byte("bcd")
	fmt.Println(bytes.EqualFold(a, b))

	c := [][]byte{
		[]byte("abc"),
		[]byte("123"),
		[]byte("Bcd"),
	}

	for _, item := range c {
		if bytes.EqualFold(b, item) {
			fmt.Println(string(item))
			break
		}
	}
}
//bytes.Runes 函数返回和s等价的[]rune切片。（将utf-8编码的unicode码值分别写入单个rune）
func runes()  {
	var s = []byte("123")
	fmt.Println(bytes.Runes(s))
}
//bytes.HasPrefix 判断s是否有前缀切片prefix
//bytes.HasSuffix 判断s是否有后缀切片suffix
func hasPreOrSufFix() {
	var s = []byte("jidsweadfjadsah")
	var prefix = []byte("jids")
	var sufFix = []byte("sah")
	fmt.Println(bytes.HasPrefix(s, prefix))
	fmt.Println(bytes.HasSuffix(s, sufFix))
}
//bytes.Contains 判断切片b是否包含子切片subslice
func contains() {
	var s = []byte("jidsweadfjadsah")
	var subFix = []byte("wead")
	fmt.Println(bytes.Contains(s, subFix))
}
//bytes.Count 计算s中有多少个不重叠的sep子切片
func count() {
	var s = []byte("jidsweadfjadsah")
	var sep = []byte("s")
	fmt.Println(bytes.Count(s, sep))
}
//bytes.Index 子切片sep在s中第一次出现的位置，不存在则返回-1
//bytes.IndexByte 字符c在s中第一次出现的位置，不存在则返回-1
//bytes.IndexRune unicode字符r的utf-8编码在s中第一次出现的位置，不存在则返回-1
//bytes.IndexAny 字符串chars中的任一utf-8编码在s中第一次出现的位置，如不存在或者chars为空字符串则返回-1
//bytes.IndexFunc s中第一个满足函数f的位置i（该处的utf-8码值r满足f(r)==true），不存在则返回-1
//bytes.LastIndex 切片sep在字符串s中最后一次出现的位置，不存在则返回-1
//bytes.LastIndexByte 字符c在s中最后一次出现的位置，不存在则返回-1
//bytes.LastIndexAny 字符串chars中的任一utf-8字符在s中最后一次出现的位置，如不存在或者chars为空字符串则返回-1
//bytes.LastIndexFunc s中最后一个满足函数f的unicode码值的位置i，不存在则返回-1
func index() {
	var s = []byte("jidsweadfjadsah")
	var sep = []byte("s")
	var sepByte byte = 97
	var sepRune rune = 97
	var sepStr string = "ea"
	fmt.Println(bytes.Index(s, sep))
	fmt.Println(bytes.IndexByte(s, sepByte))
	fmt.Println(bytes.IndexRune(s, sepRune))
	fmt.Println(bytes.IndexAny(s, sepStr))
	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		if strings.EqualFold(string(r), "w") {
			return true
		}else {
			return false
		}
	}))
	fmt.Println(bytes.LastIndex(s, sep))
	fmt.Println(bytes.LastIndexByte(s, sepByte))
	fmt.Println(bytes.LastIndexAny(s, sepStr))
	fmt.Println(bytes.LastIndexFunc(s, func(r rune) bool {
		if strings.EqualFold(string(r), "w") {
			return true
		}else {
			return false
		}
	}))
}
//bytes.ToLower 返回将所有字母都转为对应的小写版本的拷贝
//bytes.ToUpper 返回将所有字母都转为对应的大写版本的拷贝
func toLowerOrUpper()  {
	var s = []byte("AbCd")
	fmt.Println(string(bytes.ToLower(s)))
	fmt.Println(string(bytes.ToUpper(s)))
}
//bytes.Repeat 返回count个b串联形成的新的切片
func repeat() {
	fmt.Println(string(bytes.Repeat([]byte{97,98}, 5)))
}
//bytes.Replace 返回将s中前n个不重叠old切片序列都替换为new的新的切片拷贝，如果n<0会替换所有old子切片
func replace() {
	fmt.Println(string(bytes.Replace([]byte("abcaba"), []byte("abc"), []byte("123"),-1)))
}
//bytes.Map 将s的每一个unicode码值r都替换为mapping(r)，返回这些新码值组成的切片拷贝。如果mapping返回一个负值，将会丢弃该码值而不会被替换（返回值中对应位置将没有码值）
func mMap() {
	var s = []byte("jidsweadfjadsah")
	fmt.Println(string(bytes.Map(func(r rune) rune {
		if r > 127 {
			return -1
		}else {
			return r - 32
		}
	}, s)))
}

func StudyBytes () {
	//compare()
	//equal()
	//equalFold()
	//runes()
	//hasPreOrSufFix()
	//contains()
	//count()
	//index()
	//toLowerOrUpper()
	//repeat()
	//replace()
	//mMap()
}

