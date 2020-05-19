package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request error", err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("request code error",resp.StatusCode)
		return nil, fmt.Errorf("wrong status code %d",
			resp.StatusCode)
	}

	r := determineEncoding(resp.Body)
	//将GBK编码翻译成UTF-8
	UTF8CODE := transform.NewReader(resp.Body, r.NewDecoder())
	read, err := ioutil.ReadAll(UTF8CODE)
	if err != nil {
		fmt.Println("read error", err.Error())
		return nil, err
	}
	//fmt.Printf("%s\n",read)
	return read, err
}

//从获取到网页 获取编码
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
