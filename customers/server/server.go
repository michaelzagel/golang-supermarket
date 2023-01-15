package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/customers"
)

type customersServer struct {
	customers.UnimplementedCustomersServer
	postgresqlDB *sql.DB
	restClient   http.Client
}

func (srv *customersServer) Register(ctx context.Context, input *customers.RegisterInput) (output *customers.RegisterOutput, err error) {
	email := input.Email
	delivery_address := input.DeliveryAddress
	password := input.Password

	var customer_id interface{}
	srv.postgresqlDB.QueryRow("INSERT INTO customers (email, address, password) VALUES ($1, $2, $3) RETURNING customer_id", email, delivery_address, password).Scan(&customer_id)
	output.CustomerId = customer_id

	return output, nil
}

func (srv *customersServer) GetInfo(ctx context.Context, input *customers.GetInfoInput) (output *customers.GetInfoOutput, err error) {
	output, _ = srv.doGetInfo(ctx, input)

	return output, nil
}

func (srv *customersServer) doGetInfo(ctx context.Context, input *customers.GetInfoInput) (
	output *customers.GetInfoOutput,
	err error,
) {
	email := input.Email
	customer_id := input.CustomerId

	output.CustomerId = customer_id
	output.Email = email
	output.DeliveryAddress = delivery_address

	return output, nil
}

func (srv *customersServer) VerifyPassword(ctx context.Context, input *customers.VerifyPasswordInput) (output *customers.VerifyPasswordOutput, err error) {
	output, _ = srv.doVerifyPassword(ctx, input)

	return output, nil
}

func (srv *customersServer) doVerifyPassword(ctx context.Context, input *customers.VerifyPasswordInput) (
	output *customers.VerifyPasswordOutput,
	err error,
) {
	customer_id := input.CustomerId
	password := input.Password

	output.Success = success

	return output, nil
}
