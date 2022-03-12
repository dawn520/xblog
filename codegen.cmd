:: protoc --go_out=plugins=grpc:. ./rpc/CommonService.proto

protoc  ^
  --proto_path=%GOPATH%\src ^
  --proto_path=%GOPATH%\src\github.com\google\protobuf\src ^
  --proto_path=. ^
  --go_out=plugins=grpc:. ^
  --govalidators_out=. ^
  ./rpc/CommonService.proto