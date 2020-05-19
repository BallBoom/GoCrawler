package engine

import (
	"log"
	"main/fetch"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		body, err := fetch.Fetch(r.Url)
		if err != nil {
			log.Printf("fetch is error,url %s error %v", r.Url, err)
			continue
		}
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)

		for _, item := range parserResult.Items {
			log.Printf("get item %v", item)
		}

		for _, item := range parserResult.Requests {
			log.Printf("get url %s", item.Url)
		}
	}
}
