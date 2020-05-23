module worker

go 1.14

require (
	engine v0.0.0
	google.golang.org/grpc v1.29.1
	config v0.0.0
	crawlerService v0.0.0
)

replace (
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	config v0.0.0 => ../config
	crawlerService v0.0.0 => ../rpcsupport/crawlerService
)
