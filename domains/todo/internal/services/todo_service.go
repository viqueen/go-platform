package services

import (
	"context"
	"github.com/gofrs/uuid"
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	"github.com/viqueen/go-platform/domains/todo/internal/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type TodoService struct {
	todoV1.UnimplementedTodoServiceServer
	access *data.TodoAccess
}

func NewTodoService(access *data.TodoAccess) *TodoService {
	return &TodoService{access: access}
}

func (s TodoService) Create(_ context.Context, req *todoV1.CreateTodoRequest) (*todoV1.TodoResponse, error) {
	todo := &todoV1.Todo{
		Title: req.Title,
		Id:    uuid.Must(uuid.NewV4()).String(),
	}
	response, err := s.access.Todo.Writer.CreateOne(todo)
	if err != nil {
		log.Printf("could not create todo - %v", err)
		return nil, status.Error(codes.Internal, "could not create todo")
	}
	return &todoV1.TodoResponse{
		Todo: response,
	}, nil
}

func (s TodoService) Read(_ context.Context, req *todoV1.ReadTodoRequest) (*todoV1.TodoResponse, error) {
	todoId := uuid.FromStringOrNil(req.Id)
	response, err := s.access.Todo.Reader.ReadOne(todoId)
	if err != nil {
		log.Printf("could not read todo - %v", err)
		return nil, status.Error(codes.Internal, "could not read todo")
	}
	return &todoV1.TodoResponse{
		Todo: response,
	}, nil
}

func (s TodoService) Update(_ context.Context, req *todoV1.UpdateTodoRequest) (*todoV1.TodoResponse, error) {
	panic("implement me")
}

func (s TodoService) Delete(_ context.Context, req *todoV1.DeleteTodoRequest) (*todoV1.TodoResponse, error) {
	panic("implement me")
}
