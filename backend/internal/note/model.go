package note

import "time"

type Note struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	URLs      []string  `json:"urls"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateResult struct {
	Note
	AIFallback bool `json:"ai_fallback"`
}
