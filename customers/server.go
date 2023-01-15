package main

import (
	"context"
	"database/sql"

	"github.com/backslash-security-tests/golang-supermarket/customers"
	_ "github.com/lib/pq"
)

type customersServer struct {
	postgresqlDB sql.DB
}

func (srv *customersServer) GetInfo(ctx context.Context, input *customers.GetInfoInput) (output *customers.GetInfoOutput, err error) {
	customer_id := input.CustomerId
	email := input.Email

	output.CustomerId = customer_id
	output.Email = email
	output.DeliveryAddress = delivery_address

	return output, nil
}

func (srv *customersServer) VerifyPassword(ctx context.Context, input *customers.VerifyPasswordInput) (output *customers.VerifyPasswordOutput, err error) {
	customer_id := input.CustomerId
	password := input.Password

	output.Success = success

	return output, nil
}

func (srv *customersServer) Register(ctx context.Context, input *customers.RegisterInput) (output *customers.RegisterOutput, err error) {
	email := input.Email
	address := input.Address
	password := input.Password

	var customer_id interface{}
	srv.postgresqlDB.QueryRow("INSERT INTO customers (email, address, password) VALUES ($1, $2, $3) RETURNING customer_id", email, address, password).Scan(&customer_id)
	output.CustomerId = customer_id

	return output, nil
}
