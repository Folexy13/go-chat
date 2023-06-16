package user

import (
	"context"
	"time"
)

type service struct{
	Respository
	timeout time.Duration
}

func NewService(repository Respository) Service{
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, err){
	ctx,cancel:=context.WithTimeout(c, s.timeout)
	defer cancel()

	//TODO: hash password
	var hashedPwd string
	u:=&User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPwd,
	}
}