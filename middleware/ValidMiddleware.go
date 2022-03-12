package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func ValidMiddleware(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	//if arrays.ContainsString(config.GetArrayString("methodMiddleware.valid"), info.FullMethod) >= 0 {
	//	//走中间件
	//	var recommend model.Post
	//	utils.StructCopy(req, &recommend)
	//	logs.Info(reflect.TypeOf(req), recommend)
	//	result, err := govalidator.ValidateStruct(recommend)
	//	if err != nil {
	//		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	//	}
	//	println(result)
	//}
	resp, err := handler(ctx, req)
	return resp, err
}
