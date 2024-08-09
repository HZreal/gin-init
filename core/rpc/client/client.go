package main

import (
	"context"
	pb "gin-init/core/rpc/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func invoke(x, y int32) (sum, subtract int32) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//
	client := pb.NewCalculatorClient(conn)

	// 测试 Add 方法
	r, err := client.Add(ctx, &pb.AddRequest{A: x, B: y})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	log.Printf("Add: %d", r.Result)
	sum = r.Result

	// 测试 Subtract 方法
	r2, err := client.Subtract(ctx, &pb.SubtractRequest{A: x, B: y})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	log.Printf("Subtract: %d", r2.Result)
	subtract = r2.Result

	return
}
