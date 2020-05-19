package basic

import (
	"fmt"
	"regexp"
)

// 正则表达式
func RegexTest() {
	var s = "my email is 12094bc@gmail.com"
	r := regexp.MustCompile("12094")
	res := r.FindString(s)
	fmt.Println("result=", res)
}
