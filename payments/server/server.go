package main

import (
	"context"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/payments"
)

type paymentsServer struct {
	payments.UnimplementedPaymentsServer
	restClient http.Client
}

func (srv *paymentsServer) TakePayment(ctx context.Context, input *payments.TakePaymentInput) (output *payments.TakePaymentOutput, err error) {
	credit_card := input.CreditCard
	amount := input.Amount

	output.Success = success

	return output, nil
}
