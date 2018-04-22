all: handler.zip

handler.zip: hello
	zip handler.zip hello

hello: hello.go
	GOOS=linux GOARCH=amd64 go build -o hello

