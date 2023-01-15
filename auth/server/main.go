package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/backslash-security-tests/golang-supermarket/auth"
	"github.com/backslash-security-tests/golang-supermarket/customers"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 80, "The server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := &authServer{}

	// connect to the customers service via gRPC
	customersConn, err := grpc.Dial("customers", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service customers: %v", err)
	}

	defer customersConn.Close()

	srv.customersClient = customers.NewCustomersClient(customersConn)

	s := grpc.NewServer()
	auth.RegisterAuthServer(s, srv)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
