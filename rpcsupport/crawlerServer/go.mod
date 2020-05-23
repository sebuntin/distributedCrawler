module crawlerServer

go 1.14

require (
	crawlerService v0.0.0
	doubanparser v0.0.0
	google.golang.org/grpc v1.29.1
)

replace (
	config v0.0.0 => ../../config
	crawlerService v0.0.0 => ../crawlerService
	doubanparser v0.0.0 => ../../douban/doubanparser
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
	model v0.0.0 => ../../model
)
