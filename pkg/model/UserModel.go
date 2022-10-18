package model

type LoginForm struct {
	UserID   string `form:"userid"`
	Password string `form:"password"`
}
