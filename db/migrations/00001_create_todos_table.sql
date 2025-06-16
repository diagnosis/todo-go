-- +goose Up
CREATE TABLE todos (
                       item_id UUID PRIMARY KEY,
                       item_name VARCHAR(255) NOT NULL,
                       group_name VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE todos;