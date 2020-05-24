package parser

import (
	"fmt"
	"main/engine"
	"regexp"
)

const imgRe = `<img src="(.*?)" data-pic=".*?" alt="(.*?)" title=".*?" data-bd-imgshare-binded="1">`

func ParseImageUrl(content []byte) engine.ParserResult {
	r := regexp.MustCompile(imgRe)
	allSub := r.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, s := range allSub {
		fmt.Println("image all match=", string(s[0]))
		result.Items = append(result.Items, "img url"+string(s[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://pic.netbian.com" + string(s[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
