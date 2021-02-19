package service

import (
	"github.com/Mirobidjon/todo-app"
	"github.com/Mirobidjon/todo-app/pkg/repository"
)

//Autorization auth
type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
//TodoList interfeys go
type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetById(userID, listiD int) (todo.TodoList, error)
	Delete(userID, listiD int) error
	Update(userID, listiD int, input todo.UpdateListInput) error
}

//TodoItem bla bal
type TodoItem interface {
	Create(userID, listID int, input todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetById(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

// Service ...
type Service struct{
	Autorization
	TodoItem
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList: NewTodoListService(repos.TodoList),
		TodoItem: NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}