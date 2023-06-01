package main

import (
	"github.com/astaxie/beego/logs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	_ "net/http/pprof"
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
	tsl := config.GetBool("app.tsl")

	var opts []grpc.ServerOption
	var credits credentials.TransportCredentials
	if true == tsl {
		// 从输入证书文件和密钥文件为服务端构造TLS凭证
		credits, err = credentials.NewServerTLSFromFile("./tsl/my-public-key-cert.pem", "./tsl/my-private-key.pem")
		if err != nil {
			logs.Error("Failed to generate credentials %v", err)
		}
	} else {
		credits = nil
	}

	opts = []grpc.ServerOption{
		grpc.Creds(credits),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middleware.JwtAuth),
			grpc_validator.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middleware.JwtAuth),
			grpc_validator.UnaryServerInterceptor(),
			//grpc.UnaryServerInterceptor(middleware.ValidMiddleware),
		)),
	}

	grpcServer := grpc.NewServer(opts...)

	// 将 UserInfoService 注册到 gRPC
	// 注意第二个参数 UserInfoServiceServer 是接口类型的变量
	// 需要取地址传参
	pb.RegisterCommonServiceServer(grpcServer, &service.Service{})
	grpcServer.Serve(l)
}
