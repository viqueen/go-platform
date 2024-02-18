package services

import (
	"context"
	"github.com/gofrs/uuid"
	todoV1 "github.com/viqueen/go-platform/domains/_api/todo/v1"
	sharedData "github.com/viqueen/go-platform/domains/_shared/data"
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

func (s TodoService) ReadMany(_ context.Context, req *todoV1.ReadManyTodosRequest) (*todoV1.ReadManyTodosResponse, error) {
	response, err := s.access.Todo.Reader.ReadMany(sharedData.PageInfo{
		PageSize:   req.GetPageSize(),
		PageOffset: req.GetPageOffset(),
	})
	if err != nil {
		log.Printf("could not read todos - %v", err)
		return nil, status.Error(codes.Internal, "could not read todos")
	}
	nextPageOffset := req.GetPageOffset() + int32(len(response))
	return &todoV1.ReadManyTodosResponse{
		Todos:          response,
		NextPageOffset: nextPageOffset,
	}, nil
}

func (s TodoService) Update(_ context.Context, req *todoV1.UpdateTodoRequest) (*todoV1.TodoResponse, error) {
	panic("implement me")
}

func (s TodoService) Delete(_ context.Context, req *todoV1.DeleteTodoRequest) (*todoV1.TodoResponse, error) {
	panic("implement me")
}
