package main

import (
	"context"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/pricing"
)

type pricingServer struct {
	pricing.UnimplementedPricingServer
	restClient http.Client
}

func (srv *pricingServer) GetPrice(ctx context.Context, input *pricing.GetPriceInput) (output *pricing.GetPriceOutput, err error) {
	output, _ = srv.doGetPrice(ctx, input)

	return output, nil
}

func (srv *pricingServer) doGetPrice(ctx context.Context, input *pricing.GetPriceInput) (
	output *pricing.GetPriceOutput,
	err error,
) {
	quantity := input.Quantity
	sku := input.Sku

	output.Price = price

	return output, nil
}
