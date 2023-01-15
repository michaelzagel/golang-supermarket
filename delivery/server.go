package main

import (
	"context"

	"github.com/backslash-security-tests/golang-supermarket/delivery"
)

type deliveryServer struct {
}

func (srv *deliveryServer) SelectWindow(ctx context.Context, input *delivery.SelectWindowInput) (output *delivery.SelectWindowOutput, err error) {
	delivery_address := input.DeliveryAddress

	output.WindowFrom = window_from
	output.WindowTo = window_to

	return output, nil
}
