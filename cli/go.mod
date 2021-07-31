module github.com/scriptonist/todo/cli

go 1.16

replace github.com/scriptonist/todo/service => ../service

require (
	github.com/scriptonist/todo/service v1.0.0
	github.com/spf13/cobra v1.2.1
	google.golang.org/grpc v1.39.0
)
