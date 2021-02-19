package service

import (
	"github.com/Mirobidjon/todo-app"
	"github.com/Mirobidjon/todo-app/pkg/repository"
)

type TodoListService struct{
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

func (s *TodoListService) Create(UserID int, lists todo.TodoList) (int, error){
	return s.repo.Create(UserID, lists)
}

func (s *TodoListService) GetAll(userID int) ([]todo.TodoList, error){
	return s.repo.GetAll(userID)
}

func (s *TodoListService) GetById(userID, listiD int) (todo.TodoList, error){
	return s.repo.GetById(userID, listiD)
}

func (s *TodoListService) Delete(userID, listiD int) error{
	return s.repo.Delete(userID, listiD)
}

func (s *TodoListService) Update(userID, listiD int, input todo.UpdateListInput) error{
	if err := input.Validate(); err != nil {
		return err
	}
	
	return s.repo.Update(userID, listiD, input)
}