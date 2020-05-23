module itemServer

go 1.14

require (
	engine v0.0.0
	github.com/olivere/elastic/v7 v7.0.15
	google.golang.org/grpc v1.29.1
	itemService v0.0.0
	model v0.0.0
	config v0.0.0
)

replace (
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
	itemService v0.0.0 => ../itemService
	model v0.0.0 => ../../model
	config v0.0.0 => ../../config
)
