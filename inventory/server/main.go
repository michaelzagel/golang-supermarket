package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/backslash-security-tests/golang-supermarket/inventory"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 80, "The server port")

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := &inventoryServer{}

	// connect to the postgresql database using environment variables
	srv.postgresqlDB, err = sql.Open("postgres", "")
	if err != nil {
		log.Fatalf("Failed connecting to postgresql database: %v", err)
	}

	s := grpc.NewServer()
	inventory.RegisterInventoryServer(s, srv)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
