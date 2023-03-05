package repositories

type User struct {
	UserID   string `bson:"user_id"`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Role     string `bson:"role"`
	Dept     string `bson:"dept"`
	Status   int    `bson:"status"`
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(string) (*User, error)
	GetByUsername(string) (*User, error)
	GetByDept(string) (*User, error)
	Create(User) (*User, error)
	Update(*User) (*User, error)
	Delete(string) error
}
