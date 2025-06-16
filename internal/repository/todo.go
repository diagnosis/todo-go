package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
	"todo-list-backend/internal/models"
)

type TodoRepository struct {
	db *pgxpool.Pool
}

func NewTodoRepository(db *pgxpool.Pool) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	todo.ItemID = uuid.New()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	query := `INSERT INTO todos (item_id, item_name, group_name, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, todo.ItemID, todo.ItemName, todo.GroupName, todo.CreatedAt, todo.UpdatedAt)
	return err
}

func (r *TodoRepository) GetAll(ctx context.Context) ([]models.Todo, error) {
	query := `
        SELECT item_id, item_name, group_name, created_at, updated_at
        FROM todos
    `
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ItemID, &todo.ItemName, &todo.GroupName, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}

func (r *TodoRepository) GetByID(ctx context.Context, id string) (models.Todo, error) {
	var todo models.Todo
	query := `
        SELECT item_id, item_name, group_name, created_at, updated_at
        FROM todos
        WHERE item_id = $1
    `
	err := r.db.QueryRow(ctx, query, id).Scan(&todo.ItemID, &todo.ItemName, &todo.GroupName, &todo.CreatedAt, &todo.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return models.Todo{}, errors.New("todo not found")
	}
	return todo, err
}

func (r *TodoRepository) Update(ctx context.Context, todo *models.Todo) error {
	todo.UpdatedAt = time.Now()
	query := `
        UPDATE todos
        SET item_name = $2, group_name = $3, updated_at = $4
        WHERE item_id = $1
    `
	result, err := r.db.Exec(ctx, query, todo.ItemID, todo.ItemName, todo.GroupName, todo.UpdatedAt)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return errors.New("todo not found")

	}
	return nil
}
func (r *TodoRepository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM todos 
       	WHERE item_id = $1`
	results, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if results.RowsAffected() == 0 {
		return errors.New("todo not found")
	}
	return nil
}
