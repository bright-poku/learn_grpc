package main

import (
	"context"
	"log"
	"net"

	"github.com/gofrs/uuid"
	pb "productinfo/server/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)
//server is used to implement ecommerce/product_info
type server struct {
	productMap map[string]*pb.Product
}

//AddProduct implements ecommerce.AddProduct
func (s *server) AddProduct(ctx context.Context, prod *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating product ID", err)
	}
	prod.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[prod.Id] = prod
	return &pb.ProductID{Value: prod.Id}, status.New(codes.OK, "").Err()
}

//getProduct implements ecommerce.getProduct
func (s *server) GetProduct(ctx context.Context,  prod *pb.ProductID) (*pb.Product, error) {
	value, exists := s.productMap[prod.Value]
	if exists {
		return value, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "product does not exist. ", prod.Value)
}

//main function 
func main() {
	list, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})

	log.Printf("starting gRPC listener on port " + port)
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}