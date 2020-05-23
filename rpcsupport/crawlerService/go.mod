module crawlerService

go 1.14

require (
	config v0.0.0
	engine v0.0.0
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace (
	config v0.0.0 => ../../config
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
)
