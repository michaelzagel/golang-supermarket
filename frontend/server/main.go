package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/auth"
	"github.com/backslash-security-tests/golang-supermarket/customers"
	"github.com/backslash-security-tests/golang-supermarket/frontend"
	"github.com/backslash-security-tests/golang-supermarket/orders"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type FrontendServer struct {
	authClient      auth.AuthClient
	ordersClient    orders.OrdersClient
	customersClient customers.CustomersClient
	restClient      http.Client
}

func (srv *FrontendServer) FinishOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input frontend.FinishOrderInput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")
	input.OrderId = mux.Vars(r)["order_id"]

	output, err := srv.FinishOrder(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *FrontendServer) FinishOrder(ctx context.Context, input frontend.FinishOrderInput) (
	output frontend.FinishOrderOutput,
	err error,
) {
	access_token := input.AccessToken
	order_id := input.OrderId
	credit_card := input.CreditCard

	call0, _ := srv.authClient.VerifyToken(ctx, &auth.VerifyTokenInput{
		AccessToken: access_token,
	})

	customer_id := call0.CustomerId

	call1, _ := srv.ordersClient.Finish(ctx, &orders.FinishInput{
		OrderId:    order_id,
		CreditCard: credit_card,
	})

	window_from := call1.WindowFrom
	window_to := call1.WindowTo

	output.WindowFrom = window_from
	output.WindowTo = window_to
	return output, nil
}

func (srv *FrontendServer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var input frontend.RegisterInput
	json.NewDecoder(r.Body).Decode(&input)

	output, err := srv.Register(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *FrontendServer) Register(ctx context.Context, input frontend.RegisterInput) (
	output frontend.RegisterOutput,
	err error,
) {
	delivery_address := input.DeliveryAddress
	password := input.Password
	email := input.Email

	call0, _ := srv.customersClient.Register(ctx, &customers.RegisterInput{
		Email:           email,
		DeliveryAddress: delivery_address,
		Password:        password,
	})

	customer_id := call0.CustomerId

	output.CustomerId = customer_id
	return output, nil
}

func (srv *FrontendServer) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var input frontend.SignInInput
	json.NewDecoder(r.Body).Decode(&input)

	output, err := srv.SignIn(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *FrontendServer) SignIn(ctx context.Context, input frontend.SignInInput) (
	output frontend.SignInOutput,
	err error,
) {
	email := input.Email
	password := input.Password

	call0, _ := srv.authClient.SignIn(ctx, &auth.SignInInput{
		Email:    email,
		Password: password,
	})

	access_token := call0.AccessToken

	output.Success = success
	output.AccessToken = access_token
	return output, nil
}

func (srv *FrontendServer) StartOrderHandler(w http.ResponseWriter, r *http.Request) {
	var input frontend.StartOrderInput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")

	output, err := srv.StartOrder(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *FrontendServer) StartOrder(ctx context.Context, input frontend.StartOrderInput) (
	output frontend.StartOrderOutput,
	err error,
) {
	access_token := input.AccessToken

	call0, _ := srv.authClient.VerifyToken(ctx, &auth.VerifyTokenInput{
		AccessToken: access_token,
	})

	customer_id := call0.CustomerId

	call1, _ := srv.ordersClient.Start(ctx, &orders.StartInput{
		CustomerId: customer_id,
	})

	order_id := call1.OrderId

	output.OrderId = order_id
	return output, nil
}

func (srv *FrontendServer) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	var input frontend.AddItemInput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")
	input.OrderId = mux.Vars(r)["order_id"]

	output, err := srv.AddItem(r.Context(), input)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(output)
}

func (srv *FrontendServer) AddItem(ctx context.Context, input frontend.AddItemInput) (
	output frontend.AddItemOutput,
	err error,
) {
	access_token := input.AccessToken
	order_id := input.OrderId
	name := input.Name
	quantity := input.Quantity

	call0, _ := srv.authClient.VerifyToken(ctx, &auth.VerifyTokenInput{
		AccessToken: access_token,
	})

	customer_id := call0.CustomerId

	call1, _ := srv.ordersClient.AddItem(ctx, &orders.AddItemInput{
		OrderId:  order_id,
		Name:     name,
		Quantity: quantity,
	})

	item_id := call1.ItemId

	output.ItemId = item_id
	return output, nil
}

func main() {
	srv := &FrontendServer{}

	// connect to the auth service via gRPC
	authConn, err := grpc.Dial("auth", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service auth: %v", err)
	}

	defer authConn.Close()

	srv.authClient = auth.NewAuthClient(authConn)

	// connect to the orders service via gRPC
	ordersConn, err := grpc.Dial("orders", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service orders: %v", err)
	}

	defer ordersConn.Close()

	srv.ordersClient = orders.NewOrdersClient(ordersConn)

	// connect to the customers service via gRPC
	customersConn, err := grpc.Dial("customers", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed connecting to service customers: %v", err)
	}

	defer customersConn.Close()

	srv.customersClient = customers.NewCustomersClient(customersConn)
	r := mux.NewRouter()
	r.Methods("POST").PathPrefix("/orders/{order_id}/items").Handler(srv.AddItemHandler)
	r.Methods("POST").PathPrefix("/orders/{order_id}/finish").Handler(srv.FinishOrderHandler)
	r.Methods("POST").PathPrefix("/register").Handler(srv.RegisterHandler)
	r.Methods("POST").PathPrefix("/signin").Handler(srv.SignInHandler)
	r.Methods("POST").PathPrefix("/orders").Handler(srv.StartOrderHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":80", nil))
}
