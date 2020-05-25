module fetcher

go 1.14

require (
	github.com/elazarl/goproxy v0.0.0-20200426045556-49ad98f6dac1 // indirect
	github.com/parnurzeal/gorequest v0.2.16
	github.com/pkg/errors v0.9.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	golang.org/x/net v0.0.0-20200513185701-a91f0712d120
	golang.org/x/text v0.3.2
	moul.io/http2curl v1.0.0 // indirect
	config v0.0.0
)

replace (
    config v0.0.0 => ../config
)
