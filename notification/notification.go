package notification

type NotifyWindowInput struct {
	WindowTo   string `json:"window_to"`
	OrderId    int64  `json:"order_id"`
	WindowFrom string `json:"window_from"`
}

type NotifyWindowOutput struct {
	Success bool `json:"success"`
}
