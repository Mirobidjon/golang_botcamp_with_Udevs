package repository

import (
	"github.com/Mirobidjon/todo-app"
	"github.com/jmoiron/sqlx"
)

//golang auth
type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

//TodoList interfeys go
type TodoList interface {
	Create(UserID int, lists todo.TodoList) (int, error)
	GetAll(userID int) ([]todo.TodoList, error)
	GetById(userID, listiD int) (todo.TodoList, error)
	Delete(userID, listiD int) error
	Update(userID, listiD int, input todo.UpdateListInput) error
}

//TodoItem bla bal
type TodoItem interface {
	Create(listID int, input todo.TodoItem) (int, error)
	GetAll(userID, listID int) ([]todo.TodoItem, error)
	GetById(userID, itemID int) (todo.TodoItem, error)
	Delete(userID, itemID int) error
	Update(userID, itemID int, input todo.UpdateItemInput) error
}

//Repository ...
type Repository struct {
	Autorization
	TodoItem
	TodoList
}

// NewRepository get
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
		TodoList:     NewTodoListPostgres(db),
		TodoItem: NewTodoItemPostgres(db),
	}
}
