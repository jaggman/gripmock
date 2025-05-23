module github.com/tokopedia/gripmock

go 1.23.0

toolchain go1.23.4

require (
	github.com/tokopedia/gripmock/protogen v0.0.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.24.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)

replace github.com/tokopedia/gripmock/protogen v0.0.0 => ./protogen
