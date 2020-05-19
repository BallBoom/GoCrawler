package engine

type Request struct {
	Url        string                    //请求的url
	ParserFunc func([]byte) ParserResult //将Fetcher爬取到的数据放到函数中 返回解析的Request结果
}

type ParserResult struct {
	Requests []Request     //请求结构 数组
	Items    []interface{} //解析完要的数据数组
}

func NilParser([]byte) ParserResult {
	return ParserResult{}
}
