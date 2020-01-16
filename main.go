package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/hzhyvinskyi/shipper/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
}

type Repository struct {
	mu           sync.RWMutex
	consignments []*pb.Consignment
}

func (r *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	r.mu.RLock()
	updated := append(r.consignments, consignment)
	r.consignments = updated
	r.mu.RUnlock()
	return consignment, nil
}

type service struct {
	repository repository
}

func (s *service) CreateConsignment(ctx context.Context, r *pb.Consignment) (*pb.Response, error) {
	consignment, err := s.repository.Create(r)
	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repository := &Repository{}

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, &service{repository})

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}
