package auth_service

import "gin-blog/models"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
