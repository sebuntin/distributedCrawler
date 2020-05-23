package doubanparser

import (
	"engine"
	"model"
	"regexp"
	"strconv"
	"strings"
)

/*

 */
var (
	AuthorRe      = regexp.MustCompile(`<a href="https://book.douban.com/author/[\d]+/">([^<]*)</a>`)
	IdRe          = regexp.MustCompile(`https://book.douban.com/subject/([\d]+)/`)
	PressRe       = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]*)<br/>`)
	OrigNameRe    = regexp.MustCompile(`<span class="pl">原作名:</span> ([^<]*)<br/>`)
	PageNumRe     = regexp.MustCompile(`<span class="pl">页数:</span> ([\d]+)<br/>`)
	PriceRe       = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]*)<br/>`)
	IntroRe       = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([\d\D]*?)</div>`)
	DoubanScoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">\s+([^\s]*)\s+</strong>`)
)

func ParseBookInfo(content []byte, opt ...string) engine.ParserResult {
	bookinfo := model.Book{}
	bookinfo.Name = opt[1]
	bookid := extractString([]byte(opt[0]), IdRe)
	author := extractString(content, AuthorRe)
	// 去除空格
	author = strings.Replace(author, " ", "", -1)
	author = strings.Replace(author, "\n", "", -1)
	bookinfo.Author = author
	bookinfo.Press = extractString(content, PressRe)
	bookinfo.OrigName = extractString(content, OrigNameRe)
	bookinfo.Price = extractString(content, PriceRe)
	intro := extractAllString(content, IntroRe)
	briefIntro := strings.Replace(intro[0], "<p>", "", -1)
	briefIntro = strings.Replace(briefIntro, "</p>", "", -1)
	authorIntro := strings.Replace(intro[1], "<p>", "", -1)
	authorIntro = strings.Replace(authorIntro, "</p>", "", -1)
	bookinfo.BriefIntro = briefIntro
	bookinfo.AuthorIntro = authorIntro
	scoreStr := extractString(content, DoubanScoreRe)
	score, err := strconv.ParseFloat(scoreStr, 10)
	if err != nil {
		score = 0
	}
	bookinfo.DoubanScore = score
	bookinfo.PageNum = extractInt(content, PageNumRe)
	return engine.ParserResult{
		Items: []interface{}{engine.Item{Url: opt[0], Id: bookid, PayLoad: bookinfo}},
	}
}

func extractString(content []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(content)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractAllString(content []byte, re *regexp.Regexp) []string {
	match := re.FindAllSubmatch(content, -1)
	result := make([]string, 10)
	for i, m := range match {
		if len(m) != 2 {
			continue
		}
		result[i] = string(m[1])
	}
	return result
}

func extractInt(content []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(content)
	if len(match) == 2 {
		matchInt, err := strconv.Atoi(string(match[1]))
		if err != nil {
			return 0
		}
		return matchInt
	}
	return 0
}
