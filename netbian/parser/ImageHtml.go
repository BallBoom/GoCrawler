package parser

import (
	"main/engine"
	"regexp"
)

const htmlRe = `<a href="(/tupian/[0-9]+.html)" target="_blank"><img src=".*?" alt="(.*?)" />`

func ParseImageHtmlList(content []byte) engine.ParserResult {
	r := regexp.MustCompile(htmlRe)
	allSub := r.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, s := range allSub {
		result.Items = append(result.Items, "img html"+string(s[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        "http://pic.netbian.com" + string(s[1]),
			ParserFunc: ParseImageUrl,
		})
	}
	return result
}
