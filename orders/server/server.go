package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/backslash-security-tests/golang-supermarket/customers"
	"github.com/backslash-security-tests/golang-supermarket/delivery"
	"github.com/backslash-security-tests/golang-supermarket/inventory"
	"github.com/backslash-security-tests/golang-supermarket/notification"
	"github.com/backslash-security-tests/golang-supermarket/orders"
	"github.com/backslash-security-tests/golang-supermarket/payments"
	"github.com/backslash-security-tests/golang-supermarket/pricing"
)

type ordersServer struct {
	orders.UnimplementedOrdersServer
	inventoryClient inventory.InventoryClient
	pricingClient   pricing.PricingClient
	paymentsClient  payments.PaymentsClient
	customersClient customers.CustomersClient
	deliveryClient  delivery.DeliveryClient
	postgresqlDB    *sql.DB
	restClient      http.Client
}

func (srv *ordersServer) Start(ctx context.Context, input *orders.StartInput) (output *orders.StartOutput, err error) {
	output, _ = srv.doStart(ctx, input)

	return output, nil
}

func (srv *ordersServer) doStart(ctx context.Context, input *orders.StartInput) (
	output *orders.StartOutput,
	err error,
) {
	customer_id := input.CustomerId

	var order_id interface{}
	srv.postgresqlDB.QueryRow(fmt.Sprintf("INSERT INTO orders (customer_id) VALUES ('%s') RETURNING order_id", customer_id)).Scan(&order_id)
	output.OrderId = order_id

	return output, nil
}

func (srv *ordersServer) AddItem(ctx context.Context, input *orders.AddItemInput) (output *orders.AddItemOutput, err error) {
	order_id := input.OrderId
	name := input.Name
	quantity := input.Quantity

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
	output, _ = srv.doFinish(ctx, input)

	return output, nil
}

func (srv *ordersServer) doFinish(ctx context.Context, input *orders.FinishInput) (
	output *orders.FinishOutput,
	err error,
) {
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

	window_from := call2.WindowFrom
	window_to := call2.WindowTo

	call3Input, _ := json.Marshal(&notification.NotifyWindowInput{
		OrderId:    order_id,
		WindowFrom: window_from,
		WindowTo:   window_to,
	})

	req3, err := http.NewRequest("POST", "http://notification/notify_window", bytes.NewReader(call3Input))
	if err != nil {
		return output, err
	}

	res3, err := srv.restClient.Do(req3)
	if err != nil {
		return output, err
	}

	var call3 notification.NotifyWindowOutput
	err = json.NewDecoder(res3.Body).Decode(&call3)
	if err != nil {
		return output, err
	}

	output.Success = success
	output.WindowFrom = window_from
	output.WindowTo = window_to

	return output, nil
}
