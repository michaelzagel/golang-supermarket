package main

import (
	"context"

	"github.com/backslash-security-tests/golang-supermarket/pricing"
)

type pricingServer struct {
}

func (srv *pricingServer) GetPrice(ctx context.Context, input *pricing.GetPriceInput) (output *pricing.GetPriceOutput, err error) {
	sku := input.Sku
	quantity := input.Quantity

	output.Price = price

	return output, nil
}
