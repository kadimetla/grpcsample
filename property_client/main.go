package main

import (
	"context"
	"flag"
	propertyv1 "github.com/kadimetla/grpcsample/gen/proto/propertyapi/property/v1"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := propertyv1.NewPropertyServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	address := &propertyv1.Address{
		Line1: "9999 client st.",
		Line2: "",
		City:  "client",
		State: "client state",
		Lat:   "c1",
		Lng:   "c2",
	}
	property := &propertyv1.Property{
		Name:    "Client testing",
		Address: address,
	}
	r, err := c.SearchProperty(ctx, &propertyv1.PropertyRequest{Property: property})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetProperty())
}
