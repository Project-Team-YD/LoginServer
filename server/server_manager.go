package server

import (
	"context"
	"crypto/tls"
	"encoding/json"
	global_grpc "project_yd/grpc"
	"project_yd/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const SERVER_PORT = ":8081"
const SERVER_NAME = "login_server"

var Client global_grpc.GlobalGRpcServiceClient

func ConnectToNotificationServer() {
	conn, err := grpc.Dial(
		util.NotificationIp+util.NotificationPort,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})))

	if err != nil {
		println("Connect To NotificationServer Error!!")
		println(err.Error())
		return
	}
	defer conn.Close()

	Client = global_grpc.NewGlobalGRpcServiceClient(conn)
}

func GlobalGrpcMessage(rpcKey string, data interface{}) string {
	message, err := json.Marshal(data)
	if err != nil {
		println("GlobalGrpcMessage Error!!")
		println(err.Error())
		return ""
	}
	request := &global_grpc.GlobalGrpcRequest{
		RpcKey:  rpcKey,
		Message: string(message),
	}

	response, err := Client.GlobalGRpc(context.Background(), request)

	if err != nil {
		println("GlobalGrpcMessage Response Error!!")
		println(err.Error())
		return ""
	}

	return response.Message
}
