package main

import (
	"context"

	"github.com/backslash-security-tests/golang-supermarket/notification"
)

type notificationServer struct {
}

func (srv *notificationServer) NotifyWindow(ctx context.Context, input *notification.NotifyWindowInput) (output *notification.NotifyWindowOutput, err error) {
	order_id := input.OrderId
	window_from := input.WindowFrom
	window_to := input.WindowTo

	output.Success = success

	return output, nil
}
