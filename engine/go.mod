module engine

go 1.14

require (
    fetcher v0.0.0
    config v0.0.0
)

replace (
	config v0.0.0 => ../config
	fetcher v0.0.0 => ../fetcher
)
