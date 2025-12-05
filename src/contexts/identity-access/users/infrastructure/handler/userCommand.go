package handler

type UserCommand struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	CreatedAt float64 `json:"created_at"`
}
