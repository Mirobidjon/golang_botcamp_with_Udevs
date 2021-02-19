package service

import (
	"github.com/Mirobidjon/todo-app"
	"github.com/Mirobidjon/todo-app/pkg/repository"
)

// TodoItemService ...
type TodoItemService struct{
	repo repository.TodoItem
	listRepo repository.TodoList
}

// NewTodoItemService ...
func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService{
	return &TodoItemService{
		repo: repo,
		listRepo: listRepo,
	}
}

// Create ...
func (r *TodoItemService) Create(userID, listID int, input todo.TodoItem) (int, error){
	_, err := r.listRepo.GetById(userID, listID)
	if err != nil {
		return 0, err
	}

	return r.repo.Create(listID, input)
}

func (r *TodoItemService) GetAll(userID, listID int) ([]todo.TodoItem, error){
	return r.repo.GetAll(userID, listID)
}

func (r *TodoItemService) GetById(userID, itemID int) (todo.TodoItem, error){
	return r.repo.GetById(userID, itemID)
}

func (r *TodoItemService) Delete(userID, itemID int) error{
	return r.repo.Delete(userID, itemID)
}

func (r *TodoItemService) Update(userID, itemID int, input todo.UpdateItemInput) error {
	return r.repo.Update(userID, itemID, input)
}