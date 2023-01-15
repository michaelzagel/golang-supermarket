package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/auth"
	"github.com/backslash-security-tests/golang-supermarket/customers"
	"github.com/backslash-security-tests/golang-supermarket/orders"
	"github.com/gorilla/mux"
)

type FrontendServer struct {
	authClient      auth.authClient
	ordersClient    orders.ordersClient
	customersClient customers.customersClient
}

type AddItemInput struct {
	Quantity    float32 `json:"quantity"`
	AccessToken string  `json:"access_token"`
	OrderId     int64   `json:"order_id"`
	Name        string  `json:"name"`
}

type AddItemOutput struct {
	ItemId int64 `json:"item_id"`
}

func (srv *FrontendServer) AddItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input AddItemInput
	var output AddItemOutput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")
	input.OrderId = mux.Vars(r)["order_id"]

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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

type FinishOrderInput struct {
	AccessToken string `json:"access_token"`
	OrderId     int64  `json:"order_id"`
	CreditCard  string `json:"credit_card"`
}

type FinishOrderOutput struct {
	WindowFrom string `json:"window_from"`
	WindowTo   string `json:"window_to"`
}

func (srv *FrontendServer) FinishOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input FinishOrderInput
	var output FinishOrderOutput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")
	input.OrderId = mux.Vars(r)["order_id"]

	order_id := input.OrderId
	credit_card := input.CreditCard
	access_token := input.AccessToken

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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

type RegisterInput struct {
	Email           string `json:"email"`
	DeliveryAddress string `json:"delivery_address"`
	Password        string `json:"password"`
}

type RegisterOutput struct {
	CustomerId int64 `json:"customer_id"`
}

func (srv *FrontendServer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input RegisterInput
	var output RegisterOutput
	json.NewDecoder(r.Body).Decode(&input)

	email := input.Email
	delivery_address := input.DeliveryAddress
	password := input.Password

	call0, _ := srv.customersClient.Register(ctx, &customers.RegisterInput{
		Email:           email,
		DeliveryAddress: delivery_address,
		Password:        password,
	})

	customer_id := call0.CustomerId

	output.CustomerId = customer_id
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInOutput struct {
	Success     bool   `json:"success"`
	AccessToken string `json:"access_token"`
}

func (srv *FrontendServer) SignInHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input SignInInput
	var output SignInOutput
	json.NewDecoder(r.Body).Decode(&input)

	email := input.Email
	password := input.Password

	call0, _ := srv.authClient.SignIn(ctx, &auth.SignInInput{
		Email:    email,
		Password: password,
	})

	access_token := call0.AccessToken

	output.Success = success
	output.AccessToken = access_token
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

type StartOrderInput struct {
	AccessToken string `json:"access_token"`
}

type StartOrderOutput struct {
	OrderId int64 `json:"order_id"`
}

func (srv *FrontendServer) StartOrderHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input StartOrderInput
	var output StartOrderOutput
	json.NewDecoder(r.Body).Decode(&input)
	input.AccessToken = r.Header.Get("Authorization")

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
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

func main() {
	srv := &FrontendServer{}
	r := mux.NewRouter()
	r.Methods("POST").PathPrefix("/register").Handler(srv.RegisterHandler)
	r.Methods("POST").PathPrefix("/signin").Handler(srv.SignInHandler)
	r.Methods("POST").PathPrefix("/orders").Handler(srv.StartOrderHandler)
	r.Methods("POST").PathPrefix("/orders/{order_id}/items").Handler(srv.AddItemHandler)
	r.Methods("POST").PathPrefix("/orders/{order_id}/finish").Handler(srv.FinishOrderHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
