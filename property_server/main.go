package main

import (
	"context"
	"flag"
	"fmt"
	propertyv1 "github.com/kadimetla/grpcsample/gen/proto/propertyapi/property/v1"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	propertyv1.UnimplementedPropertyServiceServer
}

func (s *server) SearchProperty(ctx context.Context,

	in *propertyv1.PropertyRequest) (*propertyv1.PropertyResponse, error) {
	log.Printf("Received: %v", in.String())

	address := &propertyv1.Address{
		Line1: "1234 test st.",
		Line2: "",
		City:  "yemmiganur",
		State: "Andhra Pradesh",
		Lat:   "1",
		Lng:   "2",
	}
	property := &propertyv1.Property{
		Name:    "Adi testing",
		Address: address,
	}
	return &propertyv1.PropertyResponse{Property: property}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	propertyv1.RegisterPropertyServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
