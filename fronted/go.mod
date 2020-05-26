module fronted

go 1.14

require (
	config v0.0.0
	controller v0.0.0
	fetcher v0.0.0
	engine v0.0.0
	model v0.0.0
	view v0.0.0
)

replace (
	config v0.0.0 => ../config
	controller v0.0.0 => ./controller
	engine v0.0.0 => ../engine
	fetcher v0.0.0 => ../fetcher
	model v0.0.0 => ../model
	view v0.0.0 => ./view
	config v0.0.0 => ../config
)
