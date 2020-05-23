module persist

go 1.14

require (
	github.com/olivere/elastic/v7 v7.0.15
	engine v0.0.0
	fetcher v0.0.0
	model v0.0.0
	itemService v0.0.0
	config v0.0.0
)

replace (
	model v0.0.0 => ../model
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	model v0.0.0 => ../model
	itemService v0.0.0 => ../rpcsupport/itemService
	config v0.0.0 => ../config
)
