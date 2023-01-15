package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/backslash-security-tests/golang-supermarket/customers"
	"github.com/backslash-security-tests/golang-supermarket/delivery"
	"github.com/backslash-security-tests/golang-supermarket/inventory"
	"github.com/backslash-security-tests/golang-supermarket/orders"
	"github.com/backslash-security-tests/golang-supermarket/payments"
	"github.com/backslash-security-tests/golang-supermarket/pricing"
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

	srv := &ordersServer{}

	// connect to the delivery service via gRPC
	deliveryConn, err := grpc.Dial("delivery", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service delivery: %v", err)
	}

	defer deliveryConn.Close()

	srv.deliveryClient = delivery.NewDeliveryClient(deliveryConn)

	// connect to the inventory service via gRPC
	inventoryConn, err := grpc.Dial("inventory", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service inventory: %v", err)
	}

	defer inventoryConn.Close()

	srv.inventoryClient = inventory.NewInventoryClient(inventoryConn)

	// connect to the pricing service via gRPC
	pricingConn, err := grpc.Dial("pricing", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service pricing: %v", err)
	}

	defer pricingConn.Close()

	srv.pricingClient = pricing.NewPricingClient(pricingConn)

	// connect to the payments service via gRPC
	paymentsConn, err := grpc.Dial("payments", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service payments: %v", err)
	}

	defer paymentsConn.Close()

	srv.paymentsClient = payments.NewPaymentsClient(paymentsConn)

	// connect to the customers service via gRPC
	customersConn, err := grpc.Dial("customers", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service customers: %v", err)
	}

	defer customersConn.Close()

	srv.customersClient = customers.NewCustomersClient(customersConn)

	// connect to the postgresql database using environment variables
	srv.postgresqlDB, err = sql.Open("postgres", "")
	if err != nil {
		log.Fatalf("Failed connecting to postgresql database: %v", err)
	}

	s := grpc.NewServer()
	orders.RegisterOrdersServer(s, srv)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
