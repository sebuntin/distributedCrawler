module douban

go 1.14

require (
	config v0.0.0
	crawlerService v0.0.0
	doubanparser v0.0.0
	engine v0.0.0
	fetcher v0.0.0 // indirect
	github.com/parnurzeal/gorequest v0.2.16 // indirect
	golang.org/x/net v0.0.0-20200520004742-59133d7f0dd7 // indirect
	google.golang.org/grpc v1.29.1
	itemService v0.0.0 // indirect
	model v0.0.0 // indirect
	moul.io/http2curl v1.0.0 // indirect
	persist v0.0.0
	scheduler v0.0.0
	worker v0.0.0
)

replace (
	config v0.0.0 => ../config
	crawlerService v0.0.0 => ../rpcsupport/crawlerService
	doubanparser v0.0.0 => ./doubanparser
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	itemService v0.0.0 => ../rpcsupport/itemService
	model v0.0.0 => ../model
	persist v0.0.0 => ../persist
	scheduler v0.0.0 => ../scheduler
	worker v0.0.0 => ../worker
)
