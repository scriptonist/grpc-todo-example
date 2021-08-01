package cli

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/scriptonist/grpc-todo-example/service/pkg/api"
	"google.golang.org/grpc"
)

type CLI struct {
	apiClient api.TodoAPIClient
}

func NewCLI(apiServerAddr string) (*CLI, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(apiServerAddr, opts...)
	if err != nil {
		return nil, err
	}
	client := api.NewTodoAPIClient(conn)
	cli := &CLI{
		apiClient: client,
	}
	return cli, nil
}

type AddTodoOpts struct {
	Content string
}

func (c *CLI) AddTodo(opts AddTodoOpts) error {
	_, err := c.apiClient.Create(
		context.Background(),
		&api.CreateRequest{
			Description: opts.Content,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

type ListTodoOpts struct {
}

func (c *CLI) ListTodos(opts ListTodoOpts) error {
	stream, err := c.apiClient.Read(context.Background(), &empty.Empty{})
	if err != nil {
		return err
	}
	for {
		todo, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Listing todos: %v", err)
		}
		fmt.Println(todo)
	}
	return nil
}
