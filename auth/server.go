package main

import (
	"context"

	"github.com/backslash-security-tests/golang-supermarket/auth"
	"github.com/backslash-security-tests/golang-supermarket/customers"
)

type authServer struct {
	customersClient customers.customersClient
}

func (srv *authServer) SignIn(ctx context.Context, input *auth.SignInInput) (output *auth.SignInOutput, err error) {
	email := input.Email
	password := input.Password

	call0, _ := srv.customersClient.GetInfo(ctx, &customers.GetInfoInput{
		Email: email,
	})

	customer_id := call0.CustomerId

	call1, _ := srv.customersClient.VerifyPassword(ctx, &customers.VerifyPasswordInput{
		CustomerId: customer_id,
		Password:   password,
	})

	output.AccessToken = access_token

	return output, nil
}

func (srv *authServer) VerifyToken(ctx context.Context, input *auth.VerifyTokenInput) (output *auth.VerifyTokenOutput, err error) {
	customer_id := input.CustomerId
	access_token := input.AccessToken

	output.Success = success

	return output, nil
}
