module controller

go 1.14

require (
	engine v0.0.0
	github.com/olivere/elastic/v7 v7.0.15
	model v0.0.0
	view v0.0.0
	config v0.0.0
)

replace (
	config v0.0.0 => ../../config
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
	model v0.0.0 => ../../model
	view v0.0.0 => ../view
	config v0.0.0 => ../../config
)
