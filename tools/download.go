package tools

import (
	"fmt"
	"github.com/djimenez/iconv-go"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

func Download() {
	str := "http://www.guoxue123.com/xiaosuo/shoucao/rjr/%03d.htm"
	i := 0
	fn := "./story"
	for {
		url := fmt.Sprintf(str, i)
		resp, err := http.Get(url)
		if err != nil {
			break
		}
		if resp.StatusCode != 200 {
			break
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		content := make([]byte, len(body))
		iconv.Convert(body, content, "gb2312", "utf8")
		r := regexp.MustCompile(`(?s)<span class="s3">第(.*)`)
		content = r.Find(content)
		r = regexp.MustCompile(`<[^>]+>`)
		content = r.ReplaceAll(content, []byte(""))
		file, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
		_, err = file.Write(content)
		file.Close()
		if err != nil {
			break
		}
		i++
		//time.Sleep(time.Second)
	}
	return
}
