package datastore

import "time"

// Book represents a Book
type Book struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Publisher string     `json:"publisher"`
	ISBN      string     `json:"isbn"`
	Year      int        `json:"year"`
	Edition   int        `json:"edition,omitempty"` // omitempty for optional fields
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"` // omitempty for optional fields
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // "admin" or "user"
}

type Rating struct {
	ID        string    `json:"id"`
	BookID    string    `json:"book_id"`
	UserID    string    `json:"user_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
}
