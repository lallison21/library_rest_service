package models

type CreateUser struct {
	Login           string `binding:"required,alphanum"               example:"johndoe"             json:"login,omitempty"`
	Password        string `binding:"required,gte=8"                  example:"123456"              json:"password,omitempty"`        //nolint:lll // it's struct tags
	ConfirmPassword string `binding:"required,gte=8,eqfield=Password" example:"123456"              json:"confirmPassword,omitempty"` //nolint:lll // it's struct tags
	Email           string `binding:"required,email"                  example:"johndoe@example.com" json:"email,omitempty"`
}

func (u *CreateUser) MapToDTO() *UserDTO {
	return &UserDTO{
		Login:    u.Login,
		Password: u.Password,
		Email:    u.Email,
	}
}

type UserDTO struct {
	Login    string
	Password string
	Email    string
}

func (u *UserDTO) MapToDAO() *UserDAO {
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
