module engine

go 1.14

require (
    fetcher v0.0.0
    config v0.0.0 // indirect
)

replace (
    fetcher v0.0.0 => ../fetcher
    config v0.0.0 => ../config
)
