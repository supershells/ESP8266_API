package services

type NewUserRequest struct {
	UserID   string `json:"-"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Dept     string `json:"dept"`
}

type UserResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Dept     string `json:"dept"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserService interface {
	NewUser(NewUserRequest) (*UserResponse, error)
	GetUser(UserLogin) (*UserResponse, error)
	GetUserAll() ([]UserResponse, error)
	GetUserOne(string) (*UserResponse, error)
}
