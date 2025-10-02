package model

type Role string

const (
	Admin     Role = "admin"
	Member    Role = "member"
	Anonymous Role = "anonymous"
)

type AccountIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
