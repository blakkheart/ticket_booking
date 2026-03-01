package user

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

type Account struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  Role   `json:"role"`
}

type AccountFull interface{}

type accountFull struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  Role   `json:"role"`
}

func New(id int64, email string, name string, role Role) AccountFull {

	return accountFull{ID: id, Email: email, Name: name, Role: role}
}
