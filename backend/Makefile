PROJECT_NAME = url_shorten

run:
	export MODE_ENV=$(env) && go run *.go

init_server:
	hz new -module ${PROJECT_NAME} -idl ${PROJECT_NAME}.thrift
	go mod tidy
server:
	hz update -idl ${PROJECT_NAME}.thrift
	go mod tidy

test:
	go test $(path) -v
test_func:
	go test ./... -v -run $(func)