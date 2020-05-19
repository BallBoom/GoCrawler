package basic

import (
	"fmt"
	"regexp"
	"testing"
)

//const img = `<img src="(/uploads/allimg/[.*]\.jpg)" alt="([^"]*)>`
const img = `<span><img src="(.*?)" alt="(.*?)"</span>`

func TestA(t *testing.T) {
	var s = `<span><img src="/uploads/allimg/180222/231102-15193122629634.jpg" alt="lol娑娜大胸白丝4k壁纸"></span>`
	r := regexp.MustCompile(img)
	res := r.FindAllString(s, -1)
	fmt.Println("result=", res)
	//fmt.Println("result=",res[2])

}
