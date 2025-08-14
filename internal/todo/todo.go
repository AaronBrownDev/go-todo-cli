package todo

import "time"

type Priority string

const (
	LOW_PRIORITY    Priority = "low"
	MEDIUM_PRIORITY Priority = "medium"
	HIGH_PRIORITY   Priority = "high"
)

type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    Priority  `json:"priority"`
	Category    string    `json:"category"`
	DueDate     time.Time `json:"due_date"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
