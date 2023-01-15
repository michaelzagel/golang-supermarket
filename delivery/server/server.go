package main

import (
	"context"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/delivery"
)

type deliveryServer struct {
	delivery.UnimplementedDeliveryServer
	restClient http.Client
}

func (srv *deliveryServer) SelectWindow(ctx context.Context, input *delivery.SelectWindowInput) (output *delivery.SelectWindowOutput, err error) {
	output, _ = srv.doSelectWindow(ctx, input)

	return output, nil
}

func (srv *deliveryServer) doSelectWindow(ctx context.Context, input *delivery.SelectWindowInput) (
	output *delivery.SelectWindowOutput,
	err error,
) {
	delivery_address := input.DeliveryAddress

	output.WindowFrom = window_from
	output.WindowTo = window_to

	return output, nil
}
