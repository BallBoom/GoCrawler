package parser

import (
	"main/engine"
	"regexp"
)

const imgUrl = `<a href="(/4k[a-z]+/)"[^>]*>([^<]*)</a>`
const img = `<img src="(.*?)" alt="(.*?)"`

func ParseImageList(content []byte) engine.ParserResult {
	r := regexp.MustCompile(imgUrl)
	allSub := r.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, s := range allSub {
		result.Items = append(result.Items, string(s[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://pic.netbian.com" + string(s[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
