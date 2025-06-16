# Todo List Backend

A RESTful API for managing a todo list, built with Go, Chi router, and PostgreSQL, containerized with Docker. Uses pgx for database operations and Goose for migrations.

## Features
- Create, read, update, and delete todo items
- Todo item fields: item_id (UUID), item_name, group_name, created_at, updated_at
- PostgreSQL database with pgx driver
- Chi router for HTTP routing
- Goose for database migrations
- Dockerized application and database

## Prerequisites
- Docker
- Docker Compose
- Go (optional, for running Goose commands locally)

## Setup
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd todo-list-backend