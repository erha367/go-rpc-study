package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"rpcx/proto"
)

/**
 * @description：TODO
 * @author     ：yangsen
 * @date       ：2020/12/3 下午8:46
 * @company    ：eeo.cn
 */

type GrpcServer struct{}

func (s *GrpcServer) Multiply(ctx context.Context, req *proto.ArithRequest) (*proto.ArithResponse, error) {
	log.Println(`--- req ---`)
	return &proto.ArithResponse{Pro: req.A * req.B}, nil
}

func GrpcInit() {
	lis, err := net.Listen(`tcp`, `:8028`)
	if err != nil {
		log.Println(`net listen err`)
		return
	}
	s := grpc.NewServer()
	proto.RegisterArithServiceServer(s, &GrpcServer{})
	reflection.Register(s)
	err = s.Serve(lis)
	if  err != nil {
		log.Println("failed to serve: %v", err)
		return
	}
}
