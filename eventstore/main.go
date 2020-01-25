package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/hzhyvinskyi/shipper/pb"
)

const (
	port = ":50051"
	clusterID = "test-cluster"
	clientID = "event-store-api"
)

type server struct {
	//
}

// GetEvents gets Events from EventStore by given EventFilter
func (s server) GetEvents(context.Context, *pb.EventFilter) (*pb.EventResponse, error) {
	panic("implement me")
}

// CreateEvent creates a new Event into the EventStore
func (s server) CreateEvent(context.Context, *pb.Event) (*pb.Response, error) {
	panic("implement me")
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	pb.RegisterEventStoreServer(s, &server{})
	s.Serve(listener)
}
