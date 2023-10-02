package respons

import "time"

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
