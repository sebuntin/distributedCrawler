module itemService

go 1.14

require (
	engine v0.0.0
	github.com/golang/protobuf v1.4.2
	github.com/olivere/elastic/v7 v7.0.15
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.23.0
)

replace (
	engine v0.0.0 => ../../engine
	fetcher v0.0.0 => ../../fetcher
	model v0.0.0 => ../../model
)
