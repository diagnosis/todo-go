package models

import (
	"github.com/google/uuid"
	"time"
)

type Todo struct {
	ItemID    uuid.UUID `json:"item_id"`
	ItemName  string    `json:"item_name"`
	GroupName string    `json:"group_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
