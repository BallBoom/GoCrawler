package basic

import (
	"fmt"
	"regexp"
	"testing"
)

//const img = `<img src="(/uploads/allimg/[.*]\.jpg)" alt="([^"]*)>`
//const img = `<span><img src="(.*?)" alt="(.*?)"</span>`
//const htmlRe = `<a href="(/tupian/[0-9]+.html)" target="_blank"><img src=".*?" alt="(.*?)">`
const imgRe = `<img src="(.*?)" data-pic=".*?" alt="(.*?)" title=".*?" data-bd-imgshare-binded="1">`

func TestA(t *testing.T) {
	var s = `<a href="" id="img"><img src="/uploads/allimg/200519/004422-15898202622dae.jpg" data-pic="/uploads/allimg/200519/004422-1589820262214f.jpg" alt="短发女孩 心情 城市 唯美艺术插画4k动漫壁纸" title="短发女孩 心情 城市 唯美艺术插画4k动漫壁纸" data-bd-imgshare-binded="1">`
	r := regexp.MustCompile(imgRe)
	cont := []byte(s)
	res := r.FindAllSubmatch(cont, -1)
	for i, s1 := range res {
		fmt.Println("result=", string(s1[i]))
	}
}
