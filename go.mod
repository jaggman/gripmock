module github.com/tokopedia/gripmock

go 1.23.0

toolchain go1.23.4

require (
	github.com/go-chi/chi/v5 v5.2.1
	github.com/lithammer/fuzzysearch v1.1.8
	github.com/stretchr/testify v1.7.0
	github.com/tokopedia/gripmock/protogen v0.0.0
	github.com/tokopedia/gripmock/protogen/example v0.0.0
	golang.org/x/text v0.24.0
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.39.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
	google.golang.org/genproto v0.0.0-20250414145226-207652e42e2e // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250409194420-de1ac958c67a // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200615113413-eeeca48fe776 // indirect
)

// this is for generated server to be able to run
replace github.com/tokopedia/gripmock/protogen/example v0.0.0 => ./protogen/example

// this is for example client to be able to run
replace github.com/tokopedia/gripmock/protogen v0.0.0 => ./protogen
