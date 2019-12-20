package middleware

import (
	"context"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/wxnacy/wgo/arrays"
	"google.golang.org/grpc"
	"xblog/config"
	"xblog/service"
)

func JwtAuth(ctx context.Context) (context.Context, error) {
	logs.Info(grpc.Method(ctx))
	name, _ := grpc.Method(ctx)
	logs.Info(arrays.ContainsString(config.GetArrayString("jwt.authMethod"), name))
	if  arrays.ContainsString(config.GetArrayString("jwt.authMethod"), name) < 0 {
		return ctx, nil
	}
	tokenString, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString, &service.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.key")), nil
	})

	if claims, ok := token.Claims.(*service.MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.Zcx, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	if err != nil {
		return nil, err
	}

	return ctx, nil
}
