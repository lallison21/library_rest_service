package models

type CreateUserDTO struct {
	Login           string `json:"login,omitempty" example:"johndoe" binding:"required,alphanum"`
	Password        string `json:"password,omitempty" example:"123456" binding:"required,gte=8"`
	ConfirmPassword string `json:"confirm_password,omitempty" example:"123456" binding:"required,gte=8,eqfield=Password"`
	Email           string `json:"email,omitempty" example:"johndoe@example.com" binding:"required,email"`
}

func (u *CreateUserDTO) MapToDAO() *UserDAO {
	return &UserDAO{
		Login:    u.Login,
		Password: u.Password,
		Email:    u.Email,
	}
}

type UserDAO struct {
	Login    string
	Password string
	Email    string
}
