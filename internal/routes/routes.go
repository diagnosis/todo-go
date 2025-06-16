package routes

import (
	"context"
	"database/sql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib" // Import pgx driver for database/sql
	"github.com/pressly/goose/v3"
	"log"
	"todo-list-backend/internal/config"
	"todo-list-backend/internal/handlers"
	"todo-list-backend/internal/repository"
	"todo-list-backend/internal/service"
)

func SetupRouter(cfg *config.Config) *chi.Mux {
	// Initialize SQL database connection for Goose
	db, err := sql.Open("pgx", cfg.Database.URL)
	if err != nil {
		panic("Failed to connect to database for migrations: " + err.Error())
	}
	defer db.Close()

	// Set up Goose migrations
	goose.SetBaseFS(nil)
	if err := goose.SetDialect("postgres"); err != nil {
		panic("Failed to set Goose dialect: " + err.Error())
	}
	if err := goose.Up(db, "db/migrations"); err != nil {
		panic("Failed to apply migrations: " + err.Error())
	}
	log.Println("Database migrations applied successfully")

	// Initialize pgx connection pool
	pool, err := pgxpool.New(context.Background(), cfg.Database.URL)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	todoRepo := repository.NewTodoRepository(pool)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/todos", func(r chi.Router) {
		r.Post("/", todoHandler.Create)
		r.Get("/", todoHandler.GetAll)
		r.Get("/{id}", todoHandler.GetByID)
		r.Put("/{id}", todoHandler.Update)
		r.Delete("/{id}", todoHandler.Delete)
	})

	return r
}