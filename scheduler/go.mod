module scheduler

go 1.14

require (
    engine v0.0.0
    fetcher v0.0.0
    config v0.0.0
)

replace (
    engine v0.0.0 => ../engine
    fetcher v0.0.0 => ../fetcher
    config v0.0.0 => ../config
)