package service

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"time"
	"xblog/config"
	pb "xblog/rpc"
)

func (s *Service) Login(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, error error) {
	name := request.Username
	password := request.Password
	logs.Info(name, password)

	response = &pb.LoginResponse{Result: 1}
	//return response, nil
	return nil, errors.New("the value of filed:id can not be 0")
}

type MyCustomClaims struct {
	Uid int32 `json:"uid"`
	jwt.StandardClaims
}

func CreateJwtToken() string {

	mySigningKey := []byte(config.GetString("jwt.key"))

	// Create the Claims
	claims := MyCustomClaims{
		1,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.GetInt("jwt.ttl")),
			Issuer:    "xblog",
			Subject:   "1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	logs.Info(ss)
	return ss
}
