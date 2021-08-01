module github.com/scriptonist/grpc-todo-example/cli

go 1.16

replace github.com/scriptonist/grpc-todo-example/service => ../service

require (
	github.com/golang/protobuf v1.5.2
	github.com/scriptonist/grpc-todo-example/service v1.0.0
	github.com/spf13/cobra v1.2.1
	google.golang.org/grpc v1.39.0
)
