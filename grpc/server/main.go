package main

import (
	"fmt"
	"log"
	"net"

	Pb "server/chartroom"
	"google.golang.org/grpc"
	SE "server/server"
)

const (
	ip   = "127.0.0.1"
	port = "23333"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", ip, port))
	if err != nil {
		log.Fatalf("无法监听端口 %v %v", port, err)
	}
	s := grpc.NewServer()
	// ^ 注册服务
	Pb.RegisterChatRoomServer(s, &SE.Service{})
	log.Println("gRPC服务器开始监听", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("提供服务失败: %v", err)
	}
}
