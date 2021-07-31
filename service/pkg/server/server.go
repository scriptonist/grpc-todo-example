package server

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/scriptonist/grpc-todo-example/service/pkg/api"
)

// UnimplementedTodoAPIServer must be embedded to have forward compatible implementations.
type Server struct {
	*api.UnimplementedTodoAPIServer
	mutex sync.Mutex
	todos []*api.Todo
}

func New() *Server {
	return &Server{}
}

func (s *Server) addTodo(todo *api.Todo) {
	defer s.mutex.Unlock()
	s.mutex.Lock()
	s.todos = append(s.todos, todo)
}

func (s *Server) Create(ctx context.Context, r *api.CreateRequest) (*api.Void, error) {
	todo := &api.Todo{
		Id:          uuid.New().String(),
		Description: r.Description,
		Completed:   false,
	}
	s.addTodo(todo)
	return &api.Void{}, nil
}

func (s *Server) Read(_ *api.Void, stream api.TodoAPI_ReadServer) error {
	for _, todo := range s.todos {
		if err := stream.Send(todo); err != nil {
			return err
		}
	}
	return nil
}
