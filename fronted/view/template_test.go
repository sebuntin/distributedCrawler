package view

import (
	"engine"
	"model"
	"os"
	"testing"
)

func TestS(t *testing.T) {
	view := CreateSearchResultView("template.test.html")

	//template:= template.Must(
	//template.ParseFiles("view/template.html"))

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}

	page.Hits = 123

	item := engine.Item{
		Url: "https://book.douban.com/subject/1827374/",
		Id:  "1827374",
		PayLoad: model.Book{
			Name:     "东方快车谋杀案",
			Author:   "[英] 阿加莎·克里斯蒂",
			Press:    "新星出版社",
			OrigName: "Murder on the Orient Express",
			PageNum:  256,
			Price:    "26.00元",
			BriefIntro: `侦探波洛在叙利亚完成了一项委托，要搭乘辛普朗号东方快车回国。奇怪的是，似乎全世界的人都选在那一夜出行，这列铺位一向宽裕的豪华列车竟然一票难求。幸好他遇到了好友、国际客车公司的董事布克先生，这才挤上了车。
		午夜过后，一场大雪迫使辛普朗号停了下来。第二天一早，大家发现少了一名乘客。一个美国人死在了他的包厢里，被刺了十二刀，可他包厢的门却是反锁的。随着调查的深入，案情却似乎更加扑朔迷离，大侦探波洛想出了两种截然不同的解释……`,
			AuthorIntro: `阿加莎•克里斯蒂，无可争议的侦探小说女王，侦探文学史上最伟大的作家之一。
阿加莎•克里斯蒂原名为阿加莎•玛丽•克拉丽莎•米勒，一八九○年九月十五日生于英国德文郡托基的阿什菲尔德宅邸。她几乎没有接受过正规的教育，但酷爱阅读，尤其痴迷于歇洛克•福尔摩斯的故事。
第一次世界大战期间，阿加莎•克里斯蒂成了一名志愿者。战争结束后，她创作了自己的第一部侦探小说《斯泰尔斯庄园奇案》。几经周折，作品于一九二○正式出版，由此开启了克里斯蒂辉煌的创作生涯。一九二六年，《罗杰疑案》由哈珀柯林斯出版公司出版。这部作品一举奠定了阿加莎•克里斯蒂在侦探文学领域不可撼动的地位。之后，她又陆续出版了《东方快车谋杀案》、《ABC 谋杀案》、《尼罗河上的惨案》、《无人生还》、《阳光下的罪恶》等脍炙人口的作品。时至今日，这些作品依然是世界侦探文学宝库里最宝贵的财富。根据她的小说改编而成的...`,
			DoubanScore: 8.9,
		},
	}

	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	//err = template.Execute(out, page)
	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
