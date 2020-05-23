module rpcsupport

go 1.14

require (
	engine v0.0.0
	fetcher v0.0.0 // indirect
	doubanparser v0.0.0
	github.com/olivere/elastic/v7 v7.0.15
	google.golang.org/grpc v1.29.1
	itemService v0.0.0
	crawlerService v0.0.0
	config v0.0.0 // indirect
	model v0.0.0
)

replace (
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	doubanparser v0.0.0 => ../douban/doubanparser
	itemService v0.0.0 => ./itemService
	crawlerService v0.0.0 => ./crawlerService
	config v0.0.0 => ../config
	model v0.0.0 => ../model
)
