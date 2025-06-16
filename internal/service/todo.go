package service

import (
	"context"
	"todo-list-backend/internal/models"
	"todo-list-backend/internal/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(ctx context.Context, itemName, groupName string) (*models.Todo, error) {
	todo := &models.Todo{
		ItemName:  itemName,
		GroupName: groupName,
	}
	if err := s.repo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]models.Todo, error) {
	return s.repo.GetAll(ctx)
}

func (s *TodoService) GetTodoByID(ctx context.Context, id string) (models.Todo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TodoService) UpdateTodo(ctx context.Context, id, itemName, groupName string) (*models.Todo, error) {
	todo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	todo.ItemName = itemName
	todo.GroupName = groupName
	if err := s.repo.Update(ctx, &todo); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
