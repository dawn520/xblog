package service

import (
	"context"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
	"xblog/config"
	"xblog/dao"
	"xblog/model"
	pb "xblog/rpc"
	"xblog/utils"
)

func (s *Service) Login(ctx context.Context, request *pb.LoginRequest) (response *pb.LoginResponse, error error) {
	username := request.Username
	password := request.Password
	logs.Info(username, password)
	var user model.User
	user, _ = dao.GetUserByUsername(username, "")
	if user.ID == 0 {
		return nil, status.Errorf(codes.NotFound, "the user is not exist")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "your password is not invalid")
	}
	token := CreateJwtToken(user)
	response = &pb.LoginResponse{
		Token: token,
		UserInfo: &pb.User{
			Id:       uint32(user.ID),
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
	}
	return response, nil
}

func (s *Service) Register(ctx context.Context, request *pb.RegisterRequest) (response *pb.RegisterResponse, error error) {
	var user model.User
	if request.Password != request.PasswordConfirmed {
		return nil, errors.New("the two passwords are inconsistent")
	}
	utils.StructCopy(request, &user)
	logs.Info(user)
	userInDatabase, err := dao.GetUserByEmail(user.Email, "")
	if userInDatabase != (model.User{}) {
		return nil, status.Errorf(codes.Internal, "the email has been register")
	}
	userInDatabase, err = dao.GetUserByUsername(user.Username, "")
	if userInDatabase != (model.User{}) {
		return nil, status.Errorf(codes.Internal, "the username has been register")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	user.Password = string(hash)
	userId, err1 := dao.CreateUser(user)
	if err1 != nil {
		return nil, status.Errorf(codes.Internal, err1.Error())
	}
	logs.Info(userId)
	response = &pb.RegisterResponse{Uid: uint32(userId)}
	logs.Info(response)

	return response, nil
}

type MyCustomClaims struct {
	Uid      uint32 `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func CreateJwtToken(user model.User) string {

	mySigningKey := []byte(config.GetString("jwt.key"))

	// Create the Claims
	claims := MyCustomClaims{
		uint32(user.ID),
		user.Username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + int64(config.GetInt("jwt.ttl")),
			Issuer:    "xblog",
			Subject:   "uid:" + user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)
	logs.Info(ss)
	return ss
}
