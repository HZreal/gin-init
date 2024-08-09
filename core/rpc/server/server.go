package server

import (
	"context"
	"fmt"
	"gin-init/config"
	pb "gin-init/core/rpc/proto"
	"google.golang.org/grpc"
	"net"
)

type RPCServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *RPCServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Printf("receviced ----> %v", req)
	return &pb.AddResponse{Result: req.A + req.B}, nil
}

func (s *RPCServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.SubtractResponse, error) {
	fmt.Printf("receviced ---->  %v", req)
	return &pb.SubtractResponse{Result: req.A - req.B}, nil
}

func StartGPRC() {
	listen, _ := net.Listen("tcp", config.Conf.GRPC.GetAddr())
	// 创建grpc服务
	newServer := grpc.NewServer()
	pb.RegisterCalculatorServer(newServer, &RPCServer{})

	// 启动服务
	go func() {
		err := newServer.Serve(listen)
		if err != nil {
			fmt.Println("[Error] 启动grpc服务器失败！ ", err.Error())
		}
		fmt.Println("gRPC 服务启动成功 in ", config.Conf.GRPC.GetAddr())
	}()
}
