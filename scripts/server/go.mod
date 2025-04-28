module grpc

go 1.23.0

toolchain go1.24.2

// go mod tidy runs after copy to /go/src/grpc
replace github.com/tokopedia/gripmock/stub => /go/src/github.com/tokopedia/gripmock/stub
replace github.com/tokopedia/gripmock/protogen => /go/src/github.com/tokopedia/gripmock/protogen

require (
	github.com/tokopedia/gripmock/protogen v0.0.0
	github.com/tokopedia/gripmock/stub v0.0.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/lithammer/fuzzysearch v1.1.8 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)
