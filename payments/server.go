package main

import (
	"context"

	"github.com/backslash-security-tests/golang-supermarket/payments"
)

type paymentsServer struct {
}

func (srv *paymentsServer) TakePayment(ctx context.Context, input *payments.TakePaymentInput) (output *payments.TakePaymentOutput, err error) {
	amount := input.Amount
	credit_card := input.CreditCard

	output.Success = success

	return output, nil
}
