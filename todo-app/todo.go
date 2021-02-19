package todo

import "errors"

// TodoList ...
type TodoList struct {
	ID          int    `json:"-" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UserList ...
type UserList struct {
	ID     int
	UserID int
	ListID int
}

// TodoItem ...
type TodoItem struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

// ListsItem ...
type ListsItem struct {
	ID     int
	ListID int
	ItemID int
}

// UpdateListInput ...
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// Validate ....
func (s *UpdateListInput) Validate() error {
	if s.Title == nil && s.Description == nil {
		return errors.New("update structure hasn't values")
	}

	return nil
}

// UpdateItemInput struct
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

// Validate ...
func (s *UpdateItemInput) Validate() error {
	if s.Title == nil && s.Description == nil && s.Done == nil{
		return errors.New("update structure hasn't values")
	}

	return nil
}
