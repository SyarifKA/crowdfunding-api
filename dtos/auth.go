package dtos

type RegistUser struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}
