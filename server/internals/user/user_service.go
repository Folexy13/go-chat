package user

import (
	"context"
	"server/utils"
	"strconv"
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
func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error){
	ctx,cancel:=context.WithTimeout(c, s.timeout)
	defer cancel()

	//TODO: hash password
	hashedPassword,err:=utils.HashPasswword(req.Password)
	if err!=nil{
		return nil,err
	}
	u:=&User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}
	r,err:= s.Respository.CreateUser(ctx,u)
	if err !=nil{
		return nil,err
	}
	res:= &CreateUserRes{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}
	return res,nil
}