package user

type RegisterUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Error    error
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

type FormCreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Error    error
}

type FormUpdateUserInput struct {
	ID    int    `form:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	// Password string `json:"password" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Error error
}

type Profile struct {
	ID int `form:"id" binding:"required"`
}
