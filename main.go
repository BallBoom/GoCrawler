package main

import (
	"main/engine"
	"main/netbian/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://pic.netbian.com/",
		ParserFunc: parser.ParseImageList,
	})

}
