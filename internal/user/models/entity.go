package models

import "github.com/google/uuid"

type Role string

const (
	Admin     Role = "admin"
	Member    Role = "member"
	Anonymous Role = "anonymous"
)

type User struct {
	ID    uuid.UUID
	Email string
	Name  string
	Role  Role
}

type UserAuth struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	Role         Role
}

type CreateUserParams struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	Role     Role
}
