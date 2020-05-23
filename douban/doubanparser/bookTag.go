package doubanparser

import (
	"engine"
	"regexp"
)

var BookTagReg = regexp.MustCompile(`<td><a href="(/tag/[^"]*)">([^<]*)</a><b>\([\d]+\)</b></td>`)

const PREFIX = "https://book.douban.com"

func ParseBookTag(content []byte, opt ...string) engine.ParserResult {
	result := engine.ParserResult{}
	submatchs := BookTagReg.FindAllSubmatch(content, -1)
	//limit := 2
	for _, m := range submatchs {
		item := string(m[2])
		suffix := string(m[1])
		result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			CurURL: PREFIX + suffix,
			RefURL: opt[0],
			Parser: engine.NewFuncParser("ParseBookList",
				[]string{PREFIX + suffix}, ParseBookList),
		})
		//limit --
		//if limit == 0 {
		//	break
		//}
	}
	return result
}
