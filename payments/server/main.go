package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/backslash-security-tests/golang-supermarket/payments"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 80, "The server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := &paymentsServer{}

	s := grpc.NewServer()
	payments.RegisterPaymentsServer(s, srv)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
