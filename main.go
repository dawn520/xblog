package main

import (
	"github.com/astaxie/beego/logs"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"net"
	"xblog/config"
	_ "xblog/dao"
	"xblog/middleware"
	_ "xblog/model"
	pb "xblog/rpc"
	"xblog/service"
)


func main() {
	logs.SetLogger("console")
	port := config.GetString("app.listening")
	l, err := net.Listen("tcp", port)
	if err != nil {
		logs.Error("listening error: %v\n", err)
	}
	logs.Info("xblog service is starting, listening %s\n", port)
	s := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middleware.JwtAuth),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.JwtAuth),
		)),
		)
	token := service.CreateJwtToken()
	logs.Info("ZXXZC",token)
	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	pb.RegisterCommonServiceServer(s, &service.Service{})
	s.Serve(l)
}