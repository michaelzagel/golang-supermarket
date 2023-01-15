package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/backslash-security-tests/golang-supermarket/customers"
	"github.com/backslash-security-tests/golang-supermarket/delivery"
	"github.com/backslash-security-tests/golang-supermarket/inventory"
	"github.com/backslash-security-tests/golang-supermarket/orders"
	"github.com/backslash-security-tests/golang-supermarket/payments"
	"github.com/backslash-security-tests/golang-supermarket/pricing"
	_ "github.com/lib/pq"
)

type ordersServer struct {
	inventoryClient inventory.inventoryClient
	pricingClient   pricing.pricingClient
	paymentsClient  payments.paymentsClient
	customersClient customers.customersClient
	deliveryClient  delivery.deliveryClient
	postgresqlDB    sql.DB
}

func (srv *ordersServer) Start(ctx context.Context, input *orders.StartInput) (output *orders.StartOutput, err error) {
	customer_id := input.CustomerId

	var order_id interface{}
	srv.postgresqlDB.QueryRow(fmt.Sprintf("INSERT INTO orders (customer_id) VALUES ('%s') RETURNING order_id", customer_id)).Scan(&order_id)
	output.OrderId = order_id

	return output, nil
}

func (srv *ordersServer) AddItem(ctx context.Context, input *orders.AddItemInput) (output *orders.AddItemOutput, err error) {
	quantity := input.Quantity
	order_id := input.OrderId
	name := input.Name

	call0, _ := srv.inventoryClient.FindItem(ctx, &inventory.FindItemInput{
		Name: name,
	})

	sku := call0.Sku

	call1, _ := srv.pricingClient.GetPrice(ctx, &pricing.GetPriceInput{
		Sku:      sku,
		Quantity: quantity,
	})

	price := call1.Price

	var item_id interface{}
	srv.postgresqlDB.QueryRow("INSERT INTO order_items (order_id, sku, quantity, price) VALUES ($1, $2, $3, $4) RETURNING item_id", order_id, sku, quantity, price).Scan(&item_id)
	output.ItemId = item_id

	return output, nil
}

func (srv *ordersServer) Finish(ctx context.Context, input *orders.FinishInput) (output *orders.FinishOutput, err error) {
	credit_card := input.CreditCard
	order_id := input.OrderId

	call0, _ := srv.paymentsClient.TakePayment(ctx, &payments.TakePaymentInput{
		CreditCard: credit_card,
		Amount:     amount,
	})

	call1, _ := srv.customersClient.GetInfo(ctx, &customers.GetInfoInput{
		CustomerId: customer_id,
	})

	delivery_address := call1.DeliveryAddress

	call2, _ := srv.deliveryClient.SelectWindow(ctx, &delivery.SelectWindowInput{
		DeliveryAddress: delivery_address,
	})

	call3, _ := srv.notificationsClient.NotifyWindow(ctx, &notifications.NotifyWindowInput{
		OrderId:    order_id,
		WindowFrom: window_from,
		WindowTo:   window_to,
	})

	output.Success = success

	return output, nil
}
