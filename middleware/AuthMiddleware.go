package middleware

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/wxnacy/wgo/arrays"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"xblog/config"
	"xblog/service"
)

func JwtAuth(ctx context.Context) (context.Context, error) {
	logs.Info(grpc.Method(ctx))
	name, _ := grpc.Method(ctx)
	logs.Info(arrays.ContainsString(config.GetArrayString("methodMiddleware.auth"), name))
	if arrays.ContainsString(config.GetArrayString("methodMiddleware.auth"), name) < 0 {
		return ctx, nil
	}
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	token, err := jwt.ParseWithClaims(tokenString, &service.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	if claims, ok := token.Claims.(*service.MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Uid, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}

	return ctx, nil
}
