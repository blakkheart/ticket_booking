package user

type Role string

const (
	Admin     Role = "admin"
	Member    Role = "member"
	Anonymous Role = "anonymous"
)

type User struct {
	ID    int64
	Email string
	Name  string
	Role  Role
}

type UserAuth struct {
	ID           int64
	Email        string
	PasswordHash string
	Role         Role
}
