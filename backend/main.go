package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
       "google.golang.org/grpc/reflection"

	pb "tempconv-grpc/backend/gen"
)

// server implements the gRPC service
type server struct {
	pb.UnimplementedTempConvServiceServer
}

// C2F converts Celsius to Fahrenheit
func (s *server) C2F(ctx context.Context, req *pb.TempRequest) (*pb.TempResponse, error) {
	c := req.GetValue()
	f := c*9.0/5.0 + 32.0
	return &pb.TempResponse{Value: f}, nil
}

// F2C converts Fahrenheit to Celsius
func (s *server) F2C(ctx context.Context, req *pb.TempRequest) (*pb.TempResponse, error) {
	f := req.GetValue()
	c := (f - 32.0) * 5.0 / 9.0
	return &pb.TempResponse{Value: c}, nil
}

func main() {
	port := 50051

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTempConvServiceServer(grpcServer, &server{})
        reflection.Register(grpcServer)

	log.Printf("TempConv gRPC server listening on :%d", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}