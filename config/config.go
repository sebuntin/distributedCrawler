package config

const (
	// Parser names
	ParseBookTag  = "ParseBookTag"
	ParseBookList = "ParseBookList"
	ParseBookInfo = "ParseBookInfo"

	// Crawler Service Ports For Test
	CrawlerPort = 1235

	// ElasticSearch
	ElasticIndex = "douban_book"

	// Fronted
	Fronted = 8888
	// Template File Path
	TemplateFilePath = "D:\\go_web\\distributedCrawler\\fronted\\view\\template.html"
	View             = "D:\\go_web\\distributedCrawler\\fronted\\view"

	// Rate Limiting
	Qps = 20
)
